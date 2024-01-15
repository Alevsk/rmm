package mindmap

import (
	"strings"
)

// parseInputUsingStrategy parses the input using the given strategy
func parseInputUsingStrategy(input string, strategy parserStrategy) []string {
	return strategy.Parse(input)
}

// New reads lines from the input source and creates a mind map as a map[string]interface{}.
func New(source InputSource) (Node, error) {
	linesCh := source.ReadLines() // read lines from input source
	root := make(Node)
	for lineRes := range linesCh {
		if lineRes.Error != nil {
			return nil, lineRes.Error
		}
		if lineRes.Line == "" {
			continue
		}

		var strategy parserStrategy

		switch {
		case isIPv4(lineRes.Line):
			strategy = &ipv4Parsing{}
		default: // Default to hostname parsing
			strategy = &hostnameParsing{}
		}

		parts := parseInputUsingStrategy(lineRes.Line, strategy)

		currentNode := root
		for i := len(parts) - 1; i >= 0; i-- { // iterate over parts in reverse order
			key := parts[i]
			if !strings.Contains(parts[i], ".") {
				key = strings.Join(parts[i:], ".") // join parts into key
			}

			childNode, ok := currentNode[key] // get child node with key
			if !ok {
				childNode = make(Node) // create new child node if not found
				currentNode[key] = childNode
			}
			currentNode = childNode // set current node to child node
		}
	}
	return root, nil
}
