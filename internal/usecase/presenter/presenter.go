package presenter

import (
	"bufio"
	"fmt"
	"os"
)

type FilePresenter struct {
	pathToFile string
}

func NewFilePresenter(pathToFile string) (*FilePresenter, error) {
	if pathToFile == "" {
		return nil, fmt.Errorf("empty file path")
	}
	return &FilePresenter{pathToFile}, nil
}

func (fp *FilePresenter) Present(messages []string) error {
	file, err := os.OpenFile(fp.pathToFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, message := range messages {
		_, err := writer.WriteString(message + "\n")
		if err != nil {
			return err
		}
	}
	return writer.Flush()
}
