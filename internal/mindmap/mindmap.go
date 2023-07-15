package mindmap

import (
	"bufio"
	"net/url"
	"os"
	"strings"
)

type InputSource interface {
	ReadLines() <-chan LineResult
}

type LineResult struct {
	Line  string
	Error error
}

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

type Domain struct {
	Name       string
	SubDomains []*Domain
}

// parseHostname takes a string input representing a URL and returns the hostname and path of the URL.
func parseHostname(input string) (string, string) {
	if !strings.Contains(input, "://") {
		input = "http://" + input
	}
	u, err := url.Parse(input)
	if err != nil {
		return "", ""
	}
	return u.Host, u.Path
}

type Node map[string]Node

// CreateMindMap reads lines from the input source and creates a mind map as a map[string]interface{}.
func CreateMindMap(source InputSource) (Node, error) {
	linesCh := source.ReadLines() // read lines from input source
	root := make(Node)
	for lineRes := range linesCh {
		if lineRes.Error != nil {
			return nil, lineRes.Error
		}
		if lineRes.Line == "" {
			continue
		}
		hostname, _ := parseHostname(lineRes.Line)
		parts := strings.Split(hostname, ".") // split domain by dot
		currentNode := root
		for i := len(parts) - 1; i >= 0; i-- { // iterate over parts in reverse order
			key := strings.Join(parts[i:], ".") // join parts into key
			childNode, ok := currentNode[key]   // get child node with key
			if !ok {
				childNode = make(Node) // create new child node if not found
				currentNode[key] = childNode
			}
			currentNode = childNode // set current node to child node
		}
	}
	return root, nil
}
