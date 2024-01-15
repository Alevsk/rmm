package mindmap

type LineResult struct {
	Line  string
	Error error
}

type Domain struct {
	Name       string
	SubDomains []*Domain
}

type Node map[string]Node

type InputSource interface {
	ReadLines() <-chan LineResult
}

type parserStrategy interface {
	Parse(string) []string
}
