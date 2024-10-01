package a2s_requests

import (
	"fmt"
	"log"
	"net"
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

func GetBytes(server string, request []byte) []byte {
	raddr, err := net.ResolveUDPAddr("udp", server)
	if err != nil {
		log.Fatalln(err)
	}

	conn, err := net.DialUDP("udp", nil, raddr)
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

request:
	_, err = conn.Write(request)
	if err != nil {
		log.Fatalln(err)
	}

	response := readServerResponse(conn)

	if response[4] == 'A' {
		request = append(request[0:5], response[5:9]...)
		goto request
	}

	return response
}

func getString(array []byte, startIndex int) (string, int) {
	if startIndex >= len(array) {
		return "", -1
	}

	for i := startIndex; i < len(array); i += 1 {
		if array[i] == 0 {
			return string(array[startIndex:i]), i + 1
		}
	}

	return "", -1
}

func readServerResponse(conn *net.UDPConn) []byte {
	const maxMessageSize = 1400

	buff := make([]byte, maxMessageSize)
	_, err := conn.Read(buff)
	if err != nil {
		log.Fatalln(err)
	}

	if buff[0] == 0xFE {
		packet := buff[8]
		current_packet := (packet >> 4) + 2
		packet_max := packet & 0b1111

		for current_packet <= packet_max {
			tmp := make([]byte, maxMessageSize)
			n_read, err := conn.Read(tmp)
			if err != nil {
				log.Fatalln(err)
			}

			packet = tmp[4+4]
			current_packet = (packet >> 4) + 2
			buff = append(buff, tmp[9:n_read]...)
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
