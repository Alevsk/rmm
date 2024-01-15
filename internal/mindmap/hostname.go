package mindmap

import (
	"net/url"
	"strings"
)

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

// hostnameParsing parses hostnames
type hostnameParsing struct{}

func (p *hostnameParsing) Parse(input string) []string {
	hostname, _ := parseHostname(input)
	return strings.Split(hostname, ".")
}
