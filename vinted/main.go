package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"practice/vinted/discount"
	"practice/vinted/rule"
	"practice/vinted/shipping/courier"
	"practice/vinted/shipping/shipment"
	"practice/vinted/size"
	"strings"
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
	s.ruleEngine = rule.NewRuleEngine(s.couriers)

	// Initialize discount rule engine
	s.discountEngine = discount.NewDiscountEngine(budget)

	s.processInput()
}

type Country string

type shipping struct {
	country  Country
	couriers map[string]*courier.Courier

	ruleEngine     *rule.RuleEngine
	discountEngine *discount.Discount
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

func (s *shipping) processInput() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		transaction := strings.Split(line, " ")
		if !s.validate(transaction) {
			transaction = append(transaction, "IGNORED")
			fmt.Println(strings.Join(transaction, " "))
			continue
		}

		// parse the shipment info
		shipmentTime, _ := time.Parse("2006-01-02", transaction[0])

		shipmentSize := size.Size(transaction[1])

		courierName := string(transaction[2])

		provider := s.couriers[courierName]
		shipmentRequest := shipment.NewShipment(shipmentTime, shipmentSize, provider)

		// Rule check the shipment
		outShipment := s.ruleEngine.Process(shipmentRequest)

		// Calculate the discount price
		discountedShipment := s.discountEngine.Apply(outShipment)

		// Print the result
		discountPrice := discountedShipment.GetDiscountPrice()
		fmt.Printf("%v %s %s %0.2f ", discountedShipment.GetShippingTime().Format("2006-01-02"), discountedShipment.GetShippingSize(), discountedShipment.GetCourier().GetName(), discountedShipment.GetPrice())
		if discountPrice == 0 {
			fmt.Printf("%v\n", "-")
		} else {
			fmt.Printf("%0.2f\n", discountedShipment.GetDiscountPrice())
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
}

func (s *shipping) validate(transaction []string) bool {
	if len(transaction) != 3 {
		return false
	}

	_, err := time.Parse("2006-01-02", transaction[0])
	if err != nil {
		fmt.Printf("Invalid time format %v\n", err)
		return false
	}

	sz := transaction[1]
	provider := string(transaction[2])

	_, ok := s.couriers[provider]
	if !ok {
		return false
	}

	if size.Size(sz) == size.Unknown {
		return false
	}

	return true
}
