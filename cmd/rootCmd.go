package cmd

import (
	"errors"
	"fmt"
	"os"

	requestmanager "github.com/EzequielK-source/CurlCraft/internal/requestManager"
	"github.com/spf13/cobra"
)

var showVersion bool
var bodyData string
var isComplex bool

func init() {
	rootCmd.PersistentFlags().BoolVarP(&showVersion, "version", "v", false, "Curlcraft version")
	rootCmd.PersistentFlags().BoolVarP(&isComplex, "complex", "c", false, "Allow complex request")
	rootCmd.PersistentFlags().StringVarP(&bodyData, "body", "d", "", "Allow make request with data")
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
		body, _ := cmd.Flags().GetString("body")
		if showVersion {
			fmt.Println("CurlCraft 0.0.1")
			os.Exit(0)
		}
		requestManager := requestmanager.RequestManager(body, isComplex, args)
		requestManager.MakeRequest()

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
