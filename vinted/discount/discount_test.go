package discount

import (
	"practice/vinted/shipping/courier"
	"practice/vinted/size"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDiscount_Apply(t *testing.T) {
	testCouriers := make(map[string]*courier.Courier, 2)

	testCouriers["LP"] = courier.NewCourier("LP", map[size.Size]float64{size.Small: 1.50, size.Medium: 4.90, size.Large: 6.90})
	testCouriers["MR"] = courier.NewCourier("MR", map[size.Size]float64{size.Small: 2, size.Medium: 3, size.Large: 4})

	scenarios := []struct {
		desc             string
		inBudget         float64
		inRequest        []*testRequest
		expectedResponse []*ShipmentResponse
	}{
		{
			desc:     "discount applied",
			inBudget: 10,
			inRequest: []*testRequest{
				{
					shippingTime: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
					shippingSize: size.Small,
					courier:      testCouriers["MR"],
					price:        1.50,
				},
				{
					shippingTime: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
					shippingSize: size.Large,
					courier:      testCouriers["MR"],
					price:        4,
				},
			},
			expectedResponse: []*ShipmentResponse{
				{
					shippingTime:  time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
					shippingSize:  size.Small,
					courier:       testCouriers["LP"],
					discountPrice: 0.50,
				},
				{
					shippingTime:  time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
					shippingSize:  size.Large,
					courier:       testCouriers["MR"],
					discountPrice: 0,
				},
			},
		},
		{
			desc:     "discount not applied",
			inBudget: 0,
			inRequest: []*testRequest{
				{
					shippingTime: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
					shippingSize: size.Small,
					courier:      testCouriers["LP"],
					price:        0,
				},
			},
			expectedResponse: []*ShipmentResponse{
				{
					shippingTime:  time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
					shippingSize:  size.Small,
					courier:       testCouriers["LP"],
					discountPrice: 0,
				},
			},
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.desc, func(t *testing.T) {
			discount := NewDiscountEngine(scenario.inBudget)

			for i, request := range scenario.inRequest {
				response := discount.Apply(request)
				assert.NotNil(t, response)
				assert.True(t, scenario.expectedResponse[i].GetDiscountPrice() <= request.courier.GetPrice(request.GetShippingSize()), scenario.desc)
			}
		})
	}
}

type testRequest struct {
	shippingTime time.Time
	shippingSize size.Size
	courier      *courier.Courier
	price        float64
}

// GetPrice returns the float64 price
func (t testRequest) GetPrice() float64 {
	return t.price
}

// GetShippingTime returns the time.Time shippingTime
func (t *testRequest) GetShippingTime() time.Time {
	return t.shippingTime
}

// GetShippingSize returns the size.Size shippingSize
func (t *testRequest) GetShippingSize() size.Size {
	return t.shippingSize
}

// GetCourier returns the *courier.Courier courier
func (t *testRequest) GetCourier() *courier.Courier {
	return t.courier
}
