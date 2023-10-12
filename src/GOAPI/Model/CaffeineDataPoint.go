package main

import "time"

type CaffeineDataPoint struct {
	day       time.Time
	countInMG int
	notes     string
}
