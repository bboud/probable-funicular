package main

import (
	"fmt"
	"math"
	"net/http"
)

// https://en.wikipedia.org/wiki/Haversine_formula
func HaversinDistance(rlat, rlon, slat, slon float64) float64 {
	const earthRadius = 6378 //km

	slat = slat * math.Pi / 180
	slon = slon * math.Pi / 180
	rlat = rlat * math.Pi / 180
	rlon = rlon * math.Pi / 180

	deltaPsi := slat - rlat
	deltaLamda := slon - rlon

	h := math.Pow(math.Sin(deltaPsi/2), 2) + math.Cos(slat)*math.Cos(rlat)*math.Pow(math.Sin(deltaLamda/2), 2)

	c := math.Asin(math.Sqrt(h))

	return 2 * earthRadius * c
}

func main() {
	// f, err := os.Open("1950-2022_torn.csv")
	// defer f.Close()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// storms, err := ReadStorms(f)
	// radars := ReadRadars()

	// // Compute the radars that apply to each storm
	// for i := 0; i < len(storms); i++ {
	// 	slon, slat := storms[i].SLon, storms[i].SLat
	// 	for _, r := range radars.Features {
	// 		rlon, rlat := r.Geometry.Coordinates[0], r.Geometry.Coordinates[1]
	// 		distance := HaversinDistance(rlat, rlon, slat, slon)

	// 		if distance < 24 {
	// 			storms[i].Radars = append(storms[i].Radars, r.Properties.SiteID)
	// 			fmt.Println(distance)
	// 		}
	// 	}
	// }

	// for _, s := range storms {
	// 	for site := range s.Radars {
	// 		//emailadd=ceeandstuff%40gmail.com&startHour=00&endHour=05&id=KPAH&yyyy=2016&mm=04&dd=14&dsi=7000&product=ABL3ALL
	// 		resp, err := http.Post("https://www.ncdc.noaa.gov/nexradinv/ordercomplete.jsp", "application/x-www-form-urlencoded", "emailadd=ceeandstuff%40gmail.com&startHour=&endHour=05&id=KPAH&yyyy=2016&mm=04&dd=14&dsi=7000&product=ABL3ALL")
	// 		if err != nil {
	// 			fmt.Println(err.Error())
	// 		}
	// 		defer resp.Body.Close()
	// 	}
	// }

	resp, err := http.Post("https://www.ncdc.noaa.gov/nexradinv/ordercomplete.jsp", "application/x-www-form-urlencoded", "emailadd=ceeandstuff%40gmail.com&startHour=&endHour=05&id=KPAH&yyyy=2016&mm=04&dd=14&dsi=7000&product=ABL3ALL")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer resp.Body.Close()
}
