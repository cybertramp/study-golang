package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	filename := ""
	var fileInfo os.FileInfo
	argv := os.Args
	if len(argv) == 1 {
		log.Fatal("파일명을 입력해주세요!")
	} else {
		filename = argv[1]
	}

	// Get file size
	fileInfo, err := os.Stat(filename)
	if err != nil {
		log.Fatal(err)
	}

	filesize := fileInfo.Size()

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	buf := make([]byte, filesize)

	for {
		n, err := file.Read(buf)
		if n > 0 {
			fmt.Printf("%s", buf)
		}

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Printf("Read %d Bytes: %v", n, err)
			break
		}
	}

}
