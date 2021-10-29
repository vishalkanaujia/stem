package rule

import (
	"fmt"
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
	var shipmentResponse *ShipmentResponse
	shipmentResponse = &ShipmentResponse{
		shippingTime: ruleEngineRequest.GetShippingTime(),
		shippingSize: ruleEngineRequest.GetShippingSize(),
		courier:      ruleEngineRequest.GetCourier(),
		price:        ruleEngineRequest.GetCourier().GetPrice(ruleEngineRequest.GetShippingSize()),
	}

	shipmentResponse = r.LowestPriceBySize(shipmentResponse)
	fmt.Printf("after lowest:: processedRequest: %v\n", shipmentResponse)

	shipmentResponse = r.FreeShipmentByProvider(shipmentResponse)

	return shipmentResponse
}

// Add rules as needed by business
func (r *ruleEngine) LowestPriceBySize(request *ShipmentResponse) *ShipmentResponse {
	var bestProvider *courier.Courier = request.GetCourier()
	var lowestPrice float64 = request.GetCourier().GetPrice(request.GetShippingSize())
	fmt.Printf("xxx lowestPrice: %v\n", lowestPrice)

	if request.GetShippingSize() == size.Small {
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
		price:        lowestPrice,
	}
}

func (r *ruleEngine) FreeShipmentByProvider(request *ShipmentResponse) *ShipmentResponse {
	fmt.Printf("zzz r.calendar: %v\n", r.calendar)
	if request.GetCourier().GetName() == "LP" && request.GetShippingSize() == size.Large {
		key := r.CreateKey(request.GetShippingTime())

		_, ok := r.calendar[key]
		if !ok {
			r.calendar[key] = 1
		} else {
			r.calendar[key]++
		}

		fmt.Printf("zzz r.calendar: %v\n", r.calendar)

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
		price:        request.GetCourier().GetPrice(request.GetShippingSize()),
	}
}

func (r *ruleEngine) CreateKey(date time.Time) string {
	return strconv.Itoa(date.Year()) + "-" + date.Month().String()
}
