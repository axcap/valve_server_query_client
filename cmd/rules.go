package cmd

import (
	"github.com/spf13/cobra"
	"vmon/a2s_requests"
)

// rulesCmd represents the rules command
var rulesCmd = &cobra.Command{
	Use:   "rules <server addr>",
	Short: "Print server rules/config",
	Long:  "Returns the server rules, or configuration variables in name/value pairs.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		server := args[0]

		response := a2s_requests.GetBytes(server, a2s_requests.A2S_RULES_REQUEST)
		ptr := a2s_requests.ParseRuleResponse(response)
		a2s_requests.PrintRulesResponse(ptr)
	},
}

func init() {
	rootCmd.AddCommand(rulesCmd)
}
