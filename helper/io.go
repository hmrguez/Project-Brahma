package helper

import (
	"log"
	"os"
)

func CreateFile(name string, content string) {
	file, err := os.Create(name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	file.WriteString(content)
}

func ReadFile(name string) string{
	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	buf := make([]byte, 1024)
	n, err := file.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	return string(buf[:n])
}