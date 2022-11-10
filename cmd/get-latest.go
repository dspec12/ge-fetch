package cmd

import (
	"github.com/dspec12/getui/internal"
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

		infoFlag, err := cmd.Flags().GetBool("info")
		cobra.CheckErr(err)

		if protonFlag {
			latestProton(infoFlag, yesFlag)
		}

		if wineFlag {
			latestWine(infoFlag, yesFlag)
		}

		if !protonFlag && !wineFlag {
			latestProton(infoFlag, yesFlag)
			latestWine(infoFlag, yesFlag)
		}
	},
}

func init() {
	rootCmd.AddCommand(latestCmd)
	latestCmd.Flags().BoolP("proton", "p", false, "Installs GE-Proton")
	latestCmd.Flags().BoolP("wine", "w", false, "Installs GE-Wine")
	latestCmd.Flags().BoolP("yes", "y", false, "Skips user confirmation")
	latestCmd.Flags().BoolP("info", "i", false, "Displays info only")
	latestCmd.MarkFlagsMutuallyExclusive("yes", "info")
}

func latestProton(infoFlag, yesFlag bool) {
	protonReleases := internal.GetReleases(internal.UrlProtonGECustom)
	protonReleases[0].DisplayInfo()
	if !infoFlag {
		if yesFlag {
			protonReleases[0].Download(internal.ProtonInstallDir, false)
		} else {
			protonReleases[0].Download(internal.ProtonInstallDir, true)
		}
	}
}

func latestWine(infoFlag, yesFlag bool) {
	wineReleases := internal.GetReleases(internal.UrlWineGECustom)
	wineReleases[0].DisplayInfo()
	if !infoFlag {
		if yesFlag {
			wineReleases[0].Download(internal.WineInstallDir, false)
		} else {
			wineReleases[0].Download(internal.WineInstallDir, true)
		}
	}
}
