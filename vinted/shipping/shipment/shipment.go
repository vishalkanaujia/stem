package shipment

import (
	"practice/vinted/shipping/courier"
	"practice/vinted/size"
	"time"
)

// Create a shipment
func NewShipment(shippingTime time.Time, shippingSize size.Size, courierService *courier.Courier) *Shipment {
	return &Shipment{shippingTime, shippingSize, courierService}
}

type Shipment struct {
	shippingTime time.Time
	shippingSize size.Size
	courier      *courier.Courier
}

// GetShippingTime returns the time.Time shippingTime
func (s *Shipment) GetShippingTime() time.Time {
	return s.shippingTime
}

// GetShippingSize returns the size.Size shippingSize
func (s *Shipment) GetShippingSize() size.Size {
	return s.shippingSize
}

// GetCourier returns the *courier.Courier courier
func (s *Shipment) GetCourier() *courier.Courier {
	return s.courier
}
