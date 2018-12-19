package main

import (
	"os"
)

type File struct {
	Path string
	Data []byte

	// file permision
	Perm os.FileMode
}

type Directory struct {
	Path string

	// directory permision
	Perm os.FileMode
}

type Archive struct {
	Dirs  []Directory
	Files []File
}

/*
type Reader struct{}

func (*Reader) Read(p []byte) (int, error) {
}

func (*Reader) Close() error {
}

// Reset discards the Reader z's state and makes it equivalent to the result of
// its original state from NewReader, but reading from r instead. This permits
// reusing a Reader rather than allocating a new one.
func (z *Reader) Reset(r io.Reader) error {
}
*/
