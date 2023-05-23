package mindmap

import (
	"bufio"
	"os"
	"strings"
)

type InputSource interface {
	ReadLines() ([]string, error)
}

type FileInput struct {
	FilePath string
}

func (f FileInput) ReadLines() ([]string, error) {
	file, err := os.Open(f.FilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

type ScannerInput struct {
	Scanner *bufio.Scanner
}

func (s ScannerInput) ReadLines() ([]string, error) {
	var lines []string
	for s.Scanner.Scan() {
		lines = append(lines, s.Scanner.Text())
	}

	return lines, s.Scanner.Err()
}

type Domain struct {
	Name       string
	SubDomains []*Domain
}

type Node map[string]Node

// CreateMindMap reads lines from the input source and creates a mind map as a map[string]interface{}.
func CreateMindMap(source InputSource) (Node, error) {
	lines, err := source.ReadLines() // read lines from input source
	if err != nil {
		return nil, err
	}
	root := make(Node)
	for _, domain := range lines { // for each domain in the input lines
		if domain == "" {
			continue
		}
		parts := strings.Split(domain, ".") // split domain by dot
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
