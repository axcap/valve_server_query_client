package cmd

import (
	"github.com/spf13/cobra"
	"vmon/a2s_requests"
)

func init() {
	var infoCmd = &cobra.Command{
		Use:   "info <server-addr>",
		Args:  cobra.ExactArgs(1),
		Short: "Get server information",
		Long:  `Retrieves information about the server including, but not limited to: its name, the map currently being played, and the number of players.`,
		Run: func(cmd *cobra.Command, args []string) {
			server := args[0]

			response := a2s_requests.GetBytes(server, a2s_requests.A2S_INFO_REQUEST)
			ptr := a2s_requests.ParseInfoResponse(response)
			a2s_requests.PrintInfoResponse(ptr)
		},
	}
	rootCmd.AddCommand(infoCmd)
}
