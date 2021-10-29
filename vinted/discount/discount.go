package discount

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

type calendar map[string]float64

type Discount struct {
	budget   float64
	calendar calendar
}

func NewDiscountEngine(budget float64) *Discount {
	return &Discount{
		budget:   budget,
		calendar: make(map[string]float64, countOfMonthsInYears),
	}
}

type request interface {
	GetShippingSize() size.Size
	GetCourier() *courier.Courier
	GetShippingTime() time.Time
	GetPrice() float64
}

type ShipmentResponse struct {
	shippingTime  time.Time
	shippingSize  size.Size
	courier       *courier.Courier
	price         float64
	discountPrice float64
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

// GetDiscountPrice returns the float64 discountPrice
func (s ShipmentResponse) GetDiscountPrice() float64 {
	return s.discountPrice
}

func (d *Discount) Apply(request request) (response *ShipmentResponse) {
	var plannedDiscount float64

	key := d.CreateKey(request.GetShippingTime())
	// fmt.Printf("key: %v\n", key)
	// fmt.Printf("d.budget: %v\n", d.budget)

	monthlySpent, ok := d.calendar[key]
	if !ok {
		d.calendar[key] = 0
	}

	if monthlySpent < d.budget {
		fullPrice := request.GetCourier().GetPrice(request.GetShippingSize())
		plannedDiscount = fullPrice - request.GetPrice()

		if monthlySpent-plannedDiscount >= 0 {
			monthlySpent += plannedDiscount
			d.calendar[key] = monthlySpent
		}
	}

	//fmt.Printf("d.calendar: %v\n", d.calendar)

	return &ShipmentResponse{
		shippingTime:  request.GetShippingTime(),
		shippingSize:  request.GetShippingSize(),
		courier:       request.GetCourier(),
		price:         request.GetPrice(),
		discountPrice: plannedDiscount,
	}
}

func (d *Discount) CreateKey(date time.Time) string {
	return strconv.Itoa(date.Year()) + "-" + date.Month().String()
}
