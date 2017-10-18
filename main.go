package main

import (
	"fmt"
	. "github.com/roobert/gps-tracker-tk103/db"
	. "github.com/roobert/gps-tracker-tk103/error"
	"log"
	"net"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var logger = log.New(os.Stdout, "", 0)

const (
	CONN_HOST = "0.0.0.0"
	CONN_PORT = "9000"
	CONN_TYPE = "tcp"
)

func init() {
	OpenDB("data.db")

	schema := `
		id            INTEGER  PRIMARY KEY,
		dev_id        TEXT     NOT NULL,
		timestamp     DATETIME NOT NULL,
		latitude      REAL     NOT NULL,
		longitude     REAL     NOT NULL,
		direction     TEXT     NOT NULL
	`

	CreateTable("data", schema)
}

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
			Log("<- (connection close) EOF")
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
			data := ParseData(string(buf))
			Log(fmt.Sprintln("-- (db)", data.ToQuery("data")))
			_, err := DB.Exec(data.ToQuery("data"))
			CheckErr(err)
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

		Log(fmt.Sprintln("<- (dropped)", string(buf)))
	}
}

func Log(msg string) {
	logger.SetPrefix(time.Now().Format("2006-01-02 15:04:05") + " ")
	logger.Print(msg)
}

type Data struct {
	Header       string
	DevID        string
	Keyword      string
	TimeStamp    time.Time
	SIMCard      string
	GPSState     string
	ZeroTZTime   string
	GPSState2    string
	Latitude     float64
	HemisphereNS string
	Longitude    float64
	HemisphereEW string
	Direction    string
	Altitude     string
	AACState     string
	DoorState    string
	Oil1         string
	Oil2         string
	Temperature  string
}

func (d Data) ToQuery(db string) string {
	query := fmt.Sprintf("INSERT INTO %s (dev_id, timestamp, latitude, longitude, direction) "+
		"VALUES ('%s', '%s', '%f', '%f', '%s')",
		db, d.DevID, d.TimeStamp.Format(time.RFC3339), d.Latitude, d.Longitude, d.Direction)

	return query
}

func ParseData(data string) Data {
	// from: https://sourceforge.net/p/opengts/discussion/579835/thread/c0706b88/6fa0/4068/attachment/coban%20GPRS%20PROTOCOL_COCO%20HUO%20%2020141212.pdf
	//
	// imei:111222333444555,tracker,171016213738,,F,111111.000,A,2222.0000,N,00001.2222,W,0.00,0;
	//
	// protocol:
	//
	// <header>:<ID / IMEI>,<KEYWORD>,<TIMESTAMP>,<SIM CARD NUMBER>,<GPS STATUS>,<H,M,S OF ZERO TZ>,
	//   <GPS STATUS 2>,<LATITUDE>,<N|S>,<LONGITUDE>,<E|W>,<SPEED>,<1=ADDR-REQ, >1.0 = DIRECTION>,<ALTITUDE>,
	//   <ACC STATUS>,<VEHICLE DOOR>,<REMAINING OIL %>,<REMAINING OIL % - 2ND TANK>,<TEMPERATURE SENSOR>;
	//
	// GPS STATUS:
	// F = valid
	// L = NO GPS (LBS MODE) - LAC instead of latitude, Cellid instead of longitude
	//
	// GPS STATUS 2: - NMEA CODE: http://www.gpsinformation.org/dale/nmea.htm#RMC
	// A = ACTIVE
	// V = VOID

	// ensure length is correct

	a := strings.Split(strings.Trim(data, ";"), ",")
	b := strings.Split(a[0], ":")

	d := Data{
		Header:       b[0],
		DevID:        b[1],
		Keyword:      a[1],
		SIMCard:      a[3],
		GPSState:     a[4],
		ZeroTZTime:   a[5],
		GPSState2:    a[6],
		HemisphereNS: a[8],
		HemisphereEW: a[10],
		Direction:    a[11],
		Altitude:     a[12],
		//AACState:     a[13],
		//DoorState:    a[14],
		//Oil1:         a[15],
		//Oil2:         a[16],
		//Temperature:  a[17],
	}

	var ts time.Time
	layout := "060102150405"
	ts, err := time.Parse(layout, a[2])
	CheckErr(err)

	if a[2] != "" {
		d.TimeStamp = ts
	}

	if a[7] != "" {
		d.Latitude = latitude(a[7], d.HemisphereNS)
	}

	if a[9] != "" {
		d.Longitude = longitude(a[9], d.HemisphereEW)
	}

	return d
}

func latitude(lat, ns string) float64 {
	lat_degrees, _ := strconv.ParseInt(lat[0:2], 0, 64)
	lat_minutes, _ := strconv.ParseFloat(lat[2:], 64)
	l := float64(lat_degrees) + (lat_minutes / 60)
	if ns == "S" {
		l = l * -1
	}
	f := fmt.Sprintf("%.4f", l)
	l, _ = strconv.ParseFloat(f, 64)
	return l
}

func longitude(long, ew string) float64 {
	long_degrees, _ := strconv.ParseInt(long[0:3], 0, 64)
	long_minutes, _ := strconv.ParseFloat(long[3:], 64)
	l := float64(long_degrees) + (long_minutes / 60)
	if ew == "W" {
		l = l * -1
	}
	f := fmt.Sprintf("%.4f", l)
	l, _ = strconv.ParseFloat(f, 64)
	return l
}
