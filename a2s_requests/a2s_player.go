package a2s_requests

import (
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"os"
	"text/tabwriter"
	"time"
)

var A2S_PLAYER_REQUEST = []byte{0xFF, 0xFF, 0xFF, 0xFF, 'U', 0xFF, 0xFF, 0xFF, 0xFF}

const PLAYER_HEADER = 'D'

type A2S_PLAYER struct {
	Index    byte
	Name     string
	Score    uint32
	Duration float32
}

type A2S_PLAYER_RESPONSE struct {
	Header     byte
	NumPlayers byte
	Players    []A2S_PLAYER
}

func ParsePlayerResponse(array []byte) A2S_PLAYER_RESPONSE {
	rv := A2S_PLAYER_RESPONSE{}

	i := 4
	rv.Header, i = array[i], i+1
	if rv.Header != PLAYER_HEADER {
		log.Fatalf("Unrecognized header: '%c' (%02X)\n", rv.Header, rv.Header)
	}
	rv.NumPlayers, i = array[i], i+1

	for j := 0; j < int(rv.NumPlayers); j += 1 {
		player := A2S_PLAYER{}

		player.Index, i = array[i], i+1
		player.Name, i = getString(array, i)
		if i == -1 || i > len(array) {
			log.Fatalln("Breaking1")
			break
		}

		player.Score, i = binary.LittleEndian.Uint32(array[i:i+4]), i+4

		scoreAsUint32 := binary.LittleEndian.Uint32(array[i : i+4])
		player.Duration, i = math.Float32frombits(scoreAsUint32), i+4

		rv.Players = append(rv.Players, player)
	}

	return rv

}

func PrintPlayerResponse(resp A2S_PLAYER_RESPONSE) {
	// fmt.Fprintf(w, "Header: %c %v\n", resp.Header, resp.Header == PLAYER_HEADER)
	// fmt.Fprintf(w, "NumPlayers: %v\n", resp.NumPlayers)
	// fmt.Fprintf(w, "Len(Players): %v\n", len(resp.Players))

	// if int(resp.NumPlayers) > len(resp.Players) {
	// 	fmt.Println("Someone is connection to the server")
	// }

	if int(resp.NumPlayers) < len(resp.Players) {
		log.Fatalln("Player count mismatch")
	}

	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	fmt.Fprintln(w, "Name\tScore\tDuration")
	for _, player := range resp.Players {
		durr := fmt.Sprintf("%.0fs", player.Duration)
		duration, err := time.ParseDuration(durr)
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(w, "%s\t%d\t%v\n", player.Name, player.Score, duration)
	}
	w.Flush()
}
