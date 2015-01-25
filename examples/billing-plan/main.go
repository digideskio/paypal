package main

import (
	_ "encoding/json"
	_ "fmt"
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

	amount := &paypal.Currency{Value: "100", Currency: "USD"}
	setupFee := &paypal.Currency{Value: "1", Currency: "USD"}
	taxAmount := &paypal.Currency{Value: "12.00", Currency: "USD"}

	chargeModels := []*paypal.ChargeModel{
		&paypal.ChargeModel{
			Type:   "TAX",
			Amount: taxAmount,
		},
	}
	paymentDefinitions := []*paypal.PaymentDefinition{
		&paypal.PaymentDefinition{
			Name:              "Monthly Payments",
			Type:              "REGULAR",
			Frequency:         "MONTH",
			FrequencyInterval: "2",
			Amount:            amount,
			Cycles:            "12",
			ChargeModels:      chargeModels,
		},
	}

	merchantPreferences := &paypal.MerchantPreferences{
		SetupFee:                setupFee,
		ReturnURL:               "https://cas.easyrtc.com/api/return",
		CancelURL:               "https://cas.easyrtc.com/api/cancel",
		AutoBillAmount:          "YES",
		InitialFailAmountAction: "CONTINUE",
		MaxFailAttempts:         "0",
	}

	billingPlan := &paypal.BillingPlan{
		Name:                "White Label Tawk Monthly",
		Description:         "White Label Tawk Monthly",
		Type:                "fixed",
		State:               "ACTIVE",
		PaymentDefinitions:  paymentDefinitions,
		MerchantPreferences: merchantPreferences,
	}

	// Delete all existing billing plans in CREATED state
	plans := getPlans(client, map[string]string{"status": "ACTIVE"})
	removePlans(client, plans)

	// Create a billing plan, initial state ACTIVE
	plan, err := client.CreateBillingPlan(billingPlan)
	if err != nil {
		log.Println(err)
	}

	// Get information on billing plan
	resp, err := client.GetBillingPlan(plan.ID)
	if err != nil {
		log.Println(err)
	}
	log.Println(resp)
}

func getPlans(client *paypal.Client, filter map[string]string) []paypal.BillingPlan {
	plans, err := client.ListBillingPlans(filter)
	if err != nil {
		log.Println(err)
	}
	return plans
}

func removePlans(client *paypal.Client, plans []paypal.BillingPlan) {
	patch := []paypal.PatchRequest{
		paypal.PatchRequest{"/", "replace", map[string]string{"state": "CREATED"}, ""},
	}
	for _, p := range plans {
		if err := client.UpdateBillingPlan(p.ID, patch); err != nil {
			log.Println(err)
		}
	}
}
