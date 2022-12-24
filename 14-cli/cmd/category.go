/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// categoryCmd represents the category command
var categoryCmd = &cobra.Command{
	Use:   "category",
	Short: "Manage categories",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// name, _ := cmd.Flags().GetString("name")
		fmt.Println("category called with name: " + category)
		exists, _ := cmd.Flags().GetBool("exists")
		fmt.Println("category called with exists: " + fmt.Sprint(exists))
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("Chamado antes do Run")
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("Chamado depois do Run")
	},
}

var category string

func init() {
	rootCmd.AddCommand(categoryCmd)
	categoryCmd.PersistentFlags().StringVarP(&category, "name", "n", "Y", "Name of the category")
	categoryCmd.PersistentFlags().BoolP("exists", "e", false, "Check if category exists")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// categoryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// categoryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
