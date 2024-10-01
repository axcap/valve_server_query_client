package cmd

import (
	"github.com/spf13/cobra"
	"vmon/a2s_requests"
)

// playersCmd represents the players command
var playersCmd = &cobra.Command{
	Use:   "players <server-addr>",
	Short: "List players",
	Long:  "This query retrieves information about the players currently on the server.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		server := args[0]

		response := a2s_requests.GetBytes(server, a2s_requests.A2S_PLAYER_REQUEST)
		ptr := a2s_requests.ParsePlayerResponse(response)
		a2s_requests.PrintPlayerResponse(ptr)
	},
}

func init() {
	rootCmd.AddCommand(playersCmd)
}
