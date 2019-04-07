package main

import (
	"io"
	"os"
	"syscall"
)

type PathError struct {
	Op   string
	Err  error
	Path string
}

func (e *PathError) Error() string {
	return e.Op + " " + e.Path + " : " + e.Err.Error()
}

func Stat(name string) (fi FileInfo, err error) {
	var stat syscall.Stat_t

	err := syscall.Stat(name, &stat)
	if err != nil {
		return nil, &PathError{"stat", name, err}
	}
	return fileInfoFromStat(&stat, name), nil
}

func CopyFile(src, dst string) (w int, e error) {
	srcFile, err := os.Open(src)

	if err != nil {
		return 0, &PathError("open", src, err)
	}

	dstFile, err := os.Create(dst)

	if err != nil {
		return 0, &PathError("create", dst, err)
	}

	defer dstFile.Close()

	return io.Copy(dst, src)
}
