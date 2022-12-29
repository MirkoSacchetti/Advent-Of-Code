// --- Day 7: No Space Left On Device ---

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type File struct {
	IsDir  bool
	Size   int
	Name   string
	Files  []*File
	Parent *File
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func findFolderOnElfDrive(parent File, size int) (int, int) {
	folderSize := 0
	targetFolder := 999999999999999
	for _, file := range parent.Files {
		if file.IsDir {
			f, t := findFolderOnElfDrive(*file, size)
			folderSize += f
			if t >= size && t < targetFolder {
				targetFolder = t
			}
		} else {
			folderSize += file.Size
		}
	}
	if folderSize >= size && folderSize < targetFolder {
		targetFolder = folderSize
	}
	return folderSize, targetFolder
}

func scanFileOnElfDrive(parent File) (int, int) {
	folderSize := 0
	bigfoldersAcc := 0
	for _, file := range parent.Files {
		if file.IsDir {
			f, b := scanFileOnElfDrive(*file)
			folderSize += f
			bigfoldersAcc += b
		} else {
			folderSize += file.Size
		}
	}
	if folderSize <= 100000 {
		bigfoldersAcc += folderSize
	}
	return folderSize, bigfoldersAcc
}

func main() {
	file, err := os.ReadFile("input.txt")
	check(err)
	lines := strings.Split(string(file), "\n")
	root := File{Name: "/", IsDir: true}
	cwdPointer := &root
	currentCmd := []string{}
	for _, line := range lines {
		if len(line) < 1 {
			continue
		}
		stdout := strings.Split(line, " ")
		if stdout[0] == "$" {
			currentCmd = stdout[1:]
			if currentCmd[0] == "cd" {
				if currentCmd[1] == "/" {
					cwdPointer = &root
					continue
				}
				if currentCmd[1] == ".." {
					cwdPointer = cwdPointer.Parent
					continue
				}
				// find the folder and set to current working directory
				for _, file := range cwdPointer.Files {
					if file.IsDir && file.Name == currentCmd[1] {
						cwdPointer = file
						break
					}
				}
			}
			continue
		}
		if currentCmd[0] == "ls" {
			size, _ := strconv.Atoi(stdout[0])
			cwdPointer.Files = append(cwdPointer.Files, &File{
				IsDir:  stdout[0] == "dir",
				Name:   stdout[1],
				Size:   size,
				Parent: cwdPointer,
			})
		}
	}

	diskCapacity := 70000000
	spaceRequired := 30000000

	total, acc := scanFileOnElfDrive(root)

	spaceUnused := diskCapacity - total
	spaceToFree := spaceRequired - spaceUnused
	total, folderToDelete := findFolderOnElfDrive(root, spaceToFree)

	fmt.Println("part one:", acc)
	fmt.Println("part two:", folderToDelete)
}
