package util

import (
	"testing"
)

func TestDownload(t *testing.T) {
	err := Download("https://www.baidu.com/", "./baidu.html")
	if err != nil {
		t.Error(err)
	}
}

func TestMd5sum(t *testing.T) {
	md5, err := Md5sum("./baidu.html")
	if err != nil {
		t.Error(err)
	}
	t.Log(md5)
}

func TestScanFiles(t *testing.T) {
	files, err := ScanFiles("../../")
	if err != nil {
		t.Error(err)
	}
	for i, f := range files{
		t.Log("idx:", i, " file:", f)
	}
}

func TestCopyFile(t *testing.T)  {
	if err := CopyFile("./file.go", "../file.go"); err != nil {
		t.Error(err)
	}
}

func TestCopyDir(t *testing.T) {
	if err := CopyDir("../", "../../copy/"); err != nil {
		t.Error(err)
	}
}

func TestRemoveAll(t *testing.T) {
	if err := RemoveAll("../../copy/"); err != nil {
		t.Error(err)
	}
}

func TestReadFile(t *testing.T) {
	content, err := ReadFile("./file.go")
	if err != nil {
		t.Error(content)
	}
	t.Log(content)

}