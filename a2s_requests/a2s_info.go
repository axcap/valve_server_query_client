package a2s_requests

import (
	"fmt"
)

var A2S_INFO_REQUEST = []byte{0xFF, 0xFF, 0xFF, 0xFF, 0x54, 0x53, 0x6F, 0x75, 0x72, 0x63, 0x65, 0x20, 0x45, 0x6E, 0x67, 0x69, 0x6E, 0x65, 0x20, 0x51, 0x75, 0x65, 0x72, 0x79, 0x00}

type A2S_INFO_RESPONSE struct {
	Header        byte
	Protocol      byte
	Name          string
	Map           string
	Folder        string
	Game          string
	ID            uint16
	Players       byte
	MaxPlayers    byte
	Bots          byte
	ServerType    byte
	Environment   byte
	Visibility    byte
	VAC           byte
	Version       string
	ExtraDataFlag byte
}

var VacValues = []string{"unsecured", "secured"}
var VisibilityValues = []string{"public", "private"}

func ParseInfoResponse(array []byte) A2S_INFO_RESPONSE {
	rv := A2S_INFO_RESPONSE{}

	i := 4
	rv.Header, i = array[i], i+1
	rv.Protocol, i = array[i], i+1
	rv.Name, i = getString(array, i)
	rv.Map, i = getString(array, i)
	rv.Folder, i = getString(array, i)
	rv.Game, i = getString(array, i)
	rv.ID, i = uint16(array[i+1])<<8|uint16(array[i]), i+2
	rv.Players, i = array[i], i+1
	rv.MaxPlayers, i = array[i], i+1
	rv.Bots, i = array[i], i+1
	rv.ServerType, i = array[i], i+1
	rv.Environment, i = array[i], i+1
	rv.Visibility, i = array[i], i+1
	rv.VAC, i = array[i], i+1
	rv.Version, i = getString(array, i)
	rv.ExtraDataFlag = array[i]
	return rv
}

func PrintInfoResponse(resp A2S_INFO_RESPONSE) {
	fmt.Printf("Header: %c %v\n", resp.Header, resp.Header == 'I')
	fmt.Printf("Protocol: %d\n", resp.Protocol)
	fmt.Printf("Name: %v\n", resp.Name)
	fmt.Printf("Map: %v\n", resp.Map)
	fmt.Printf("Folder: %v\n", resp.Folder)
	fmt.Printf("Game: %v\n", resp.Game)
	fmt.Printf("ID: %v\n", resp.ID)
	fmt.Printf("Players: %v\n", resp.Players)
	fmt.Printf("Max.Players: %v\n", resp.MaxPlayers)
	fmt.Printf("Bots: %v\n", resp.Bots)
	fmt.Printf("ServerType: %s\n", getServerType(resp.ServerType))
	fmt.Printf("Environment: %s\n", getEnvironment(resp.Environment))
	fmt.Printf("Visibility: %s\n", VisibilityValues[resp.Visibility])
	fmt.Printf("VAC: %s\n", VacValues[resp.VAC])
	fmt.Printf("Version: %s\n", resp.Version)
	fmt.Printf("EDF: %x\n", resp.ExtraDataFlag)
}

func getServerType(serverType byte) string {
	switch serverType {
	case 'd':
		return "dedicated"
	case 'l':
		return "non-dedicated"
	case 'p':
		return "SourceTv (proxy)"
	default:
		return "unknown"
	}
}

func getEnvironment(env byte) string {
	switch env {
	case 'l':
		return "Linux"
	case 'w':
		return "Windows"
	case 'm', 'o':
		return "Mac"
	default:
		return "unknown"
	}
}
