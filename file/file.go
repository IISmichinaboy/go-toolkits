package util

import (
	"github.com/ddliu/go-httpclient"
	"io/ioutil"
	"crypto"
	"os"
	"bufio"
	"io"
	"encoding/hex"
	"path/filepath"
)


// 下载文件
func Download(url string, filename string) error {
	res, err := httpclient.Get(url, nil)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	data, err := res.ReadAll()
	if err != nil {
		return err
	}
	err1 := ioutil.WriteFile(filename, data, 0755)
	if err1 != nil {
		return err1
	}
	return nil
}

// 删除所有目录或文件
func RemoveAll(path string) error {
	return os.RemoveAll(path)
}

func ReadFile(filename string) (string, error)  {
	data, err := ioutil.ReadFile(filename)
	if  err != nil {
		return "", err
	}
	return string(data), nil
}

// 计算md5
func Md5sum(filename string) (string, error) {
	md5 := crypto.MD5.New()
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()
	r := bufio.NewReader(f)

	_, err = io.Copy(md5, r)
	if err != nil {
		return "", err
	}
	temp := hex.EncodeToString(md5.Sum(nil))
	return temp, nil
}

// 扫描指定目录及其子目录,获取全部文件路径
func ScanFiles(dirPath string) ([]string, error)  {
	dirPath, err := filepath.Abs(dirPath)
	if err != nil {
		return nil, err
	}

	files := make([]string, 0, 10)
	dirs, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	PathSep := string(os.PathSeparator)
	for _, file := range dirs{
		if file.IsDir() {
			childFiles, err := ScanFiles(dirPath + PathSep + file.Name() + PathSep)
			if err != nil {
				return nil, err
			}
			files = append(files, childFiles...)
		} else {
			files = append(files,  dirPath + PathSep + file.Name())
		}
	}
	return files, nil
}

// 拷贝文件
func CopyFile(source string, dest string) error {
	sourcefile, err := os.Open(source)
	if err != nil {
		return err
	}
	defer sourcefile.Close()

	destfile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destfile.Close()

	_, err1 := io.Copy(destfile, sourcefile)
	if err1 != nil {
		return err1
	}

	sourceinfo, err := os.Stat(source)
	if err != nil {
		return err
	}

	if err := os.Chmod(dest, sourceinfo.Mode()); err != nil {
		return err
	}
	return nil
}

// 递归拷贝目录下的所有文件到指定目录下
func CopyDir(srcDir string, destDir string) error {
	// get properties of source dir
	sourceinfo, err := os.Stat(srcDir)
	if err != nil {
		return err
	}

	// create dest dir
	err = os.MkdirAll(destDir, sourceinfo.Mode())
	if err != nil {
		return err
	}

	directory, _ := os.Open(srcDir)
	defer directory.Close()

	objects, err := directory.Readdir(-1)

	for _, obj := range objects {

		sourcefilepointer := srcDir + "/" + obj.Name()
		destinationfilepointer := destDir + "/" + obj.Name()

		if obj.IsDir() {
			// create sub-directories - recursively
			err = CopyDir(sourcefilepointer, destinationfilepointer)
			if err != nil {
				return err
			}
		} else {
			err = CopyFile(sourcefilepointer, destinationfilepointer)
			if err != nil {
				return err
			}
		}

	}
	return nil
}
