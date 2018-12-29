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
var exact, nocolour bool
var ignoredir []string

// Executes the Cobra function
func main() {
	execute()
}

// find lists files that matches a pattern in a specific directory
func find(filename string) {
	var foundFiles bool

	red := color.New(color.FgRed).SprintFunc()

	if !exact {
		filename = ".*" + filename + ".*"
	}

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

			matched, err := regexp.MatchString(filename, path)
			if matched && err == nil {
				if !foundFiles {
					foundFiles = true
				}
				if !nocolour {
					filename = strings.Replace(filename, ".*", "", -1)
					path = strings.Replace(path, filename, red(filename), -1)
				}
				fmt.Println(path)
			} else if err != nil {
				return err
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
	cmd.PersistentFlags().BoolVarP(&exact, "exact", "e", false, "Only return results if the name matches exactly")
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

ff [filename]

By default it wraps the filename in a wildcard expression, so "foo" would
be defined by the regex:

/.*foo.*/

The "-e" flag removes this wildcard expression, so further exact patterns may
be passed in.`,
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
