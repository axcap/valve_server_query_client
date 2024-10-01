package a2s_requests

import (
	"encoding/binary"
	"fmt"
	"log"
	"os"
	"text/tabwriter"
)

var A2S_RULES_REQUEST = []byte{0xFF, 0xFF, 0xFF, 0xFF, 'V', 0xFF, 0xFF, 0xFF, 0xFF}

const _A2S_RULES_HEADER = 'E'

type A2S_RULE struct {
	Name  string
	Value string
}

type A2S_RULES_RESPONSE struct {
	Header   byte
	NumRules uint16
	Rules    []A2S_RULE
}

func ParseRuleResponse(array []byte) A2S_RULES_RESPONSE {
	rv := A2S_RULES_RESPONSE{}

	i := 4
	rv.Header, i = array[i], i+1
	if rv.Header != _A2S_RULES_HEADER {
		log.Fatalf("Unrecognized header: '%c' (%02X)\n", rv.Header, rv.Header)
	}
	rv.NumRules, i = binary.LittleEndian.Uint16(array[i:i+2]), i+2

	for j := 0; j < int(rv.NumRules); j += 1 {
		rule := A2S_RULE{}
		rule.Name, i = getString(array, i)
		rule.Value, i = getString(array, i)
		rv.Rules = append(rv.Rules, rule)
	}

	return rv
}

func PrintRulesResponse(resp A2S_RULES_RESPONSE) {
	// fmt.Printf("Header: %c %v\n", resp.Header, resp.Header == _A2S_RULES_HEADER)
	// fmt.Printf("Rules num: %v\n", resp.NumRules)
	// fmt.Printf("Rules num: %v\n", len(resp.Rules))

	if int(resp.NumRules) < len(resp.Rules) {
		log.Fatalln("Rule count mismatch")
	}

	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	fmt.Fprintln(w, "Name\tValue")
	for _, rule := range resp.Rules {
		fmt.Fprintf(w, "%v\t%v\n", rule.Name, rule.Value)
	}
	w.Flush()
}
