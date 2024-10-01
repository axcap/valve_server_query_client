package cmd

import (
	"cs_monitor/a2s_requests"
	"github.com/spf13/cobra"
)

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Get server information",
	Long:  `Retrieves information about the server including, but not limited to: its name, the map currently being played, and the number of players.`,
	Run: func(cmd *cobra.Command, args []string) {
		server, err := cmd.Flags().GetString("server")
		if err != nil {
			panic(err)
		}

		response := a2s_requests.GetBytes(server, a2s_requests.A2S_INFO_REQUEST)
		ptr := a2s_requests.ParseInfoResponse(response)
		a2s_requests.PrintInfoResponse(ptr)
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// infoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// infoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	infoCmd.Flags().StringP("server", "s", "", "set server")
	infoCmd.MarkFlagRequired("server")
}
