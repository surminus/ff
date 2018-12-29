package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var directory string
var nocolour bool
var ignoredir []string

// Executes the Cobra function
func main() {
	execute()
}

// find lists files that matches a pattern in a specific directory
func find(filename string) {
	var foundFiles bool

	red := color.New(color.FgRed).SprintFunc()

	err := filepath.Walk(directory,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				fmt.Println(err)
				return err
			}

			for _, ignored := range ignoredir {
				if info.IsDir() && info.Name() == ignored {
					return filepath.SkipDir
				}
			}

			// Only work on files, not directories
			if !info.IsDir() {
				matched, err := regexp.MatchString(filename, filepath.Base(path))
				if matched && err == nil {
					if !foundFiles {
						foundFiles = true
					}

					// We want to highlight only the name of the file in path,
					// not any directory names since we are only matching files
					if !nocolour {
						// Split the path into a slice
						pathSlice := strings.Split(path, "/")
						// Fetch the name of the file
						match := pathSlice[len(pathSlice)-1]
						// Remove the original file name
						pathSlice = pathSlice[:len(pathSlice)-1]
						// Re-colour the match in the filename only and add
						// back to slice
						pathSlice = append(pathSlice, strings.Replace(match, filename, red(filename), -1))
						// Rejoin the slice
						path = strings.Join(pathSlice, "/")
					}

					fmt.Println(path)
				} else if err != nil {
					return err
				}
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
	cmd.PersistentFlags().StringSliceVar(&ignoredir, "ignore-dir", []string{}, "Do not search in the specified directory. Can be specified multiple times")
	cmd.PersistentFlags().BoolVar(&nocolour, "no-colour", false, "Disable outputting in colour")
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// cmd is the base command
var cmd = &cobra.Command{
	Use:   "ff",
	Short: "Find a file matching a given name",
	Long: `Find a file matching a given name.

Specify the pattern as the first argument:

ff [filename]`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please supply a file to search for")
			fmt.Println()
			cmd.Usage()
			os.Exit(1)
		}

		find(args[0])
	},
}
