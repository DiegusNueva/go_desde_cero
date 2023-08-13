package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path"
	"regexp"
	"runtime"
	"sort"
	"strings"
	"time"

	"github.com/AJRDRGZ/fileinfo"
	"github.com/fatih/color"
	"golang.org/x/exp/constraints"
)

// Windows os  system
const Windows = "windows"

func main() {
	// filter patter
	flagPattern := flag.String("p", "", "filter by pattern")
	flagAll := flag.Bool("a", false, "all files including hide files")
	flagNumberRecords := flag.Int("n", 0, "number of records")

	// order flags
	hasOrderByTime := flag.Bool("t", false, "sort by time, oldest first")
	hasOrderBySize := flag.Bool("s", false, "sort by file size, smallest first")
	hasOrderByReverse := flag.Bool("r", false, "reverse order while sorting")

	flag.Parse()
	path := flag.Arg(0)

	if path == "" {
		path = "."
	}

	dirs, err := os.ReadDir(path)

	if err != nil {
		panic(err)
	}

	fs := []file{}

	for _, dir := range dirs {
		isHidden := isHidden(dir.Name(), path)

		if isHidden && !*flagAll {
			continue
		}

		if *flagPattern != "" {

			isMatched, err := regexp.MatchString("(?i)"+*flagPattern, dir.Name())

			if err != nil {
				panic(err)
			}

			if !isMatched {
				continue
			}
		}

		f, err := getFile(dir, isHidden)
		if err != nil {
			panic(err)
		}

		fs = append(fs, f)
	}

	if !*hasOrderBySize || !*hasOrderByTime {
		orderByName(fs, *hasOrderByReverse)
	}

	if *hasOrderBySize && !*hasOrderByTime {
		orderBySize(fs, *hasOrderByReverse)
	}

	if *hasOrderByTime {
		orderByTime(fs, *hasOrderByReverse)
	}

	if *flagNumberRecords == 0 || *flagNumberRecords > len(fs) {
		*flagNumberRecords = len(fs)
	}
	printList(fs, *flagNumberRecords)

}

func mySort[T constraints.Ordered](i, j T, isReverse bool) bool {
	if isReverse {
		return i > j
	}

	return i < j
}

func orderByTime(files []file, isReverse bool) {
	sort.SliceStable(files, func(i, j int) bool {

		return mySort(
			files[i].modificationTime.Unix(),
			files[j].modificationTime.Unix(),
			isReverse,
		)
	})
}

func orderByName(files []file, isReverse bool) {
	sort.SliceStable(files, func(i, j int) bool {

		return mySort(strings.ToLower(files[i].name), strings.ToLower(files[j].name), isReverse)

	})
}

func orderBySize(files []file, isReverse bool) {

	sort.SliceStable(files, func(i, j int) bool {
		return mySort(
			files[i].size,
			files[j].size,
			isReverse,
		)
	})

}

func printList(fs []file, nRecords int) {
	for _, file := range fs[:nRecords] {
		style := mapStyleByFileType[file.fileType]
		fmt.Printf("%s %s %s %10d %s %s %s%s %s\n",
			file.mode, file.userName, file.groupName, file.size, file.modificationTime.Format(time.DateTime),
			style.icon, setColor(file.name, style.color), style.symbol, markHidden(file.isHidden))
	}
}

func getFile(dir fs.DirEntry, isHidden bool) (file, error) {

	info, err := dir.Info()

	if err != nil {
		return file{}, fmt.Errorf("dir.Info(): %v", err)
	}

	userName, groupName := fileinfo.GetUserAndGroup(info.Sys())

	f := file{
		name:             dir.Name(),
		isDir:            dir.IsDir(),
		isHidden:         isHidden,
		userName:         userName,
		groupName:        groupName,
		size:             info.Size(),
		modificationTime: info.ModTime(),
		mode:             info.Mode().String(),
	}

	setFile(&f)

	return f, nil
}

func setFile(f *file) {
	switch {
	case isLink(*f):
		f.fileType = fileLink
	case f.isDir:
		f.fileType = fileDirectory
	case isExec(*f):
		f.fileType = fileExecutable
	case isCompress(*f):
		f.fileType = fileCompress
	case isImage(*f):
		f.fileType = fileImage
	default:
		f.fileType = fileRegular
	}
}

func isLink(f file) bool {
	return strings.HasPrefix(strings.ToUpper(f.mode), "L")
}

func isExec(f file) bool {
	if runtime.GOOS == Windows {
		return strings.HasSuffix(f.name, exe)
	}

	return strings.Contains(f.mode, "x")
}

func isCompress(f file) bool {
	return strings.HasSuffix(f.name, zip) ||
		strings.HasSuffix(f.name, gz) ||
		strings.HasSuffix(f.name, tar) ||
		strings.HasSuffix(f.name, rar) ||
		strings.HasSuffix(f.name, deb)
}

func isImage(f file) bool {
	return strings.HasSuffix(f.name, png) ||
		strings.HasSuffix(f.name, jpg) ||
		strings.HasSuffix(f.name, gif)
}

func isHidden(fileName, basePath string) bool {

	filePath := fileName

	if runtime.GOOS == Windows {
		filePath = path.Join(basePath, fileName)
	}

	return fileinfo.IsHidden(filePath)

}

func setColor(nameFile string, styleColor color.Attribute) string {
	switch styleColor {
	case color.FgBlue:
		return blue(nameFile)
	case color.FgGreen:
		return green(nameFile)
	case color.FgRed:
		return red(nameFile)
	case color.FgMagenta:
		return magenta(nameFile)
	case color.FgCyan:
		return cyan(nameFile)
	}

	return nameFile
}

func markHidden(isHidden bool) string {
	if !isHidden {
		return ""
	}

	return yellow("ø")
}
