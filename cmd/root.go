package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

type TestResult struct {
	Time    string
	Action  string
	Package string
	Test    string
	Elapsed string
}

var rootCmd = &cobra.Command{
	Use:   "showit",
	Short: "showit is a simple example of a go test parser",
	Run: func(cmd *cobra.Command, args []string) {
		s := bufio.NewScanner(os.Stdin)
		results := []TestResult{}
		for s.Scan() {
			result := TestResult{}
			if err := json.Unmarshal([]byte(s.Text()), &result); err != nil {
				fmt.Println(err)
				continue
			}
			results = append(results, result)
		}
		for _, r := range results {
			if r.Action == "pass" {
				fmt.Printf("%s %s\n", color.GreenString("✅ PASS"), r.Test)
			}
			if r.Action == "fail" {
				fmt.Printf("%s %s\n", color.RedString("❌ FAIL"), r.Test)
			}
		}
	},
}

func Execute(version string) {
	rootCmd.Version = version
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
