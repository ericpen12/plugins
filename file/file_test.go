package file

import (
	"testing"
)

func TestReadLine(t *testing.T) {
	count := 1
	ReadLine("file.go", func(line string, err error) {
		if err != nil {
			t.Error(err)
			return
		}
		t.Logf("count: %d, content: %s", count, line)
		count++
	})
}
