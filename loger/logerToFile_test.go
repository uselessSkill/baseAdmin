package loger

import "testing"

func TestWriteToFile(t *testing.T) {
	WriteToFile("king.txt", "demo")
}
