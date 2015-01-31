package main

import (
	"encoding/json"
	_ "fmt"
	"log"
	"os"
	_ "time"

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

	/*
			plan, err := client.ListBillingPlans(map[string]string{"status": "ACTIVE", "page_size": "1"})
			if err != nil {
				log.Println(err)
			}
			p := &paypal.BillingPlan{ID: plan[0].ID}

			now := time.Now()
			now = now.AddDate(0, 0, 1)
			payer := &paypal.Payer{PaymentMethod: paypal.PaymentMethodPaypal}

			agree := &paypal.BillingAgreement{
				Name:        "White Label Tawk",
				Description: "Your customized Tawk server",
				StartDate:   now.Format("2006-01-02T15:04:05Z"),
				Payer:       payer,
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

		patch := []paypal.PatchRequest{
			paypal.PatchRequest{
				Op:    "replace",
				Path:  "/",
				Value: map[string]string{"description": "Modified"},
			},
		}
	*/

	id := "I-3DLV8D44YSRY"

	/*
		if err := client.UpdateBillingAgreement(id, patch); err != nil {
			log.Println(err)
		}

	*/
	descr := &paypal.AgreementStateDescriptor{
		Note: "Suspended for test purposes.",
	}

	err := client.ReactivateBillingAgreement(id, descr)
	if err != nil {
		log.Println(err)
	}

	agree, err := client.GetBillingAgreement(id)
	if err != nil {
		log.Println(err)
	}

	a, _ := json.Marshal(agree)
	log.Println(string(a))
}
