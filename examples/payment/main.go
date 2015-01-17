package main

import (
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

    payments, err := client.ListPayments(map[string]string{
        "count":   "10",
        "sort_by": "create_time",
    })
    if err != nil {
        log.Fatal("Could not retrieve payments: ", err)
    }

    fmt.Println(payments)
}
