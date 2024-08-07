package producer

import (
	"bufio"
	"os"
)

type FileProducer struct {
	pathToFile string
}

func NewFileProducer(pathToFile string) *FileProducer {
	return &FileProducer{pathToFile}
}

func (fp *FileProducer) Produce() ([]string, error) {
	file, err := os.Open(fp.pathToFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var messages []string
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		messages = append(messages, scanner.Text())
	}
	return messages, nil
}
