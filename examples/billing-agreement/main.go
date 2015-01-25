package main

import (
	"time"
   "encoding/json"
    "fmt"
    "log"
    "os"

    "github.com/rmorriso/paypal"
)

func main() {
    clientID := os.Getenv("PAYPAL_CLIENTID")
    if clientID == "" {
        panic("Paypal clientID is missing")
    }

    secret := os.Getenv("PAYPAL_SECRET")
    if secret == "" {
        panic("Paypal secret is missing")
    }

    client := paypal.NewClient(clientID, secret, paypal.APIBaseSandBox)

    plan, err := client.ListBillingPlans(map[string]string {"status":"ACTIVE", "page_size":"1"})
    if err != nil {
		log.Println(err)
	}
	p := &paypal.BillingPlan{ID: plan[0].ID}

    now := time.Now()
    now = now.AddDate(0, 0, 1)
    payer := &paypal.Payer{PaymentMethod: paypal.PaymentMethodPaypal}

    agree := &paypal.BillingAgreement{
		Name: "White Label Tawk",
		Description: "Your customized Tawk server",
		StartDate: now.Format("2006-01-02T15:04:05Z"),
		Payer: payer,
		BillingPlan: p,
    }
    a, err := json.Marshal(agree)
	if err != nil {
		log.Println(err)
	}
	log.Println(string(a))
 

    agreement, err := client.CreateBillingAgreement(agree)
	if err != nil {
		log.Println(err)
	}
	
	a, err = json.Marshal(agreement)
	if err != nil {
		log.Println(err)
	}
    fmt.Println(string(a))
}
