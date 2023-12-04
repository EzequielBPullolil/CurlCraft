package cmd

import (
	"fmt"
	"os"

	argumentManager "github.com/EzequielK-source/CurlCraft/internal/argumentManager"
	basicRequest "github.com/EzequielK-source/CurlCraft/internal/basicRequest"
	complexRequest "github.com/EzequielK-source/CurlCraft/internal/complexRequest"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.PersistentFlags().BoolVarP(&showVersion, "version", "v", false, "Curlcraft version")
	rootCmd.PersistentFlags().BoolVarP(&isComplex, "complex", "c", false, "Allow complex request")
}

var showVersion bool
var isComplex bool
var rootCmd = &cobra.Command{
	Use: "CurlCraft [URL] [METHOD/S] [FLAGS..]",
	Run: func(cmd *cobra.Command, args []string) {
		if showVersion {
			fmt.Println("CurlCraft 0.0.1")
			os.Exit(0)
		}
		url := argumentManager.Url(args)
		if isComplex {
			complexRequest.Request(url, argumentManager.Methods(args))
		} else {
			basicRequest.Request(url, argumentManager.Method(args))
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
