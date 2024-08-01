package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/caarlos0/env"
	"github.com/stripe/stripe-go/v79"
	"github.com/stripe/stripe-go/v79/checkout/session"
)

type config struct {
	Key       string `env:"stripe_key"`
	DomainURL string `env:"domainUrl" envDefault:"localhost"`
	Port      string `env:"port" envDefault:"4242"`
	HostUrl   string `env:"hostUrl" envDefault:"http://localhost:4242"`
}

type StripeBody struct {
	FoodName   string `json: "foodName"`
	Qty        int64  `json: "qty"`
	CouponCode string `json: "couponCode"`
}

func main() {
	// This is your test secret API key.
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}

	stripe.Key = cfg.Key

	http.Handle("/", http.FileServer(http.Dir("public")))
	http.HandleFunc("/create-checkout-session", createCheckoutSession)
	addr := cfg.DomainURL + ":" + cfg.Port
	log.Printf("Listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

// example request
// curl -X GET http://localhost:4242/create-checkout-session --data '[{"qty": 1, "foodName":"Chicken Tamales w/ Cheese"}, {"qty": 3, "foodName": "Chicken Tamales w/ Cheese"},{"qty": 1, "foodName": "Bistek Ranchero Taco"}]'
func createCheckoutSession(w http.ResponseWriter, r *http.Request) {
	var items []StripeBody
	json.NewDecoder(r.Body).Decode(&items)

	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}

	foods := map[string]string{
		"Bistek Ranchero Plate":     "price_1Piqw0GyTtiKluPRcBx9RlBp",
		"Carne Guisada Plate":       "price_1PiqvfGyTtiKluPRKscd58Wp",
		"Bistek Ranchero Taco":      "price_1PiqvMGyTtiKluPRyfvXboAm",
		"Carne Guisada Tacos":       "price_1Piqv6GyTtiKluPRPTN9L7B1",
		"Bean Tamales w/ Cheese":    "price_1PiqlCGyTtiKluPRzcr7VJNr",
		"Pork Tamales w/ Cheese":    "price_1PiqkuGyTtiKluPRzUuTmDYE",
		"Chicken Tamales w/ Cheese": "price_1PiqkRGyTtiKluPRhFO21Eql",
		"Bean Tamales":              "price_1Piqk6GyTtiKluPRDa1wDA51",
		"Pork Tamales":              "price_1PiqjuGyTtiKluPRQ3mKyJSC",
		"Chicken Tamales":           "price_1PimvtGyTtiKluPR0imgyXVG",
	}

	checkoutItems := []*stripe.CheckoutSessionLineItemParams{}

	for _, item := range items {
		price_id := foods[item.FoodName]
		fmt.Println("price_id: ", price_id)
		addItem := &stripe.CheckoutSessionLineItemParams{
			Price:    stripe.String(price_id), // the price ID has to be fetched from the stripe dashboard, so we have to create items beforehand
			Quantity: stripe.Int64(item.Qty),  //the quantity value can be passed as a body
		}

		checkoutItems = append(checkoutItems, addItem)
	}

	fmt.Println("items-list-count: ", len(checkoutItems))

	domain := cfg.HostUrl
	params := &stripe.CheckoutSessionParams{
		LineItems:                checkoutItems,
		AllowPromotionCodes:      stripe.Bool(true),
		BillingAddressCollection: stripe.String("auto"),
		Mode:                     stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL:               stripe.String(domain + "/success.html"), //redirect back to website???
		CancelURL:                stripe.String(domain + "/cancel.html"),  //redirect back to website???
		AutomaticTax: &stripe.CheckoutSessionAutomaticTaxParams{
			Enabled: stripe.Bool(true),
		},
		// CustomerEmail: stripe.String("ruben.valdez87@gmail.com"),
	}

	s, err := session.New(params)

	if err != nil {
		log.Printf("session.New: %v\n\n", err)
	}

	http.Redirect(w, r, s.URL, http.StatusSeeOther)
}
