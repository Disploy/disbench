package cmd

import (
	"github.com/Disploy/disbench/internal_benchmark"
	"github.com/spf13/cobra"
)

var internalBenchmarkCmd = &cobra.Command{
	Use:   "benchmark",
	Short: "Run a benchmark test on Disploy.",
	Run: func(cmd *cobra.Command, args []string) {
		iterations, _ := cmd.Flags().GetInt("iterations")
		url, _ := cmd.Flags().GetString("url")

		if url == "" {
			cmd.Help()
			return
		}

		if iterations < 1 {
			iterations = 1
		}

		var benchmarks = internal_benchmark.Start(iterations, url)

		for _, benchmark := range benchmarks {
			benchmark.Print()
		}
	},
}

func init() {
	internalCmd.AddCommand(internalBenchmarkCmd)

	internalBenchmarkCmd.Flags().IntP("iterations", "i", 100, "The number of times to run the benchmark.")
	internalBenchmarkCmd.Flags().StringP("url", "u", "", "The URL to send the benchmark request to.")

}
