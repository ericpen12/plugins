package file

import (
	"bufio"
	"io"
	"os"
)

func Mkdir(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return os.Mkdir(dir, os.ModePerm)
	}
	return nil
}

func ReadAsLine(filename string, handle func(line string, err error)) {
	f, err := os.Open(filename)
	if err != nil {
		handle("", err)
		return
	}
	buf := bufio.NewReader(f)
	for {
		b, _, err := buf.ReadLine()
		if err == io.EOF {
			break
		}
		handle(string(b), err)
	}
}
