package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
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
	ELat   float64
	ELon   float64
}

func ReadStorms() map[string]Storm {
	storms := make(map[string]Storm)

	f, err := os.Open("1950-2022_torn.csv")
	defer f.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	reader := csv.NewReader(f)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println(err.Error())
		}

		if record[0] == "om" {
			continue
		}

		var conversion [4]float64

		for i := 16; i <= 19; i++ {
			conversion[i-16], err = strconv.ParseFloat(record[i], 64)
			if err != nil {
				fmt.Println(err.Error())
			}
		}

		storms[record[0]] = Storm{
			Year:   record[1],
			Month:  record[2],
			Day:    record[3],
			Time:   record[4],
			Date:   record[5],
			FScale: record[10],
			SLat:   conversion[0],
			SLon:   conversion[1],
			ELat:   conversion[2],
			ELon:   conversion[3],
		}
	}

	return storms
}
