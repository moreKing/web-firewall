package dir

import (
	"crypto/md5"
	"io"
	"io/fs"
	"io/ioutil"
	"os"
	"path"
)

type FileInfo struct {
	Name    string
	Size    int64
	MD5     []byte
	ModTime int64
}

func GetFileInfo(filepath string) (*FileInfo, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	info, err := f.Stat()
	if err != nil {
		return nil, err
	}

	md5hash := md5.New()
	if _, err := io.Copy(md5hash, f); err != nil {
		// fmt.Println("Copy", err)
		return nil, err
	}

	fileInfo := &FileInfo{
		Name:    info.Name(),
		Size:    info.Size(),
		MD5:     md5hash.Sum(nil),
		ModTime: info.ModTime().Unix(),
	}

	return fileInfo, err

}

// 获取指定目录下所有文件，包括子目录
func GetDirAllFileInfo(fullPath string) []fs.FileInfo {
	files, err := ioutil.ReadDir(fullPath)
	if err != nil {
		return nil
	}
	var infos []fs.FileInfo
	for _, v := range files {
		// 判断是否是文件夹
		if v.IsDir() {
			newPath := path.Join(fullPath, v.Name())
			subFile := GetDirAllFileInfo(newPath)
			if len(subFile) > 0 {
				infos = append(infos, subFile...)
			}
		} else {
			infos = append(infos, v)
		}
	}
	return infos
}

// 获取指定目录下所有文件，不包括子目录
func GetDirFileInfo(fullPath string) []os.DirEntry {
	files, err := os.ReadDir(fullPath)
	if err != nil {
		return nil
	}
	var infos []os.DirEntry
	for _, v := range files {
		// 判断是否是文件夹
		if !v.IsDir() {
			infos = append(infos, v)
		}
	}
	return infos
}

func DirExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func CreateDir(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

func CreateAllFile(path, filename string) {
	_ = os.MkdirAll(path, os.ModePerm)
	_, _ = os.Create(path + filename)
}
