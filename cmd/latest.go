/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// latestCmd represents the latest command
var latestCmd = &cobra.Command{
	Use:   "latest",
	Short: "Fetches and installs the latest GE release",
	Long: `Fetches and installs the latest GE release

Example:
	getui latest -proton`,
	Run: func(cmd *cobra.Command, args []string) {

		// Set Fags
		protonFlag, err := cmd.Flags().GetBool("proton")
		cobra.CheckErr(err)

		wineFlag, err := cmd.Flags().GetBool("wine")
		cobra.CheckErr(err)

		yesFlag, err := cmd.Flags().GetBool("yes")
		cobra.CheckErr(err)

		infoFlag, err := cmd.Flags().GetBool("yes")
		cobra.CheckErr(err)

		if protonFlag {
			fmt.Println("proton")
		}

		if wineFlag {
			fmt.Println("wine")
		}

		// if yesFlag {
		// 	fmt.Println("yes")
		// }

		// if infoFlag {
		// 	fmt.Println("info")
		// }

		if !protonFlag && !wineFlag {
			fmt.Println("proton and wine")
		}
	},
}

func init() {
	rootCmd.AddCommand(latestCmd)
	latestCmd.Flags().BoolP("proton", "p", false, "Installs GE-Proton")
	latestCmd.Flags().BoolP("wine", "w", false, "Installs GE-Wine")
	latestCmd.Flags().BoolP("yes", "y", false, "Skips user confirmation")
	latestCmd.Flags().BoolP("info", "i", false, "Displays info only")
}
