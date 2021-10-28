package courier

import "practice/vinted/size"

type Courier struct {
	name    string
	sizes   []size.Size
	pricing map[size.Size]float64
}

func NewCourier(name string, pricing map[size.Size]float64) *Courier {
	return &Courier{name: name, pricing: pricing}
}

func (c *Courier) GetName() string {
	return c.name
}

func (c *Courier) GetSizes() []size.Size {
	if c.sizes != nil {
		return c.sizes
	}

	c.sizes = make([]size.Size, len(c.pricing))
	for size := range c.pricing {
		c.sizes = append(c.sizes, size)
	}

	return c.sizes
}

func (c *Courier) GetPrice(size size.Size) float64 {
	return c.pricing[size]
}

func (c *Courier) SetPrice(size size.Size, price float64) {
	c.pricing[size] = price
}
