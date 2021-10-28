package main

import (
	"fmt"
	"practice/vinted/discount"
	ruleengine "practice/vinted/rule"
	"practice/vinted/shipping/courier"
	"practice/vinted/shipping/shipment"
	"practice/vinted/size"
	"time"
)

func main() {
	// Start shipping service for a new country
	s := NewShipping("France")

	// Add couriers
	s.addCourier("LP", map[size.Size]float64{size.Small: 1.50, size.Medium: 4.90, size.Large: 6.90})
	s.addCourier("MR", map[size.Size]float64{size.Small: 2, size.Medium: 3, size.Large: 4})

	// parse the shipment info
	shipmentTime, shipmentSize, courierName := parseShipmentInfo()
	shipmentRequest := shipment.NewShipment(shipmentTime, shipmentSize, courierName)

	// Rule check the shipment
	outShipment := ruleengine.Process(shipmentRequest)

	// Calculate the discount price
	discountedShipment := discount.Apply(outShipment)
	fmt.Printf("discountedShipment: %v\n", discountedShipment)
}

type Country string

type shipping struct {
	country  Country
	couriers map[string]*courier.Courier
}

func NewShipping(countryName string) *shipping {
	return &shipping{country: Country(countryName)}
}

func (s *shipping) addCourier(courierName string, pricing map[size.Size]float64) {
	c := courier.NewCourier(courierName, pricing)
	s.couriers[courierName] = c
}

func parseShipmentInfo() (time.Time, size.Size, string) {
	return time.Now(), size.Small, "LP"
}
