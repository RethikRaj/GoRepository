// Embedded structs => Composition + Field/Method promotion
package structs

import "fmt"

type CreditCard struct {
    provider string
    limit    int
}

func (c CreditCard) ShowLimit() {
    fmt.Println("Limit:", c.limit)
}

type RupayCard struct {
    CreditCard  // embedded
    rewardRate  float64
}

func Struct4() {
	r := RupayCard{
		CreditCard: CreditCard{provider: "Rupay", limit: 50000},
		rewardRate: 1.2,
	}
	fmt.Println(r)
	// Field Promotion => Limit , provider are directly accesible
	fmt.Println(r.limit, r.provider, r.CreditCard.limit, r.CreditCard.provider, r.rewardRate) 
	// Method Promotion 
	r.ShowLimit()
}