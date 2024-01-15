package mindmap

import "bufio"

type ScannerInput struct {
	Scanner *bufio.Scanner
}

func (s ScannerInput) ReadLines() <-chan LineResult {
	linesCh := make(chan LineResult)
	go func() {
		defer close(linesCh)

		for s.Scanner.Scan() {
			linesCh <- LineResult{Line: s.Scanner.Text()}
		}

		if s.Scanner.Err() != nil {
			linesCh <- LineResult{Error: s.Scanner.Err()}
		}
	}()
	return linesCh
}
