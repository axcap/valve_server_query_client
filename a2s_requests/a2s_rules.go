package a2s_requests

import (
	"fmt"
	"log"
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

func parseRuleResponse(array []byte) A2S_RULES_RESPONSE {
	log.Println("PARSE PARSE PARSE")
	rv := A2S_RULES_RESPONSE{}

	i := 4
	rv.Header, i = array[i], i+1
	if rv.Header != _A2S_RULES_HEADER {
		log.Fatalf("Unrecognized header: '%c' (%02X)\n", rv.Header, rv.Header)
	}
	rv.NumRules, i = uint16(array[i])|uint16(array[i+1])<<8, i+2
	fmt.Printf("Numver of rules: %v\n", rv.NumRules)

	//fmt.Printf("Array len: %v\n", len(array))

	for j := 0; j < int(rv.NumRules); j += 1 {
		rule := A2S_RULE{}
		rule.Name, i = getString(array, i)
		if i == -1 || i > len(array) {
			log.Println("Breaking1")
			break
		}
		rule.Value, i = getString(array, i)
		if i == -1 || i > len(array) {
			log.Println(i)
			log.Println(rule.Name)
			//log.Fatalln("Breaking2")
			//break
		}
		rv.Rules = append(rv.Rules, rule)
	}

	return rv
}

func printRulesResponse(resp A2S_RULES_RESPONSE) {
	fmt.Printf("Header: %c %v\n", resp.Header, resp.Header == _A2S_RULES_HEADER)
	fmt.Printf("Rules num: %v\n", resp.NumRules)
	fmt.Printf("Rules num: %v\n", len(resp.Rules))

	if int(resp.NumRules) < len(resp.Rules) {
		log.Fatalln("Rule count mismatch")
	}

	for _, rule := range resp.Rules {
		log.Printf("%v: %v\n", rule.Name, rule.Value)
	}
}
