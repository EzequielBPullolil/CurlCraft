package cmd

import (
	"errors"
	"fmt"
	"os"

	requestmanager "github.com/EzequielK-source/CurlCraft/internal/requestManager"
	"github.com/spf13/cobra"
)

var showVersion bool
var haveBodyData bool
var isComplex bool

func init() {
	rootCmd.PersistentFlags().BoolVarP(&showVersion, "version", "v", false, "Curlcraft version")
	rootCmd.PersistentFlags().BoolVarP(&isComplex, "complex", "c", false, "Allow complex request")
	rootCmd.PersistentFlags().BoolVarP(&haveBodyData, "bodyData", "d", false, "Allow make request with data")
}

var rootCmd = &cobra.Command{
	Use: "CurlCraft [URL] [METHOD/S] [CONTENT-TYPE] [FLAGS..]",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 && !showVersion {
			return errors.New("Requires at least URL arg")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		if showVersion {
			fmt.Println("CurlCraft 0.0.1")
			os.Exit(0)
		}
		requestManager := requestmanager.RequestManager(haveBodyData, isComplex)
		requestManager.MakeRequest(args)

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
