package internal_benchmark

import "time"

type Benchmark struct {
	Iterations int
	Errors     int
	Successes  int
	Durations  []time.Duration
	Average    float64
	Print      func()
}
