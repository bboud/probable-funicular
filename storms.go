package main

import (
	"encoding/csv"
	"io"
	"strconv"
)

type Storm struct {
	Year   string
	Month  string
	Day    string
	Date   string
	Time   string
	FScale string
	SLat   float64
	SLon   float64
	Radars []string
}

func ReadStorms(r io.Reader) ([]Storm, error) {
	var storms []Storm
	var err error
	reader := csv.NewReader(r)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		if record[0] == "om" {
			continue
		}

		slon, err := strconv.ParseFloat(record[16], 64)
		slat, err := strconv.ParseFloat(record[15], 64)

		storms = append(storms, Storm{
			Year:   record[1],
			Month:  record[2],
			Day:    record[3],
			Time:   record[4],
			Date:   record[5],
			FScale: record[10],
			SLon:   slon,
			SLat:   slat,
		})
	}

	return storms, err
}
