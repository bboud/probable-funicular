package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Feature struct {
	Properties struct {
		SiteID string
	}
	Geometry struct {
		Coordinates []float64
	}
}

type FeatureCollection struct {
	Features []Feature
}

func ReadRadars() FeatureCollection {
	f, err := os.Open("Weather_Radar_Stations.geojson")
	defer f.Close()
	if err != nil {
		fmt.Println(err.Error())
	}

	var features FeatureCollection
	content, err := io.ReadAll(f)
	if err != nil {
		fmt.Println(err.Error())
	}

	json.Unmarshal(content, &features)

	return features
}
