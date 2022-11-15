package internal_benchmark

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
)

func ExecutePingBenchmark(
	iterations int,
	url string,
	expectedContent string,
) Benchmark {
	var benchmark Benchmark
	benchmark.Iterations = iterations

	for i := 0; i < iterations; i++ {
		var startTime = time.Now()
		var response, err = http.Post(url, "application/json", strings.NewReader(PingPayload))

		if err != nil {
			fmt.Printf("client: could not create request: %s\n", err)
			os.Exit(1)
		}

		var endTime = time.Now()
		var duration = endTime.Sub(startTime)

		if err != nil {
			benchmark.Errors++
			continue
		}

		if response.StatusCode != 200 {
			benchmark.Errors++
			continue
		}

		var payload PingResponse
		err = json.NewDecoder(response.Body).Decode(&payload)

		if err != nil {
			benchmark.Errors++
			continue
		}

		if payload.Data.Content != expectedContent {
			benchmark.Errors++
			continue
		}

		benchmark.Successes++
		benchmark.Durations = append(benchmark.Durations, duration)
	}

	var totalDuration time.Duration

	for _, duration := range benchmark.Durations {
		totalDuration += duration
	}

	benchmark.Average = float64(totalDuration.Milliseconds()) / float64(len(benchmark.Durations))

	benchmark.Print = func() {
		fmt.Printf("%s: %s | %s\n", color.RedString("Benchmark"), color.MagentaString("execute `/ping` command"), color.HiGreenString("%d iterations", benchmark.Iterations))
		fmt.Printf("  %s: %s\n", color.HiMagentaString("Successes"), color.HiGreenString("%d", benchmark.Successes))
		fmt.Printf("  %s: %s\n", color.HiMagentaString("Errors"), color.HiRedString("%d", benchmark.Errors))
		fmt.Printf("  %s: %s\n", color.HiMagentaString("Average"), color.HiGreenString("%fms", benchmark.Average))
		fmt.Printf("  %s: %s\n", color.HiMagentaString("Total"), color.HiGreenString("%fms", float64(totalDuration.Milliseconds())))
	}

	benchmark.Markdown = func() string {
		return fmt.Sprintf(
			"## `/ping` Command Benchmark\n\n"+
				"| Iterations | Successes | Errors | Average | Total |\n"+
				"| --- | --- | --- | --- | --- |\n"+
				"| %d | %d | %d | %fms | %fms |\n",
			benchmark.Iterations,
			benchmark.Successes,
			benchmark.Errors,
			benchmark.Average,
			float64(totalDuration.Milliseconds()),
		)
	}

	return benchmark
}
