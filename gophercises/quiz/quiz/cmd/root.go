/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var csvFile string
var answerCount int
var runtimeSeconds int

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "quiz",
	Short: "Quiz to solve csv file problems within time",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		if !strings.HasSuffix(csvFile, ".csv") {
			fmt.Println("Use only .csv format")
			return
		}

		timer := time.NewTimer(time.Duration(runtimeSeconds) * time.Second)
		answersCh := make(chan struct{})

		go func() {
			readFromFile()
			answersCh <- struct{}{}
		}()

		select {
		case <-timer.C:
		case <-answersCh:
		}

		fmt.Printf("\nYour result is: %d\n", answerCount)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&csvFile, "csv", "problemList/problems.csv", "csv file with problems (in format - problem,answer)")
	rootCmd.PersistentFlags().IntVar(&runtimeSeconds, "sec", 30, "seconds to run")
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
