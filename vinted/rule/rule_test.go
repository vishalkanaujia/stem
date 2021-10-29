package rule

import (
	"practice/vinted/shipping/courier"
	"practice/vinted/size"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRule_Process(t *testing.T) {
	testCouriers := make(map[string]*courier.Courier, 2)

	testCouriers["LP"] = courier.NewCourier("LP", map[size.Size]float64{size.Small: 1.50, size.Medium: 4.90, size.Large: 6.90})
	testCouriers["MR"] = courier.NewCourier("MR", map[size.Size]float64{size.Small: 2, size.Medium: 3, size.Large: 4})

	scenarios := []struct {
		desc              string
		inRequests        []*testRequest
		expectedResponses []*ShipmentResponse
	}{
		{
			desc: "lowest price shipment for size Small",
			inRequests: []*testRequest{
				{
					shippingTime: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
					shippingSize: size.Small,
					courier:      testCouriers["LP"],
				},
				{
					shippingTime: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
					shippingSize: size.Small,
					courier:      testCouriers["MR"],
				},
				{
					shippingTime: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
					shippingSize: size.Large,
					courier:      testCouriers["MR"],
				},
			},
			expectedResponses: []*ShipmentResponse{
				{
					shippingTime: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
					shippingSize: size.Small,
					courier:      testCouriers["LP"],
					price:        1.50,
				},
				{
					shippingTime: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
					shippingSize: size.Small,
					courier:      testCouriers["LP"],
					price:        1.50,
				},
				{
					shippingTime: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
					shippingSize: size.Large,
					courier:      testCouriers["MR"],
					price:        4,
				},
			},
		},
		{
			desc: "the third L shipment via LP is free in a calende month",
			inRequests: []*testRequest{
				{
					shippingTime: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
					shippingSize: size.Large,
					courier:      testCouriers["LP"],
				},
				{
					shippingTime: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
					shippingSize: size.Large,
					courier:      testCouriers["LP"],
				},
				{
					shippingTime: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
					shippingSize: size.Large,
					courier:      testCouriers["LP"],
				},
				{
					shippingTime: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
					shippingSize: size.Large,
					courier:      testCouriers["LP"],
				},
			},
			expectedResponses: []*ShipmentResponse{
				{
					shippingTime: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
					shippingSize: size.Large,
					courier:      testCouriers["LP"],
					price:        6.90,
				},
				{
					shippingTime: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
					shippingSize: size.Large,
					courier:      testCouriers["LP"],
					price:        6.90,
				},
				{
					shippingTime: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
					shippingSize: size.Large,
					courier:      testCouriers["LP"],
					price:        0,
				},
				{
					shippingTime: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
					shippingSize: size.Large,
					courier:      testCouriers["LP"],
					price:        6.90,
				},
			},
		},
		{
			desc: "the third L shipment via LP is not free across calender month",
			inRequests: []*testRequest{
				{
					shippingTime: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
					shippingSize: size.Large,
					courier:      testCouriers["LP"],
				},
				{
					shippingTime: time.Date(2020, time.February, 1, 0, 0, 0, 0, time.UTC),
					shippingSize: size.Large,
					courier:      testCouriers["LP"],
				},
				{
					shippingTime: time.Date(2020, time.March, 1, 0, 0, 0, 0, time.UTC),
					shippingSize: size.Large,
					courier:      testCouriers["LP"],
				},
				{
					shippingTime: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
					shippingSize: size.Large,
					courier:      testCouriers["LP"],
				},
			},
			expectedResponses: []*ShipmentResponse{
				{
					shippingTime: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
					shippingSize: size.Large,
					courier:      testCouriers["LP"],
					price:        6.90,
				},
				{
					shippingTime: time.Date(2020, time.February, 1, 0, 0, 0, 0, time.UTC),
					shippingSize: size.Large,
					courier:      testCouriers["LP"],
					price:        6.90,
				},
				{
					shippingTime: time.Date(2020, time.March, 1, 0, 0, 0, 0, time.UTC),
					shippingSize: size.Large,
					courier:      testCouriers["LP"],
					price:        6.90,
				},
				{
					shippingTime: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
					shippingSize: size.Large,
					courier:      testCouriers["LP"],
					price:        6.90,
				},
			},
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.desc, func(t *testing.T) {
			for _, scenario := range scenarios {
				ruleEngine := NewRuleEngine(testCouriers)
				for i, request := range scenario.inRequests {
					response := ruleEngine.Process(request)
					assert.Equal(t, scenario.expectedResponses[i].price, response.price, scenario.desc)
				}
			}
		})
	}
}

type testRequest struct {
	shippingTime time.Time
	shippingSize size.Size
	courier      *courier.Courier
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
