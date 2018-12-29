package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"github.com/spf13/cobra"
)

var directory string
var exact bool

// Executes the Cobra function
func main() {
	execute()
}

// find lists files that matches a pattern in a specific directory
func find(filename string) {
	var foundFiles bool

	if !exact {
		filename = ".*" + filename + ".*"
	}

	err := filepath.Walk(directory,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if matched, _ := regexp.MatchString(filename, path); matched {
				if !foundFiles {
					foundFiles = true
				}
				fmt.Println(path)
			}
			return nil
		})

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if !foundFiles {
		os.Exit(1)
	}
}

func execute() {
	cmd.PersistentFlags().StringVarP(&directory, "dir", "d", ".", "Specify directory to search in")
	cmd.PersistentFlags().BoolVarP(&exact, "exact", "e", false, "Only return results if the name matches exactly")
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// cmd is the base command
var cmd = &cobra.Command{
	Use:   "ff",
	Short: "Find a file matching a given name",
	Long:  `Find a file matching a given name`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please supply a file to search for")
			os.Exit(1)
		}

		find(args[0])
	},
}
