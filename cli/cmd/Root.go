package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
)

func Execute() {
	RootCmd.Execute()
}

var RootCmd = &cobra.Command{
	Use:   "nais",
	Short: "Nais is the CLI for the Nais PAAS",
	Long: "Nais is the CLI for the Nais PAAS",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Velkommen til enn skikkelig nais plattform!!")
	},
}