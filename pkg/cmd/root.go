package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "estore",
	Short: "CLI for managing estore",
}

func init() {
	createCmd.Flags().StringVarP(&customerName, "name", "n", "", "Name of the customer (required)")
	createCmd.MarkFlagRequired("name")
	RootCmd.AddCommand(createCmd)
	RootCmd.AddCommand(getCustomerCmd)
}
