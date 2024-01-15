package mindmap

import (
	"bufio"
	"os"
)

type FileInput struct {
	FilePath string
}

func (f FileInput) ReadLines() <-chan LineResult {
	linesCh := make(chan LineResult)
	go func() {
		defer close(linesCh)
		file, err := os.Open(f.FilePath)
		if err != nil {
			linesCh <- LineResult{Error: err}
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			linesCh <- LineResult{Line: scanner.Text()}
		}

		if scanner.Err() != nil {
			linesCh <- LineResult{Error: scanner.Err()}
		}
	}()
	return linesCh
}
