package cmd

import (
	"fmt"
	"os"

	internal "github.com/EzequielK-source/CurlCraft/internal"
	basicRequest "github.com/EzequielK-source/CurlCraft/internal/basicRequest"
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
		url := args[0]
		if showVersion {
			fmt.Println("CurlCraft 0.0.1")
			os.Exit(0)
		}

		if isComplex {
			fmt.Println("Complex request")
			for _, v := range args[1:] {
				if internal.IsHttpMethod(v) {
					basicRequest.Request(url, v)
				}
			}
		} else {
			if len(args) > 1 {
				method := args[1]
				basicRequest.Request(url, method)
			} else {
				basicRequest.Request(url, "get")
			}

		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
