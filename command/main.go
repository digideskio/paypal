// Copyright 2014 The Priologic Authors. All rights reserved.
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/rmorriso/paypal"
	"github.com/spf13/cobra"
)

const (
	ndays = 1
)

var (
	configFile                    string
	config                        *Config
	verbosity                     int
	cmdPaypal                     *cobra.Command
	cmdAuthorization              *cobra.Command
	cmdBillingAgreement           *cobra.Command
	subcmdBillingAgreementExecute *cobra.Command
	cmdBillingPlans               *cobra.Command
	cmdCapture                    *cobra.Command
	cmdInvoice                    *cobra.Command
	cmdOrder                      *cobra.Command
	cmdPayments                   *cobra.Command
	subcmdPaymentsCreate          *cobra.Command
	subcmdPaymentsGet             *cobra.Command
	subcmdPaymentsExecute         *cobra.Command
	subcmdPaymentsUpdate          *cobra.Command
	subcmdPaymentsList            *cobra.Command
	subcmdPaymentsListAll         *cobra.Command
	subcmdPaymentsListCreated     *cobra.Command
	subcmdPaymentsListFailed      *cobra.Command
	cmdPayout                     *cobra.Command
	cmdRefund                     *cobra.Command
	cmdSale                       *cobra.Command

	token    string
	clientID string
	secret   string
	client   *paypal.Client
	now      = time.Now()
)

// initialize command parameters
func init() {
	cmdPaypal = &cobra.Command{Use: "paypal"}
	cmdPaypal.PersistentFlags().IntVarP(&verbosity, "verbosity", "v", 0, "the verbosity level for logging messages")
	cmdPaypal.PersistentFlags().StringVarP(&configFile, "config", "c", "./sandbox.yaml", "the Paypal config file (default is to use Sandbox config)")

	// Billing Agreement commands
	cmdBillingAgreement = &cobra.Command{
		Use:   "billing-agreement",
		Short: "paypal billing agreement operations",
		Long: `a customer must execute a billing agreement for a given billing plan.
        `,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("replace with call to PayPal API")
			initConfig()
		},
	}
	subcmdBillingAgreementExecute = &cobra.Command{
		Use:   "execute",
		Short: "execute a billing agreement",
		Long: `a customer must execute a billing agreement for a given billing plan.
        `,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("replace with call to PayPal API")
		},
	}
	subcmdBillingAgreementExecute.Flags().StringVarP(&token, "payment-token", "t", "", "payment token for the billing agreement")

	cmdBillingAgreement.AddCommand(subcmdBillingAgreementExecute)

	// Billing Planss commands
	cmdBillingPlans = &cobra.Command{
		Use:   "billing-plans",
		Short: "paypal billing plan operations",
		Long: `a billing plan is used for subscription-based products.
        `,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("replace with call to PayPal API")
		},
	}

	// Payments commands
	cmdPayments = &cobra.Command{
		Use:   "payments",
		Short: "paypal payments operations",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("replace with call to PayPal API for payment list")
		},
	}
	// Payments List command
	subcmdPaymentsList = &cobra.Command{
		Use:   "list",
		Short: "list payments",
		Long: `list payments to merchant identified by API key.
        `,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("replace with call to PayPal API")
		},
	}
	subcmdPaymentsListAll = &cobra.Command{
		Use:   "all",
		Short: "list all payments",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("replace with call to PayPal API")
		},
	}
	subcmdPaymentsListCreated = &cobra.Command{
		Use:   "created",
		Short: "list payments in created state",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("replace with call to PayPal API")
		},
	}
	subcmdPaymentsListFailed = &cobra.Command{
		Use:   "failed",
		Short: "list payments in failed state",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("replace with call to PayPal API")
		},
	}
	subcmdPaymentsList.AddCommand(subcmdPaymentsListAll, subcmdPaymentsListCreated, subcmdPaymentsListFailed)

	// Payments Create command
	subcmdPaymentsCreate = &cobra.Command{
		Use:   "create",
		Short: "create payment",
		Long: `Create a payment, specifying the intent (sale, authorize, or order), and other
transactions details like amount, recipients, and items.
        `,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("replace with call to PayPal API")
		},
	}
	subcmdPaymentsCreate.Flags().StringVarP(&token, "file", "f", "", "read payment details from file (JSON)")
	subcmdPaymentsGet = &cobra.Command{
		Use:   "get",
		Short: "get payment",
		Long: `Use this call to get details about a payment by ID.
        `,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("replace with call to PayPal API")
		},
	}
	subcmdPaymentsGet.Flags().StringVarP(&token, "id", "i", "", "payment id to be executed")
	subcmdPaymentsExecute = &cobra.Command{
		Use:   "execute",
		Short: "execute payment",
		Long: `Use this call to execute (complete) a PayPal payment that has been approved by the payer.
               You can optionally update transaction information when executing the payment by passing
               in one or more transactions.
        `,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("replace with call to PayPal API")
		},
	}
	subcmdPaymentsExecute.Flags().StringVarP(&token, "id", "i", "", "payment id to be executed")
	subcmdPaymentsUpdate = &cobra.Command{
		Use:   "update",
		Short: "update payment details",
		Long: `Use this call to partially update the payment resource for the given identifier.
        `,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("replace with call to PayPal API")
		},
	}
	subcmdPaymentsUpdate.Flags().StringVarP(&token, "id", "i", "", "payment id to be executed")
	subcmdPaymentsUpdate.Flags().StringVarP(&token, "file", "f", "", "read payment details from file (JSON)")

	cmdPayments.AddCommand(subcmdPaymentsList, subcmdPaymentsCreate, subcmdPaymentsGet, subcmdPaymentsUpdate, subcmdPaymentsExecute)

	cmdPaypal.AddCommand(cmdBillingAgreement, cmdBillingPlans, cmdPayments)

	clientID = os.Getenv("PAYPAL_CLIENTID")
	if clientID == "" {
		panic("Paypal clientID is missing")
	}

	secret = os.Getenv("PAYPAL_SECRET")
	if secret == "" {
		panic("Paypal secret is missing")
	}
	client = paypal.NewClient(clientID, secret, paypal.APIBaseSandBox)
}

func main() {
	cmdPaypal.PersistentFlags().Parse(os.Args[1:])
	initConfig()
	cmdPaypal.Execute()

}

func initConfig() {
	config, err := Init(configFile)
	if err != nil {
		log.Fatalf("Error in configuration: %s\n", err)
	}
	if verbosity >= 5 {
		log.Printf("Read config from %s: %v\n", configFile, config)
	}
}
