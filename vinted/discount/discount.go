package discount

import (
	"practice/vinted/shipping/courier"
	"practice/vinted/size"
	"time"
)

type request interface {
	GetSize() size.Size
	GetCourier() *courier.Courier
	GetTime() time.Time
}

type ShipmentResponse struct {
	shippingTime  time.Time
	shippingSize  size.Size
	courier       *courier.Courier
	discountPrice float64
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

// GetDiscountPrice returns the float64 discountPrice
func (s ShipmentResponse) GetDiscountPrice() float64 {
	return s.discountPrice
}

func Apply(request request) (response *ShipmentResponse) {
	return &ShipmentResponse{
		shippingTime:  time.Now(),
		shippingSize:  request.GetSize(),
		courier:       request.GetCourier(),
		discountPrice: 0,
	}
}
