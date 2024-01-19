package gopixanalyzer

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gopixanalyzer",
	Short: "gopixanalyzer is a analyzer for pictures.",
	Long:  `gopixanalyzer is a fast CLI tool for analyzing pictures. It can be used to extracting metadata from pictures.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
