package main

const (
	CONN_HOST = "localhost"
	CONN_PORT = "42069"
	CONN_TYPE = "tcp"
)

type ControlData struct {
	Key string `json:"key"`
	Value int `json:"value"`
}

type ServerStatus struct {
	Success bool `json:"success"`
}

func main() {
	ch := make(chan ControlData)
	go udpServer(ch)
	leds(ch)
}