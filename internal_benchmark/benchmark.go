package internal_benchmark

func Start(
	iterations int,
	timeBetween float32,
	url string,
) []Benchmark {
	var benchmarks []Benchmark

	benchmarks = append(benchmarks, ExecutePingBenchmark(iterations, timeBetween, url, "hello world!!!!!!!!"))

	return benchmarks
}
