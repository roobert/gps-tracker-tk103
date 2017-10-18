package handler

import (
	"bytes"
	"io"
	"net/http"
	"time"

	. "github.com/roobert/golang-db"
	. "github.com/roobert/golang-error"

	"github.com/twpayne/go-gpx"
)

func API(w http.ResponseWriter, r *http.Request) {
	query := "select timestamp, latitude, longitude from data"
	rows, err := DB.Query(query)
	CheckErr(err)

	var wpts []*gpx.WptType
	var timestamp time.Time
	var latitude float64
	var longitude float64

	for rows.Next() {
		err = rows.Scan(&timestamp, &latitude, &longitude)
		CheckErr(err)

		wpt := &gpx.WptType{
			Lat:  latitude,
			Lon:  longitude,
			Time: timestamp,
		}

		wpts = append(wpts, wpt)
	}

	rows.Close()

	g := &gpx.GPX{
		Version: "1.0",
		Creator: "Whatever",
		Wpt:     wpts,
	}

	buf := new(bytes.Buffer)

	err = g.WriteIndent(buf, "", "  ")
	CheckErr(err)

	io.WriteString(w, buf.String())
}
