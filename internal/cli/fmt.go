package cli

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"

	tui "github.com/charmbracelet/lipgloss"
	"gopkg.in/yaml.v2"
)

var errPrefix = tui.NewStyle().Foreground(tui.Color("#ac0000")).Render("Error: ")

func Fatal(v ...any) {
	fmt.Fprint(os.Stderr, errPrefix)
	fmt.Fprint(os.Stderr, v...)
	fmt.Fprintln(os.Stderr)
	os.Exit(1)
}

func Fatalf(format string, v ...any) {
	fmt.Fprintf(os.Stderr, errPrefix+format+"\n", v...)
	os.Exit(1)
}

func Print(v ...any) (int, error) { return fmt.Print(v...) }

func Printf(format string, v ...any) (int, error) { return fmt.Printf(format, v...) }

func Println(v ...any) (int, error) { return fmt.Println(v...) }

func PrintYAML(tree map[string]interface{}) {
	yamlBytes, err := yaml.Marshal(tree)
	if err != nil {
		Fatalf("%v. See 'rmm --help'", err)
	}

	fmt.Println(string(yamlBytes))
}

// PrintJSON takes in a map of type string to interface{}, which represents a JSON tree.
func PrintJSON(tree map[string]interface{}) {
	// Marshal the JSON tree into a formatted string with the given indentation prefix and spaces.
	data, err := json.MarshalIndent(tree, "", "  ")
	if err != nil {
		// If an error occurred during marshalling, exit with a fatal error message with the given error message.
		Fatalf("%v. See 'rmm --help'", err)
	}
	// Print the JSON string.
	fmt.Println(string(data))
}

// TreeToList takes in a map of type string to interface{} and returns a string that represents a list of the key-value pairs in the map.
// The map can be nested, and its keys will be sorted alphabetically.
// The list will have a nested structure that reflects the structure of the map.
func TreeToList(data map[string]interface{}, markdown bool) string {
	// Define a struct that represents a key-value pair and its prefix in the markdown list.
	type entry struct {
		key    string
		prefix string
		value  interface{}
	}

	// Initialize a stack with the top level of the map.
	stack := []entry{{"", "", data}}
	// Initialize a string builder for the result.
	var result strings.Builder

	// Loop through the stack until it is empty.
	for len(stack) > 0 {
		// Pop the last element from the stack.
		e := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// Get the key, prefix, and value from the popped element.
		key := e.key
		prefix := e.prefix
		value := e.value

		switch v := value.(type) {
		case map[string]interface{}:
			// If the value is a map, add its prefix and key (if not empty) to the result as a markdown list item.
			if key != "" {
				listSymbol := ""
				if markdown {
					listSymbol = "- "
				}
				result.WriteString(prefix + listSymbol + key + "\n")
			}

			// Sort the keys of the nested map.
			keys := make([]string, 0, len(v))
			for k := range v {
				keys = append(keys, k)
			}
			sort.Strings(keys)
			// Add the entries to the stack in reverse order to preserve the order of the keys.
			for i := len(keys) - 1; i >= 0; i-- {
				k := keys[i]
				if key == "" && prefix == "" {
					// If the key and prefix are both empty, add the key-value pair to the stack with an empty prefix.
					stack = append(stack, entry{k, "", v[k]})
				} else {
					// Otherwise, add the key-value pair to the stack with a prefix that is two spaces more than its parent.
					indentation := "\t"
					if markdown {
						indentation = "  "
					}
					stack = append(stack, entry{k, prefix + indentation, v[k]})
				}
			}
		}
	}

	// Return the list.
	return result.String()
}

// PrintMarkdown output the tree in markdown format using an unsorted list
func PrintMarkdown(tree map[string]interface{}) {
	fmt.Println(TreeToList(tree, true))
}

// PrintList output the tree in plaintext format using tabs
func PrintList(tree map[string]interface{}) {
	fmt.Println(TreeToList(tree, false))
}
