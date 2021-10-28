package rule

import (
	"practice/vinted/shipping/courier"
	"practice/vinted/size"
	"strconv"
	"time"
)

const (
	countOfYears         = 10
	countOfMonthsInYears = 12 * countOfYears
)

type calendar map[string]int

type RuleEngine struct {
	couriers map[string]*courier.Courier
	calendar calendar
}

func NewRuleEngine(couriers map[string]*courier.Courier) *RuleEngine {
	return &RuleEngine{
		couriers: couriers,
		calendar: make(map[string]int, countOfMonthsInYears),
	}
}

type request interface {
	GetShippingSize() size.Size
	GetCourier() *courier.Courier
	GetShippingTime() time.Time
}

type ShipmentResponse struct {
	shippingTime time.Time
	shippingSize size.Size
	courier      *courier.Courier
	price        float64
}

// GetPrice returns the float64 price
func (s ShipmentResponse) GetPrice() float64 {
	return s.price
}

// GetShippingTime returns the time.Time shippingTime
func (s ShipmentResponse) GetShippingTime() time.Time {
	return s.shippingTime
}

// GetShippingSize returns the size.Size shippingSize
func (s ShipmentResponse) GetShippingSize() size.Size {
	return s.shippingSize
}

// GetCourier returns the *courier.Courier courier
func (s ShipmentResponse) GetCourier() *courier.Courier {
	return s.courier
}

func (r *RuleEngine) Process(ruleEngineRequest request) *ShipmentResponse {
	var shipmentResponse *ShipmentResponse
	shipmentResponse = &ShipmentResponse{
		shippingTime: ruleEngineRequest.GetShippingTime(),
		shippingSize: ruleEngineRequest.GetShippingSize(),
		courier:      ruleEngineRequest.GetCourier(),
		price:        ruleEngineRequest.GetCourier().GetPrice(ruleEngineRequest.GetShippingSize()),
	}

	shipmentResponse = r.LowestPriceBySize(shipmentResponse)

	shipmentResponse = r.FreeShipmentByProvider(shipmentResponse)

	return shipmentResponse
}

// Add rules as needed by business
func (r *RuleEngine) LowestPriceBySize(request *ShipmentResponse) *ShipmentResponse {
	var lowestPrice float64 = request.GetCourier().GetPrice(request.GetShippingSize())

	if request.GetShippingSize() == size.Small {
		for _, provider := range r.couriers {
			price := provider.GetPrice(size.Small)
			if price < lowestPrice {
				lowestPrice = price
			}
		}
	}

	return &ShipmentResponse{
		shippingTime: request.GetShippingTime(),
		shippingSize: request.GetShippingSize(),
		courier:      request.GetCourier(),
		price:        lowestPrice,
	}
}

func (r *RuleEngine) FreeShipmentByProvider(request *ShipmentResponse) *ShipmentResponse {
	if request.GetCourier().GetName() == "LP" && request.GetShippingSize() == size.Large {
		key := r.CreateKey(request.GetShippingTime())

		_, ok := r.calendar[key]
		if !ok {
			r.calendar[key] = 1
		} else {
			if r.calendar[key] != 0 {
				r.calendar[key]++
			}
		}

		// fmt.Printf("r.calendar[key]: %v key=%v\n", r.calendar[key], key)
		if r.calendar[key] != 0 && r.calendar[key]%3 == 0 {
			r.calendar[key] = 0

			return &ShipmentResponse{
				shippingTime: request.GetShippingTime(),
				shippingSize: request.GetShippingSize(),
				courier:      request.GetCourier(),
				price:        0,
			}
		}
	}

	return &ShipmentResponse{
		shippingTime: request.GetShippingTime(),
		shippingSize: request.GetShippingSize(),
		courier:      request.GetCourier(),
		price:        request.GetPrice(),
	}
}

func (r *RuleEngine) CreateKey(date time.Time) string {
	return strconv.Itoa(date.Year()) + "-" + date.Month().String()
}
