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

type calendar map[string]int

type Discount struct {
	budget   float64
	calendar calendar
}

func NewDiscountEngine(budget float64) *Discount {
	return &Discount{
		budget:   budget,
		calendar: make(map[string]int, countOfMonthsInYears),
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

// GetPrice returns the int price
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
	var plannedDiscount int

	key := d.CreateKey(request.GetShippingTime())

	// fmt.Printf("key: %v\n", key)
	monthlySpent, ok := d.calendar[key]
	if !ok {
		d.calendar[key] = int(d.budget * 100)
	}

	// fmt.Printf("monthlySpent: %v\n", monthlySpent)
	if monthlySpent > 0 {
		var fullPrice int = int(float64(100) * request.GetCourier().GetPrice(request.GetShippingSize()))

		// fmt.Printf("request.GetPrice(): %v\n", request.GetPrice())
		plannedDiscount = fullPrice - int(float64(100)*request.GetPrice())

		if monthlySpent-plannedDiscount >= 0 {
			monthlySpent -= plannedDiscount
			d.calendar[key] = monthlySpent

			return &ShipmentResponse{
				shippingTime:  request.GetShippingTime(),
				shippingSize:  request.GetShippingSize(),
				courier:       request.GetCourier(),
				price:         request.GetPrice(),
				discountPrice: float64(plannedDiscount) / float64(100),
			}
		}

		// partial discount
		if monthlySpent-plannedDiscount < 0 {
			d.calendar[key] -= monthlySpent

			discountPrice := float64(monthlySpent) / float64(100)

			return &ShipmentResponse{
				shippingTime:  request.GetShippingTime(),
				shippingSize:  request.GetShippingSize(),
				courier:       request.GetCourier(),
				price:         request.GetCourier().GetPrice(request.GetShippingSize()) - discountPrice,
				discountPrice: discountPrice,
			}
		}
	}

	return &ShipmentResponse{
		shippingTime:  request.GetShippingTime(),
		shippingSize:  request.GetShippingSize(),
		courier:       request.GetCourier(),
		price:         request.GetCourier().GetPrice(request.GetShippingSize()),
		discountPrice: 0,
	}
}

func (d *Discount) CreateKey(date time.Time) string {
	return strconv.Itoa(date.Year()) + "-" + date.Month().String()
}
