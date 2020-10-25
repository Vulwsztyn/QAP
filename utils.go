package main

import "time"

func measureTime(fn func()) int64 {
	start := time.Now()
	fn()
	stop := time.Since(start)
	return stop.Microseconds()
}
