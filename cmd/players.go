package cmd

import (
	"cmp"
	"log/slog"
	"slices"
	"strings"
	"vmon/a2s_requests"

	"github.com/spf13/cobra"
)

var playersCmd = &cobra.Command{
	Use:   "players <server-addr>",
	Short: "List players",
	Long:  "This query retrieves information about the players currently on the server.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		server := args[0]

		slog.Debug("ListPlayers", "server-addr", server)
		responseBytes := a2s_requests.GetBytes(server, a2s_requests.A2S_PLAYER_REQUEST)
		response := a2s_requests.ParsePlayerResponse(responseBytes)

		sortBy, err := cmd.Flags().GetString("sort")
		if err != nil {
			slog.Error("Could not get sorting by value")
			panic(err)
		}
		sort(response, sortBy)

		a2s_requests.PrintPlayerResponse(response)
	},
}

func sort(resp a2s_requests.A2S_PLAYER_RESPONSE, sortBy string) {
	switch strings.ToLower(sortBy) {
	case "score":
		slices.SortFunc(resp.Players, cmpByScore)
	case "duration":
		slices.SortFunc(resp.Players, cmpByDuration)
	case "name":
		slices.SortFunc(resp.Players, cmpByName)
	default:
		slog.Error("Unknown sorting argument", "sortBy", sortBy)
		panic(sortBy)
	}
}

func cmpByScore(a, b a2s_requests.A2S_PLAYER) int {
	return cmp.Compare(b.Score, a.Score)
}
func cmpByDuration(a, b a2s_requests.A2S_PLAYER) int {
	return cmp.Compare(b.Duration, a.Duration)
}
func cmpByName(a, b a2s_requests.A2S_PLAYER) int {
	return cmp.Compare(b.Name, a.Name)
}

func init() {
	rootCmd.AddCommand(playersCmd)

	playersCmd.Flags().StringP("sort", "s", "score", "")
}
