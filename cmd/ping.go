package cmd

import (
	"cs_monitor/a2s_requests"
	"fmt"
	"github.com/spf13/cobra"
	"time"
)

// pingCmd represents the ping command
var pingCmd = &cobra.Command{
	Use:   "ping <server-addr>",
	Short: "Get ping to the server",
	Args:  cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		server := args[0]

		start := time.Now()

		response := a2s_requests.GetBytes(server, a2s_requests.A2S_PING_REQUEST)
		a2s_requests.ParsePingResponse(response)

		elapsed := time.Since(start)
		fmt.Println(elapsed.Round(time.Millisecond))
	},
}

func init() {
	rootCmd.AddCommand(pingCmd)
}
