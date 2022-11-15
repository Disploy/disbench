package cmd

import (
	"fmt"
	"os"

	"github.com/Disploy/disbench/github"
	"github.com/Disploy/disbench/internal_benchmark"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var internalBenchmarkCmd = &cobra.Command{
	Use:   "benchmark",
	Short: "Run a benchmark test on Disploy.",
	Run: func(cmd *cobra.Command, args []string) {
		iterations, _ := cmd.Flags().GetInt("iterations")
		url, _ := cmd.Flags().GetString("url")
		githubTarget, _ := cmd.Flags().GetString("github")
		title, _ := cmd.Flags().GetString("title")
		parsedGithub, err := github.ParseGitHubTarget(githubTarget)
		githubToken := os.Getenv("GITHUB_TOKEN")

		if err != nil {
			fmt.Printf("client: could not parse GitHub target: %s\n", err)
			os.Exit(1)
		}

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

			if githubToken != "" {
				github.InitClient(githubToken)
				github.PostComment(parsedGithub.Owner, parsedGithub.Repo, parsedGithub.IssueNumber, fmt.Sprintf("Benchmark: %s\n\n%s", title, benchmark.Markdown()))
				color.Magenta("[github] Posted benchmark results to %s/%s#%d", parsedGithub.Owner, parsedGithub.Repo, parsedGithub.IssueNumber)
			}
		}
	},
}

func init() {
	internalCmd.AddCommand(internalBenchmarkCmd)

	internalBenchmarkCmd.Flags().IntP("iterations", "i", 100, "The number of times to run the benchmark.")
	internalBenchmarkCmd.Flags().StringP("url", "u", "", "The URL to send the benchmark request to.")
	internalBenchmarkCmd.Flags().StringP("github", "g", "", "The GitHub target to post the benchmark results to formatted as <owner>/<repo>#<issue number>.")
	internalBenchmarkCmd.Flags().StringP("title", "t", "Untitled benchmark", "The title of the benchmark.")
}
