package internal_benchmark

func Start(
	iterations int,
	url string,
) []Benchmark {
	var benchmarks []Benchmark

	benchmarks = append(benchmarks, ExecutePingBenchmark(iterations, url, "hello world!!!!!!!!"))

	return benchmarks
}
