package common

import (
	"math/rand"
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

var CHARS = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
	"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}

func RandAllString(lenNum int) string {
	str := strings.Builder{}
	length := len(CHARS)
	for i := 0; i < lenNum; i++ {
		l := CHARS[rand.Intn(length)]
		str.WriteString(l)
	}
	return str.String()
}
