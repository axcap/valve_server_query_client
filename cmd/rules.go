package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"vmon/a2s_requests"
)

func init() {
	var singleRule string

	var rulesCmd = &cobra.Command{
		Use:   "rules <server addr>",
		Short: "Print server rules/config",
		Long:  "Returns the server rules, or configuration variables in name/value pairs.",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			server := args[0]

			response_bytes := a2s_requests.GetBytes(server, a2s_requests.A2S_RULES_REQUEST)
			response := a2s_requests.ParseRuleResponse(response_bytes)

			if singleRule != "" {
				for _, rule := range response.Rules {
					if rule.Name == singleRule {
						fmt.Println(rule.Value)
						os.Exit(0)
					}
				}
				fmt.Println("Not found")
			} else {
				a2s_requests.PrintRulesResponse(response)
			}
		},
	}

	rootCmd.AddCommand(rulesCmd)
	rulesCmd.Flags().StringVarP(&singleRule, "rule", "r", "", "Only show single rule/config (e.g: mp_friendlyfire)")
}
