package main

const (
	CONN_HOST = "localhost"
	CONN_PORT = "42069"
	CONN_TYPE = "tcp"
)

type ControlData struct {
	Key string `json:"key"`
	Value uint8 `json:"value"`
}

type ServerStatus struct {
	Success bool `json:"success"`
}

func main() {
	ch := make(chan ControlData, 5)
	hasData := make(chan bool)
	go udpServer(ch, hasData)
	leds(ch, hasData)
}