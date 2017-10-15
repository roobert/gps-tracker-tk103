package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"regexp"
	"time"
)

var logger = log.New(os.Stdout, "", 0)

const (
	CONN_HOST = "0.0.0.0"
	CONN_PORT = "9000"
	CONN_TYPE = "tcp"
)

func main() {
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)

	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}

	defer l.Close()

	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)

	for {
		conn, err := l.Accept()

		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)

	for {
		buf = make([]byte, 1024)
		len, err := conn.Read(buf)

		// match EOF
		if len == 0 {
			conn.Close()
			Log("<- EOF: closing connection")
			break
		}

		if err != nil {
			fmt.Println("Error:", buf, len)
			fmt.Println("Error reading:", err.Error())
			panic("1")
			continue
		}

		// match handshake
		matched, err := regexp.Match("^##,imei:[0-9]{15},.;", buf)

		if err != nil {
			fmt.Println("Error:", buf)
			fmt.Println("Error reading:", err.Error())
			panic("2")
			continue
		}

		if matched == true {
			Log(fmt.Sprintln("<- (handshake)", string(buf)))
			Log("-> (handshake) LOAD\\r\\n")
			conn.Write([]byte("LOAD\r\n"))
			continue
		}

		// match data
		matched, err = regexp.Match("^imei:[0-9]{15},.*$", buf)

		if err != nil {
			fmt.Println("Error:", buf)
			fmt.Println("Error reading:", err.Error())
			panic("3")
			continue
		}

		if matched == true {
			Log(fmt.Sprintln("<- (data)", string(buf)))
			continue
		}

		// match ping
		matched, err = regexp.Match("^[0-9]{15}", buf)

		if err != nil {
			fmt.Println("Error:", buf)
			fmt.Println("Error reading:", err.Error())
			panic("4")
			continue
		}

		if matched == true {
			Log(fmt.Sprintln("<- (ping)", string(buf)))
			Log("-> (pong) OK\\r\\n")
			conn.Write([]byte("OK\r\n"))
			continue
		}

		Log("no match")
	}
}

func Log(msg string) {
	logger.SetPrefix(time.Now().Format("2006-01-02 15:04:05") + " ")
	logger.Print(msg)
}
