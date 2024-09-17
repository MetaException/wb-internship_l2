package utils

/*

import (
	"bufio"
	"io"
	"os"
)

type TextReader struct {
	sc *bufio.Scanner
}

func NewTextReader() *TextReader {
	sc := bufio.NewScanner(os.Stdin)
	return &TextReader{
		sc: sc,
	}
}

func (tr *TextReader) ReadLine() (string, error) {

	if tr.sc.Scan() {
		return tr.sc.Text(), nil
	}

	if err := tr.sc.Err(); err != nil {
		return "", err
	}

	return "", io.EOF
}

func (tr *TextReader) ReadText() ([]string, error) {
	text := make([]string, 0)

	var line string
	var err error
	for ; err != io.EOF; line, err = tr.ReadLine() {

		if err != nil {
			return nil, err
		}
		text = append(text, line)
	}

	return text, nil
}
*/
