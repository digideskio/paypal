package main

import (
	"log"
	"os"
	"time"

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

	token := ""
	if len(os.Args) == 2 {
		token = os.Args[1]
	}

	client := paypal.NewClient(clientID, secret, paypal.APIBaseSandBox)

	if token == "" {
		now := time.Now()
		now = now.AddDate(0, 0, 1)
		startDate := now.Format("2006-01-02T15:04:05Z")
		payer := paypal.Payer{PaymentMethod: paypal.PaymentMethodPaypal}

		p, err := client.ListBillingPlans(map[string]string{"status": "ACTIVE", "page_size": "1"})
		if err != nil {
			log.Println(err)
		}
		plan := &paypal.BillingPlan{ID: p[0].ID}

		agreement := &paypal.BillingAgreement{
			Name:        "White Label Tawk Monthly",
			Description: "Your own tawk server.",
			StartDate:   startDate,
			Payer:       payer,
			BillingPlan: plan,
		}

		agree, err := client.CreateBillingAgreement(agreement)
		if err != nil {
			log.Println(err)
		}

		log.Println(agree.BillingAgreement)
	} else {
		resp, err := client.ExecuteBillingAgreement(token)
		if err != nil {
			log.Println(err)
		}
		a, _ := client.GetBillingAgreement(resp.ID)
		log.Println("agreement: ", a)
	}
}
