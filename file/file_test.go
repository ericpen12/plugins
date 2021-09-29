package file

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	ReadLine("file.go", func(line string, err error) {
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(line)
	})
}
