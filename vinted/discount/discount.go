package discount

import (
	"fmt"
	"math/rand"
	"practice/vinted/shipping/courier"
	"practice/vinted/size"
	"strconv"
	"time"
)

const (
	countOfYears         = 10
	countOfMonthsInYears = 12 * countOfYears
)

type calendar map[string]float64

type discount struct {
	budget   float64
	calendar calendar
}

func NewDiscountEngine(budget float64) *discount {
	return &discount{
		budget:   budget,
		calendar: make(map[string]float64, countOfMonthsInYears),
	}
}

type request interface {
	GetShippingSize() size.Size
	GetCourier() *courier.Courier
	GetShippingTime() time.Time
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

func (d *discount) Apply(request request) (response *ShipmentResponse) {
	var plannedDiscount float64

	key := d.CreateKey(request.GetShippingTime())
	fmt.Printf("key: %v\n", key)
	fmt.Printf("d.budget: %v\n", d.budget)

	monthlySpent, ok := d.calendar[key]
	if !ok {
		d.calendar[key] = 0
	}

	if monthlySpent < d.budget {
		rand.Seed(time.Now().UnixNano())

		plannedDiscount = request.GetCourier().GetPrice(request.GetShippingSize()) * rand.Float64()
		if monthlySpent-plannedDiscount >= 0 {
			monthlySpent += plannedDiscount
			d.calendar[key] = monthlySpent
		}
	}

	fmt.Printf("d.calendar: %v\n", d.calendar)

	return &ShipmentResponse{
		shippingTime:  request.GetShippingTime(),
		shippingSize:  request.GetShippingSize(),
		courier:       request.GetCourier(),
		discountPrice: plannedDiscount,
	}
}

func (d *discount) CreateKey(date time.Time) string {
	return strconv.Itoa(date.Year()) + "-" + date.Month().String()
}
