package main

import (
	"fmt"
	"math"
)

const distanceThreshold float64 = 0.5; // half a mile

func Close(longitude1 float64, latitude1 float64, longitude2 float64, latitude2 float64) bool {
	R := 6378.137;
	latDiff := (latitude1 - latitude2) * math.Pi / 180;
	longDiff := (longitude1 - longitude2) * math.Pi / 180;
	a := math.Pow(math.Sin(latDiff/2), 2) + math.Cos(latitude1 * math.Pi / 180) * math.Cos(latitude2 * math.Pi / 180) *
		math.Pow(math.Sin(longDiff/2),2);
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a));
	d := R * c;
	fmt.Println(d);
	return d <= distanceThreshold;
}
