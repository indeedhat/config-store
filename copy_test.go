package configstore

import (
	"io/ioutil"
	"os"
	"path"
	"strings"
	"testing"
)

const (
	t_SRC = "testdata/src"
	t_DST = "testdata/dst"
)

func setup() {
	// tearing down before test rather than after allows for visual inspection if things go wrong
	teardown()

	if err := os.MkdirAll(t_DST, os.ModePerm); nil != err {
		panic(err)
	}
}

func teardown() {
	if _, err := os.Stat(t_DST); nil != err || !os.IsNotExist(err) {
		_ = os.RemoveAll(t_DST)
	}
}

func TestMain(m *testing.M) {
	setup()

	code := m.Run()

	os.Exit(code)
}

func TestCopyFile(t *testing.T) {
	src := path.Join(t_SRC, "file_1")
	dst := path.Join(t_DST, "file_1")

	if err := copyR(src, dst, false, false); nil != err {
		t.Errorf("copy returned error %s", err)
	}

	if _, err := os.Stat(dst); nil != err {
		t.Error("Failed to copy file")
	}
}

func TestCopyFileNotExists(t *testing.T) {
	src := path.Join(t_SRC, "file_2")
	dst := path.Join(t_DST, "file_2")

	if err := copyR(src, dst, false, false); !os.IsNotExist(err) {
		t.Errorf("did not produce error on none existant file")
	}
}

func TestCopyFileToExists(t *testing.T) {
	src := path.Join(t_SRC, "file_1")
	dst := path.Join(t_DST, "file_1")

	if err := copyR(src, dst, false, false); !os.IsExist(err) {
		t.Errorf("did not produce error on existing destination: %s", err)
	}
}

func TestCopyFileOverwrite(t *testing.T) {
	src := path.Join(t_SRC, "file_1")
	dst := path.Join(t_DST, "file_1")
	contents := "data"

	if err := ioutil.WriteFile(dst, []byte(contents), os.ModePerm); nil != err {
		t.Error("Failed to setup test")
	}

	if err := copyR(src, dst, true, false); nil != err {
		t.Errorf("failed to overwrite file: %s", err)
	}

	if data, err := ioutil.ReadFile(dst); nil != err || 0 != strings.Compare("d\n", string(data)) {
		t.Errorf("file was not overwritten: ;%s;", string(data))
	}
}

func TestCopyDirectory(t *testing.T) {
	src := path.Join(t_SRC, "dir_1")
	dst := path.Join(t_DST, "dir_1")

	if err := copyR(src, dst, false, false); nil != err {
		t.Error("failed to copy dir")
	}

	dstFiles := []string{
		path.Join(t_DST, "dir_1/file_2"),
		path.Join(t_DST, "dir_1/file_3"),
		path.Join(t_DST, "dir_1/dir_2/file_4"),
	}

	for _, dstFile := range dstFiles {
		if _, err := os.Stat(dstFile); nil != err {
			t.Error("not all files were found in the destination directory")
		}
	}
}

func TestCopyDirectoryNotExists(t *testing.T) {
	src := path.Join(t_SRC, "not_exist")
	dst := path.Join(t_DST, "not_exist")

	if err := copyR(src, dst, false, false); nil == err {
		t.Error("did not produce an error on none existing source dir")
	}
}

func TestCopyDirectoryToExists(t *testing.T) {
	src := path.Join(t_SRC, "dir_1")
	dst := path.Join(t_DST, "dir_1")

	if err := copyR(src, dst, false, false); nil == err {
		t.Error("did not produce error on existing destination directory")
	}
}

func TestCopyDirectoryOverwrite(t *testing.T) {
	src := path.Join(t_SRC, "dir_3")
	dst := path.Join(t_DST, "dir_1")

	if err := copyR(src, dst, true, false); nil != err {
		t.Error("cound not overwrite existing dir")
	}

	if _, err := os.Stat(path.Join(dst, "file_5")); nil != err {
		t.Error("did not actually overwrite the directory")
	}

}

func TestCopyDirectoryIgnoreExisting(t *testing.T) {
	src := path.Join(t_SRC, "dir_1")
	dst := path.Join(t_DST, "dir_1")

	if err := copyR(src, dst, false, true); nil != err {
		t.Error("failed to copy dir")
	}

	dstFiles := []string{
		path.Join(t_DST, "dir_1/file_2"),
		path.Join(t_DST, "dir_1/file_3"),
		path.Join(t_DST, "dir_1/dir_2/file_4"),
		path.Join(t_DST, "dir_1/file_5"),
	}

	for _, dstFile := range dstFiles {
		if _, err := os.Stat(dstFile); nil != err {
			t.Error("not all files were found in the destination directory")
		}
	}
}
