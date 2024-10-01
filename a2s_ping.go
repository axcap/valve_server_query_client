package main

import (
	"fmt"
	"log"
)

var A2S_PING_REQUEST = []byte{0xFF, 0xFF, 0xFF, 0xFF, 'i'}

const _PING_HEADER = 'j'

type A2S_PING_RESPONSE struct {
	Header  byte
	Payload string
}

func parsePingResponse(array []byte) A2S_PING_RESPONSE {
	rv := A2S_PING_RESPONSE{}

	i := 4
	rv.Header, i = array[i], i+1
	if rv.Header != _PING_HEADER {
		log.Fatalf("Unrecognized header: '%c' (%02X), Expected: '%c' (%02X)\n", rv.Header, rv.Header, _PING_HEADER, _PING_HEADER)
	}
	rv.Payload, i = getString(array, i)

	return rv

}

func printPingResponse(resp A2S_PING_RESPONSE) {
	fmt.Printf("Header: %c %v\n", resp.Header, resp.Header == PLAYER_HEADER)
	fmt.Printf("Payload: %v\n", resp.Payload)
}
