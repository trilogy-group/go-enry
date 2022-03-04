package loc

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/go-enry/go-enry/v2"
)

func GetLines(file string, content []byte) (total, blank int) {
	var r io.Reader
	if content != nil {
		r = bytes.NewReader(content)
	} else {
		// file not loaded to memory - stream it
		f, err := os.Open(file)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		r = f
	}
	br := bufio.NewReader(r)
	lastBlank := true
	empty := true
	for {
		data, prefix, err := br.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println(err)
			break
		}
		if prefix {
			continue
		}
		empty = false
		total++
		lastBlank = len(data) == 0
		if lastBlank {
			blank++
		}
	}
	if !empty && lastBlank {
		total++
		blank++
	}
	nonBlank := total - blank
	return total, nonBlank
}

func getFileType(file string, content []byte) string {
	switch {
	case enry.IsImage(file):
		return "Image"
	case enry.IsBinary(content):
		return "Binary"
	default:
		return "Text"
	}
}
