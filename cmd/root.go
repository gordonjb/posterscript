package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "posterscript",
	Short: "Plex related scripts",
	Args:  cobra.NoArgs,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(version string, date string, commit string) {
	rootCmd.Version = version
	rootCmd.Long = fmt.Sprintf(`Plex related scripts. 
Version: %s built from #: %s on: %s`, version, commit, date)
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
