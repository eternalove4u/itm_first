package presenter

import (
	"bufio"
	"os"
)

type FilePresenter struct {
	pathToFile string
}

func NewFilePresenter(pathToFile string) *FilePresenter {
	return &FilePresenter{pathToFile}
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
