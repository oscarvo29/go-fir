package main

import "github.com/google/gopacket/pcap"

func GetPcapPacket() *pcap.Handle {

	handle, err := pcap.OpenLive("eth0", 1600, true, pcap.BlockForever)
	if err != nil {
		panic(err)
	}

	return handle
}
