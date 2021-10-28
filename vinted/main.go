package main

import (
	"fmt"
	"practice/vinted/discount"
	"practice/vinted/rule"
	"practice/vinted/shipping/courier"
	"practice/vinted/shipping/shipment"
	"practice/vinted/size"
	"time"
)

const budget = 10

func main() {
	// Start shipping service for a new country
	s := NewShipping("France")

	// Add couriers
	s.addCourier("LP", map[size.Size]float64{size.Small: 1.50, size.Medium: 4.90, size.Large: 6.90})
	s.addCourier("MR", map[size.Size]float64{size.Small: 2, size.Medium: 3, size.Large: 4})

	// Initialize shipping rule engine
	ruleEngine := rule.NewRuleEngine(s.couriers)

	// Initialize discount rule engine
	discountEngine := discount.NewDiscountEngine(budget)

	// parse the shipment info
	shipmentTime, shipmentSize, courierName := parseShipmentInfo()

	provider := s.couriers[courierName]
	shipmentRequest := shipment.NewShipment(shipmentTime, shipmentSize, provider)

	// Rule check the shipment
	outShipment := ruleEngine.Process(shipmentRequest)

	// Calculate the discount price
	discountedShipment := discountEngine.Apply(outShipment)
	fmt.Printf("discountedShipment: %v %v %v\n", discountedShipment.GetCourier().GetName(), discountedShipment.GetCourier().GetPrice(shipmentSize), discountedShipment.GetDiscountPrice())
}

type Country string

type shipping struct {
	country  Country
	couriers map[string]*courier.Courier
}

func NewShipping(countryName string) *shipping {
	return &shipping{
		country:  Country(countryName),
		couriers: make(map[string]*courier.Courier, 10),
	}
}

func (s *shipping) addCourier(courierName string, pricing map[size.Size]float64) {
	c := courier.NewCourier(courierName, pricing)
	s.couriers[courierName] = c
}

func parseShipmentInfo() (time.Time, size.Size, string) {
	return time.Now(), size.Small, "MR"
}
