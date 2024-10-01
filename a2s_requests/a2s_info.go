package a2s_requests

import (
	"encoding/binary"
	"fmt"
)

var A2S_INFO_REQUEST = []byte{0xFF, 0xFF, 0xFF, 0xFF, 0x54, 0x53, 0x6F, 0x75, 0x72, 0x63, 0x65, 0x20, 0x45, 0x6E, 0x67, 0x69, 0x6E, 0x65, 0x20, 0x51, 0x75, 0x65, 0x72, 0x79, 0x00}

type A2S_INFO_RESPONSE struct {
	Header        byte
	Protocol      byte
	Address       string // Obsolete GoldSource field
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
	Mod           byte // Obsolete GoldSource field
	VAC           byte
	Version       string
	ExtraDataFlag byte

	Link              string // Obsolete GoldSource field
	DownloadLink      string // Obsolete GoldSource field
	NULL              byte   // Obsolete GoldSource field
	GoldSourceVersion uint64 // Obsolete GoldSource field
	Size              uint64 // Obsolete GoldSource field
	Type              byte   // Obsolete GoldSource field
	DLL               byte   // Obsolete GoldSource field
}

var vacValues = []string{"unsecured", "secured"}
var visibilityValues = []string{"public", "private"}
var serverType = map[byte]string{
	'd': "dedicated",
	'l': "non-dedicated",
	'p': "SourceTv (proxy)",
}

var hostOS = map[byte]string{
	'l': "Linux",
	'w': "Windows",
	'm': "Mac",
	'o': "Mac",
}

var mods = map[byte]string{
	0: "Half-Life",
	1: "Half-Life mod",
}

func ParseInfoResponse(array []byte) A2S_INFO_RESPONSE {
	rv := A2S_INFO_RESPONSE{}

	i := 4
	rv.Header, i = array[i], i+1

	if IsGoldSourceServer(rv.Header) {
		parseInfoResponse_GoldSourceResponse(array, &rv, i)
	} else {
		_parseInfoResponse(array, &rv, i)
	}

	return rv
}

func _parseInfoResponse(array []byte, rv *A2S_INFO_RESPONSE, i int) {
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
}

func parseInfoResponse_GoldSourceResponse(array []byte, rv *A2S_INFO_RESPONSE, i int) {
	rv.Address, i = getString(array, i)
	rv.Name, i = getString(array, i)
	rv.Map, i = getString(array, i)
	rv.Folder, i = getString(array, i)
	rv.Game, i = getString(array, i)
	rv.Players, i = array[i], i+1
	rv.MaxPlayers, i = array[i], i+1
	rv.Protocol, i = array[i], i+1
	rv.ServerType, i = array[i], i+1
	rv.Environment, i = array[i], i+1
	rv.Visibility, i = array[i], i+1
	rv.Mod, i = array[i], i+1

	if rv.Mod == 1 {
		rv.Link, i = getString(array, i)
		rv.DownloadLink, i = getString(array, i)
		rv.NULL, i = array[i], i+1
		rv.GoldSourceVersion, i = binary.LittleEndian.Uint64(array[i:i+4]), i+4
		rv.Size, i = binary.LittleEndian.Uint64(array[i:i+4]), i+4
		rv.Type, i = array[i], i+1
		rv.DLL, i = array[i], i+1
	}
	rv.VAC, i = array[i], i+1
	rv.Bots, i = array[i], i+1
}

func PrintInfoResponse(resp A2S_INFO_RESPONSE) {
	fmt.Printf("Header: %c %v\n", resp.Header, resp.Header == 'I' || IsGoldSourceServer(resp.Header))
	if IsGoldSourceServer(resp.Header) {
		fmt.Printf("Address: %v\n", resp.Address)
	}

	fmt.Printf("Name: %v\n", resp.Name)
	fmt.Printf("Protocol: %d\n", resp.Protocol)
	fmt.Printf("Map: %v\n", resp.Map)
	fmt.Printf("Folder: %v\n", resp.Folder)
	fmt.Printf("Game: %v\n", resp.Game)

	if resp.ID != 0 {
		fmt.Printf("ID: %v\n", resp.ID)
	}

	fmt.Printf("Players: %v\n", resp.Players)
	fmt.Printf("Max.Players: %v\n", resp.MaxPlayers)
	fmt.Printf("Bots: %v\n", resp.Bots)
	fmt.Printf("ServerType: %s\n", serverType[resp.ServerType])
	fmt.Printf("Environment: %s\n", hostOS[resp.Environment])
	fmt.Printf("Visibility: %s\n", visibilityValues[resp.Visibility])
	fmt.Printf("Mod: %s\n", mods[resp.Mod])

	if resp.Mod == 1 {
		fmt.Printf("Link: %v\n", resp.Link)
		fmt.Printf("DownloadLink: %v\n", resp.DownloadLink)
		fmt.Printf("Version: %v\n", resp.GoldSourceVersion)
		fmt.Printf("Size: %v\n", resp.Size)
		fmt.Printf("Type: %v\n", resp.Type)
		fmt.Printf("DLL: %v\n", resp.DLL)
	}

	if !IsGoldSourceServer(resp.Header) {
		fmt.Printf("Version: %s\n", resp.Version)
		fmt.Printf("ExtraDataFlag: %x\n", resp.ExtraDataFlag)
	}

	fmt.Printf("VAC: %s\n", vacValues[resp.VAC])
}

func IsGoldSourceServer(header byte) bool {
	return header == 'm'
}
