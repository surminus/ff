// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// Executes the Cobra function
func main() {
	execute()
}

func directories(d string) []os.FileInfo {
	var dirs []os.FileInfo
	files, err := ioutil.ReadDir(d)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, file := range files {
		if file.IsDir() {
			dirs = append(dirs, file)
		}
	}
	return dirs
}

// find lists files that matches a pattern
func find(dir, filename string) {
	for _, d := range directories(dir) {
		files, err := filepath.Glob(filepath.Join(d.Name(), "/*"+filename+"*"))

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		output(files)
	}
}

// output simply prints the slice of files and exits with an error if none
// are found
func output(files []string) {
	for _, f := range files {
		fmt.Println(f)
	}
}

func execute() {
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

		find(".", args[0])
	},
}
