package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.PersistentFlags().BoolVarP(&version, "version", "v", false, "Curlcraft version")
}

var version bool
var rootCmd = &cobra.Command{
	Use: "CurlCraft [URL] [METHOD] [FLAGS..]",
	Run: func(cmd *cobra.Command, args []string) {
		showVersion, _ := cmd.Flags().GetBool("version")

		if showVersion {
			fmt.Println("CurlCraft 0.0.1")
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
