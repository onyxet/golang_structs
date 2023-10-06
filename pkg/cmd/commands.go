package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"structs/pkg/first"
)

var customerName string
var customers first.Customers

var createCmd = &cobra.Command{
	Use:   "createCustomer",
	Short: "Create a new customer",
	Run: func(cmd *cobra.Command, args []string) {
		if _, err := os.Stat("estore.json"); err == nil {
			fileContent, err := os.ReadFile("estore.json")
			if err != nil {
				panic(err)
			}

			if err := json.Unmarshal(fileContent, &customers); err != nil {
				panic(err)
			}
		}
		if customerName == "" {
			panic("Customer name is required")
		}
		customer := first.Customer{
			Name: customerName,
		}
		customers.AddCustomer(customer)
		toFile, _ := json.MarshalIndent(customers, "", " ")
		err := os.WriteFile("estore.json", toFile, 0644)
		if err != nil {
			panic(err)
		}
	},
}

var getCustomerCmd = &cobra.Command{
	Use:   "getCustomers",
	Short: "List existing customers",
	Run: func(cmd *cobra.Command, args []string) {
		if _, err := os.Stat("estore.json"); err == nil {
			fileContent, err := os.ReadFile("estore.json")
			if err != nil {
				panic(err)
			}

			if err := json.Unmarshal(fileContent, &customers); err != nil {
				panic(err)
			}
		}
		fmt.Println(customers)
	},
}

var BuySomethingCmd = &cobra.Command{
	Use:   "buySomething",
	Short: "Randomly buy something",
	Run: func(cmd *cobra.Command, args []string) {
		items := make(chan first.Items)
		go first.GenerateRandomOrders(items)
	},
}
