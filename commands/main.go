// Copyright 2014 The Priologic Authors. All rights reserved.
package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/golang/glog"
	"github.com/rmorriso/paypal"
)

const (
	ndays = 1
)

var (
	configFile                                  string
	fs1, fs2, fs3, fs4, fs5, fs6, fs7, fs8, fs9 *flag.FlagSet
	config                                      *Config
	all                                         bool
	token                                       string
	clientID                                    string
	secret                                      string
	client                                      *paypal.Client
	now                                         = time.Now()
)

// initialize command parameters
func init() {
	// global options
	flag.StringVar(&configFile, "f", "./sandbox.yaml", "the Paypal config file (default is to use Sandbox config)")
	// forecasts command options (locations has no options)
	fs1 = flag.NewFlagSet("billing-agreement", flag.ContinueOnError)
	fs1.StringVar(&token, "payment-token", "", "execute a billing agreement for the supplied payment token")
	fs1.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s options billing-agreement [create | update | retrieve | suspend | reactivate | cancel | set-balance | bill-balance ] sub-options\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Options:\n")
		fmt.Fprintf(os.Stderr, "\nbilling-agreement subcommand options:\n")
		fs1.PrintDefaults()
	}
	fs2 = flag.NewFlagSet("billing-plans", flag.ContinueOnError)
	fs2.BoolVar(&all, "all", false, "list all billing plans (default is to list only active plans)")
	fs2.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s options billing-plans [create | update | retrieve | list ] \n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Options:\n")
		fmt.Fprintf(os.Stderr, "\nbilling-agreement subcommand options:\n")
		fs2.PrintDefaults()
	}
	fs3 = flag.NewFlagSet("captures", flag.ContinueOnError)
	fs3.BoolVar(&all, "all", false, "list all capture resources (default is to list only approved and pending resources)")
	fs4 = flag.NewFlagSet("invoices", flag.ContinueOnError)
	fs4.BoolVar(&all, "all", false, "list all invoice resources (default is to list only approved and pending resources)")
	fs5 = flag.NewFlagSet("payments", flag.ContinueOnError)
	fs5.BoolVar(&all, "all", false, "list all payment resources (default is to list only approved and pending resources)")
	fs5.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s options payments [create | update | execute | get | list ] \n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Options:\n")
		fmt.Fprintf(os.Stderr, "\npayments subcommand options:\n")
		fs5.PrintDefaults()
	}
	fs6 = flag.NewFlagSet("payouts", flag.ContinueOnError)
	fs6.BoolVar(&all, "all", false, "list all payout resources (default is to list only approved and pending resources)")
	fs7 = flag.NewFlagSet("refunds", flag.ContinueOnError)
	fs7.BoolVar(&all, "all", false, "list all refund resources (default is to list only approved and pending resources)")
	fs8 = flag.NewFlagSet("sale", flag.ContinueOnError)
	fs8.BoolVar(&all, "all", false, "list all sale resources (default is to list only approved and pending resources)")
	fs9 = flag.NewFlagSet("transactions", flag.ContinueOnError)
	fs9.BoolVar(&all, "all", false, "list all transaction resources (default is to list only approved and pending resources)")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s options [ billing-agreement | billing-plans | capture | invoicing | payments | payouts | refund | sale | transactions ] subcommand options\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nbilling-agreement subcommand options:\n")
		fs1.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nbilling-plans subcommand options:\n")
		fs2.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\ncaptures subcommand options:\n")
		fs3.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\ninvoices subcommand options:\n")
		fs4.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\npayments subcommand options:\n")
		fs5.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\npayouts subcommand options:\n")
		fs6.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nrefunds subcommand coptions:\n")
		fs7.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nsale subcommand options:\n")
		fs8.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\ntransactions subcommand options:\n")
		fs9.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nNote: use -v 5 for debug output\n\n")
	}
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
	defer glog.Flush()

	flag.Parse()

	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		glog.Fatalf("Paypal config file: %s\n", err)
	}

	args := flag.Args()
	glog.V(5).Infof("Args: %v\n", args)

	var err error
	config, err = Init(configFile)
	if err != nil {
		glog.Fatalf("Error in configuration: %s\n", err)
	}

	// must provide a subcommand
	if len(args) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	// note: flag.Parse() must be called before subcommand processing to ensure
	// global options are consumed first from arguments array

	cmd := args[0]
	switch cmd {
	case "billing-agreement":
		// add support for create, execute, update, retrieve, suspend, cancel, reactivate
		if len(args) < 2 {
			flag.Usage()
			fs1.Usage()
			return
		}
		subcmd := args[1]
		if subcmd != "create" {
			fs1.Usage()
			return
		}
		// get billing-agreement subcommand options
		if err := fs1.Parse(args[2:]); err != nil {
			fs1.Usage()
			return
		}
	case "billing-plans":
		// get billing-plans subcommand options
		if err := fs1.Parse(args[1:]); err != nil {
			os.Exit(1)
		}
		plans, err := client.ListBillingPlans(map[string]string{
			"count": "10",
		})
		if err != nil {
			glog.Fatal("Could not retrieve billing plans: ", err)
		}
		for p := range plans {
			fmt.Printf("%v\n", p)
		}
	case "payments":
		// get payments subcommand options
		if err := fs5.Parse(args[1:]); err != nil {
			os.Exit(1)
		}
		payments, err := client.ListPayments(map[string]string{
			"count":   "10",
			"sort_by": "create_time",
		})
		if err != nil {
			glog.Fatal("Could not retrieve payments: ", err)
		}
		for p := range payments {
			fmt.Printf("%v\n", p)
		}
	case "sale":
		// get sale subcommand options
		if err := fs2.Parse(args[1:]); err != nil {
			os.Exit(1)
		}
		transactionID := "2W932963DR146200D"
		sale, err := client.GetSale(transactionID)
		if err != nil {
			glog.Fatal("Could not get sale %s: %s", transactionID, err)
		}
		fmt.Printf("%v\n", sale)
	default:
		flag.Usage()
		os.Exit(1)
	}

}
