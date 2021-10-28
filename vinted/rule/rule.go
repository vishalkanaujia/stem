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
	maxPrice             = 1000
)

type calendar map[string]int

type ruleEngine struct {
	couriers map[string]*courier.Courier
	calendar calendar
}

func NewRuleEngine(couriers map[string]*courier.Courier) *ruleEngine {
	return &ruleEngine{
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

func (r *ruleEngine) Process(ruleEngineRequest request) *ShipmentResponse {
	processedRequest := r.LowestPriceBySize(ruleEngineRequest)
	processedRequest = r.LowestPriceBySize(processedRequest)

	return processedRequest
}

// Add rules as needed by business
func (r *ruleEngine) LowestPriceBySize(request request) *ShipmentResponse {
	var bestProvider *courier.Courier = request.GetCourier()

	if request.GetShippingSize() == size.Small {
		var lowestPrice float64 = maxPrice

		for _, provider := range r.couriers {
			price := provider.GetPrice(size.Small)
			if price < lowestPrice {
				lowestPrice = price
				bestProvider = provider
			}
		}
	}

	return &ShipmentResponse{
		shippingTime: request.GetShippingTime(),
		shippingSize: request.GetShippingSize(),
		courier:      bestProvider,
	}
}

func (r *ruleEngine) FreeShipmentByProvider(request request) *ShipmentResponse {
	if request.GetCourier().GetName() == "LP" {
		key := r.CreateKey(request.GetShippingTime())

		_, ok := r.calendar[key]
		if !ok {
			r.calendar[key] = 1
		} else {
			r.calendar[key]++
		}

		if r.calendar[key]%3 == 0 {
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
	}
}

func (r *ruleEngine) CreateKey(date time.Time) string {
	return strconv.Itoa(date.Year()) + "-" + date.Month().String()
}
