package common

import (
	"log"
	"testing"
	"time"
)

func TestGetTimeUnix(t *testing.T) {
	got := GetTimeUnix()
	want := time.Now().Unix()

	if got != want {
		log.Fatal("TestGetTimeUnix ERROR!")
	}
}

func TestGetDateUnix(t *testing.T) {
	got := GetDateUnix()
	want := time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05")

	if got != want {
		log.Fatal("TestGetDateUnix ERROR!")
	}
}

func TestDateFromUnix(t *testing.T) {
	got := DateFromUnix(1620978453)
	want := time.Unix(1620978453, 0).Format("2006-01-02 15:04:05")

	if got != want {
		log.Fatal("TestDateFromUnix ERROR!")
	}
}
