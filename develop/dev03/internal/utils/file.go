package utils

import (
	"bufio"
	"os"
)

func ReadText(path string) ([]string, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	text := make([]string, 0)

	for sc.Scan() {
		text = append(text, sc.Text())
	}

	if err := sc.Err(); err != nil {
		return nil, err
	}

	return text, nil
}

func WriteTextToFilePath(text []string, path string) error {

	newFile, err := os.Create(path)
	if err != nil {
		return err
	}
	defer newFile.Close()

	WriteText(text, newFile)

	return nil
}

func WriteText(text []string, dist *os.File) error {

	writer := bufio.NewWriter(dist)
	for _, line := range text {
		_, err := writer.WriteString(line + "\n") //..
		if err != nil {
			return err
		}
	}
	writer.Flush()

	return nil
}
