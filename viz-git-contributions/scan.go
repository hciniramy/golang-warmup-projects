package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/user"
	"strings"
)

func scanGitFolders(folders []string, folder string) []string {
	folder = strings.TrimSuffix(folder, "/")

	f, err := os.Open(folder)
	defer f.Close()

	if err != nil {
		log.Fatal(err)
	}

	files, err := f.Readdir(-1)

	if err != nil {
		log.Fatal(err)
	}

	var path string
	for _, file := range files {
		if file.IsDir() {
			path = folder + "/" + file.Name()
			if file.Name() == ".git" {
				path = strings.TrimSuffix(path, "/.git")
				fmt.Printf(path)
				folders = append(folders, path)
				continue
			}
			if file.Name() == "vendor" || file.Name() == "nodes_modules" {
				continue
			}

			folders = scanGitFolders(folders, path)
		}
	}

	return folders
}
func addNewSliceElementsToFile(path string, newRepos []string) {
	existingRepos := parseFileLinesToSlice(path)
	repos := joinSlices(newRepos, existingRepos)
	dumpStringSliceToFile(repos, path)
}

func dumpStringSliceToFile(repos []string, path string) {
	content := strings.Join(repos, "\n")
	os.WriteFile(path, []byte(content), 0755)
}

func joinSlices(new []string, existing []string) []string {
	for _, i := range new {
		if !sliceContains(existing, i) {
			existing = append(existing, i)
		}
	}
	return existing
}

func sliceContains(existing []string, i string) bool {
	for _, v := range existing {
		if v == i {
			return true
		}
	}
	return false
}

func parseFileLinesToSlice(path string) []string {
	f := openFile(path)
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		if err != io.EOF {
			panic(err)
		}
	}

	return lines
}

func openFile(path string) *os.File {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_RDWR, 0755)
	if err != nil {
		if os.IsNotExist(err) {
			_, err = os.Create(path)
			if err != nil {
				panic(err)
			}
		} else {
			panic(err)
		}
	}

	return f
}

func getDotFilePath() string {
	usr, err := user.Current()

	if err != nil {
		log.Fatal(err)
	}

	dotfile := usr.HomeDir + "/.gogitlocalstats"

	return dotfile
}

func recursiveScanFolder(folder string) []string {
	return scanGitFolders(make([]string, 0), folder)
}
