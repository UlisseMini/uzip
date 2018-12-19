package main

import (
	"fmt"
	"github.com/UlisseMini/utils/gobber"
	"io/ioutil"
	"os"
	"path/filepath"
)

// archive will be set from crawler
var archive Archive

// zip is a simple wrapper for crawler
// TODO make this work with a bunch of files and not require a directory
func zip(dir string, archiveName string) error {
	err := filepath.Walk(dir, crawler)
	if err != nil {
		return err
	}

	// write it to a file
	return gobber.Write(archiveName, archive)
}

// unzip unzips an archive
func unzip(zipfile string) error {
	err := gobber.Read(zipfile, &archive)
	if err != nil {
		return err
	}

	// create all the directories
	for _, dir := range archive.Dirs {
		fmt.Println(dir.Path)
		err := os.Mkdir(dir.Path, dir.Perm)
		if err != nil {
			return err
		}
	}

	// create all the files
	for _, file := range archive.Files {
		fmt.Println(file.Path)
		err := ioutil.WriteFile(file.Path, file.Data, file.Perm)
		if err != nil {
			return err
		}
	}

	return nil
}

// WARNING can't be called concurrently (i should add mutexes later)
func crawler(path string, f os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	fmt.Println(path)

	if f.IsDir() == true {
		dir := Directory{
			Path: path,
			Perm: f.Mode(),
		}

		archive.Dirs = append(archive.Dirs, dir)
	} else {
		data, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		file := File{
			Path: path,
			Data: data,
			Perm: f.Mode(),
		}

		archive.Files = append(archive.Files, file)
	}

	return nil
}
