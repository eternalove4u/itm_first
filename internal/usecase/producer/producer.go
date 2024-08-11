package producer

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type FileProducer struct {
	pathToFile string
}

func NewFileProducer(pathToFile string) (*FileProducer, error) {
	if pathToFile == "" {
		return nil, fmt.Errorf("empty path in new producer")
	}
	return &FileProducer{pathToFile}, nil
}

func (fp *FileProducer) Produce() ([]string, error) {
	file, err := os.Open(fp.pathToFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	text, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("read input file: %w", err)
	}

	messages := strings.Split(string(text), "\n")
	for i, m := range messages {
		if len(m) == 0 {
			continue
		}
		if m[len(m)-1] == '\n' {
			messages[i] = m[:len(m)-2]
		}
	}

	return messages, nil
}
