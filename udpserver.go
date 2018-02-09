package main

import (
	"encoding/json"
	"os"
	"fmt"
	"net"
)

func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func udpServer(ch chan<- ControlData) {
	addr, err := net.ResolveUDPAddr("udp", ":" + CONN_PORT)
	checkError(err)

	conn, err := net.ListenUDP("udp", addr)
	checkError(err)
	fmt.Println("Listening on port " + CONN_PORT)

	defer conn.Close()

	buf := make([]byte, 1024)
	
	for  {
		n, addr, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("Error: " + err.Error())
		}

		var inObj ControlData
		err = json.Unmarshal(buf[:n], &inObj)

		if err != nil {
			fmt.Println("Error: " + err.Error())
		}
		fmt.Println("Raw data: ", string(buf[:n]))
		fmt.Println("Recieved ", inObj, " from ", addr)

		ch <- inObj
	}
}