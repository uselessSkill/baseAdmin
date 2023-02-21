package loger

import (
	"os"
	"sync"
)

var lock sync.Mutex
var filePath = "./"

func WriteToFile(fileName string, writeData string) {
	lock.Lock()
	file, err := os.OpenFile(filePath+fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)

	if err != nil {
		panic("create log file error !! ")
		return
	}
	defer file.Close()
	defer lock.Unlock()
	_, err = file.WriteString(writeData + "\n")
	if err != nil {
		panic("write to log file error !! ")
		return
	}
}
