package main

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

// copy files and directories
func copyR(source, destination string, overwrite, ignoreExisting bool) error {
	stat, err := os.Stat(source)
	if nil != err {
		return err
	}

	if stat.IsDir() {
		return copyDirectory(source, destination, overwrite, ignoreExisting)
	}

	if err = copyFile(source, destination, overwrite); nil != err && (!ignoreExisting && os.IsExist(err)) {
		return err
	} else {
		logIgnoredError(err)
	}

	return nil
}

func copyFile(source, destination string, overwrite bool) error {
	_, err := os.Stat(destination)
	if nil == err && !overwrite {
		return os.ErrExist
	} else {
		logIgnoredError(err)
	}

	logVerbose("making dir %s", filepath.Dir(destination))
	_ = os.MkdirAll(filepath.Dir(destination), os.ModePerm)

	src, err := os.Open(source)
	if nil != err {
		return err
	}
	defer src.Close()

	dst, err := os.OpenFile(destination, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if nil != err {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	return err
}

func copyDirectory(source, destination string, overwrite, ignoreExisting bool) error {
	_, err := os.Stat(destination)
	if nil == err && overwrite {
		os.RemoveAll(destination)
	} else if nil == err && !ignoreExisting {
		return os.ErrExist
	}

	if nil == err || os.IsNotExist(err) {
		os.MkdirAll(destination, 0744)
	}

	files, err := ioutil.ReadDir(source)
	if nil != err {
		return err
	}

	for _, file := range files {
		src := filepath.Join(source, file.Name())
		dst := filepath.Join(destination, file.Name())
		if err := copyR(src, dst, overwrite, ignoreExisting); nil != err && (!ignoreExisting && os.IsExist(err)) {
			return err
		} else {
			logIgnoredError(err)
		}
	}

	return nil
}
