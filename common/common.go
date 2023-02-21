package common

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

func GetTimeUnix() int64 {
	return time.Now().Unix()
}

func GetDateUnix() string {
	return time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05")
}

func DateFromUnix(Unix int64) string {
	return time.Unix(Unix, 0).Format("2006-01-02 15:04:05")
}

func GetCurrPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))
	ret := path[:index]
	return ret
}
