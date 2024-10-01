package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

/*
TODO:
* [x] A2S_INFO
* [x] Suport for challenge request
* [x] A2S_RULES
* [x] A2S_PLAYER
* [x] A2S_PING
*
* [] Source / GO  engine
*/

func main() {
	serverUrl := "45.144.155.163:27015"
	if len(os.Args) > 1 {
		serverUrl = os.Args[1]
	}

	raddr, err := net.ResolveUDPAddr("udp", serverUrl)
	if err != nil {
		log.Fatalln(err)
	}

	conn, err := net.DialUDP("udp", nil, raddr)
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	log.Printf("The UDP server is %s\n", conn.RemoteAddr().String())

	request := A2S_PING_REQUEST
request:
	log.Printf("Request len: %d\n", len(request))
	printHexArray("request", request)
	_, err = conn.Write(request)

	if err != nil {
		log.Fatalln(err)
	}

	response := readServerResponse(conn)
	err = os.WriteFile("buff.hex", response, 0644)
	printHexArray("Response: ", response)

	if response[4] == 'A' {
		log.Println("Challenge requested")
		request = append(A2S_PING_REQUEST[0:5], response[5:9]...)
		goto request
	}

	//printHexArray("Resp: ", response)
	ptr := parsePingResponse(response)
	//ptr := parsePlayerResponse(response)
	// log.Println(ptr)
	printPingResponse(ptr)
}

func readServerResponse(conn *net.UDPConn) []byte {
	const maxMessageSize = 1400 // max UDP message size, unless you know better

	buff := make([]byte, maxMessageSize)
	_, err := conn.Read(buff)
	if err != nil {
		log.Fatalln(err)
	}

	if buff[0] == 0xFE {
		//return buff[9:]
		log.Println("Multipacket")

		header := buff[:4]
		printHexArray("Header: ", header)

		id := buff[4:8]
		printHexArray("ID: ", id)
		packet := buff[8]
		printHexArray("Packet: ", buff[8:9])
		log.Printf("PacketNumber: %x\n", packet)
		current_packet := (packet >> 4) + 2
		packet_max := packet & 0b1111

		log.Printf("Packets: %d/%d\n", current_packet, packet_max)

		for current_packet <= packet_max {
			log.Printf("%d / %d \n", current_packet, packet_max)
			tmp := make([]byte, maxMessageSize)
			n_read, err := conn.Read(tmp)
			fmt.Printf("Read: %d\n", n_read)
			if err != nil {
				log.Fatalln(err)
			}

			printHexArray("TMP: ", tmp)
			packet = tmp[4+4]
			current_packet = (packet >> 4) + 2
			log.Printf("Current packet: %d\v", current_packet)
			buff = append(buff, tmp[9:n_read]...)
			err = os.WriteFile("buff.hex", buff[9:], 0644)
			//panic("")
		}
		return buff[9:]
	}

	return buff
}

func printHexArray(title string, array []byte) {
	str := ""
	str += fmt.Sprintf("%s: [", title)
	for i, char := range array {
		if i != len(array)-1 {
			str += fmt.Sprintf("%02X ", char)
		} else {
			str += fmt.Sprintf("%02X", char)
		}
	}
	str += fmt.Sprintln("]")

	log.Println(str)
}
