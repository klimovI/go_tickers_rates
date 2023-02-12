package cmd

import (
	"github.com/spf13/cobra"

	"github.com/klimovI/go_tickers_rates/server"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Starts server",
	Run: func(cmd *cobra.Command, args []string) {
		server.StartServer()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
