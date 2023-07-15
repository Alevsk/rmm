package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"github.com/Alevsk/rmm/internal/cli"
	"github.com/Alevsk/rmm/internal/mindmap"
	"github.com/Alevsk/rmm/internal/sys"
	flag "github.com/spf13/pflag"
)

type commands = map[string]func([]string)

const usage = `Usage:
    rmm [options] <command>

Commands:
    server                   Start a RMM server (TODO).
    update                   Update RMM binary (TODO).

Options:
    -f, --file               Filename from where to read input.
    -h, --help               Print command line options.
    -v, --version            Print version information.
    -o, --output             Display result in different formats list|markdown|json|yaml|obsidian (default: list)
`

func main() {

	cmd := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	cmd.Usage = func() { fmt.Fprint(os.Stderr, usage) }

	var (
		showVersion  bool
		fileName     string
		outputFormat string
	)

	cmd.StringVarP(&fileName, "file", "f", "", "Filename from where to read input.")
	cmd.BoolVarP(&showVersion, "version", "v", false, "Print version information.")
	cmd.StringVarP(&outputFormat, "output", "o", "list", "Display result in different formats list|markdown|json|yaml|obsidian (default: list)")

	if err := cmd.Parse(os.Args); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			os.Exit(2)
		}
		cli.Fatalf("%v. See 'rmm --help'", err)
	}

	if showVersion {
		buildInfo := sys.BinaryInfo()
		cli.Printf("rmm %s (commit=%s)\n", buildInfo.Version, buildInfo.CommitID)
		return
	}

	stat, err := os.Stdin.Stat()
	if err != nil {
		cli.Fatalf("%v. See 'rmm --help'", err)
	}

	var source mindmap.InputSource

	if fileName != "" {
		source = mindmap.FileInput{
			FilePath: fileName,
		}
	} else if (stat.Mode() & os.ModeCharDevice) == 0 {
		source = mindmap.ScannerInput{
			Scanner: bufio.NewScanner(os.Stdin),
		}
	}

	if source != nil {
		tree, err := mindmap.CreateMindMap(source)
		if err != nil {
			cli.Fatalf("%v. See 'rmm --help'", err)
		}
		switch outputFormat {
		case "json":
			cli.PrintJSON(tree)
		case "yaml":
			cli.PrintYAML(tree)
		case "markdown":
			cli.PrintMarkdown(tree)
		case "obsidian":
			cli.PrintObsidianCanvas(tree)
		case "list":
			cli.PrintList(tree)
		default:
			cli.PrintList(tree)
		}
		return
	}

	subCmds := commands{
		"server": func(s []string) { fmt.Println("running server") },
		"update": func(s []string) { fmt.Println("updating binary") },
	}
	if len(os.Args) > 1 {
		if subCmd, ok := subCmds[os.Args[1]]; ok {
			subCmd(os.Args[1:])
			return
		}
	}
	if cmd.NArg() > 1 {
		cli.Fatalf("%q is not a rmm command. See 'rmm --help'", cmd.Arg(1))
	}

	cmd.Usage()
	os.Exit(2)

}
