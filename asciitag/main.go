package main

import (
	"flag"
	"io"
	"log"
	"os"
	"strings"

	"github.com/ojn/asciitag"
	"github.com/ojn/asciitag/flag"
)

// Flags
var (
	assignees = flag.String("a", "", "Only TAGs assigned to this username(s) will be displayed")
	format    = flag.String("f", "text", "Format to use when outputting TAGs (supported formats: text, csv, json)")
	output    = flag.String("o", "stdout", "Destination for output (can be stdout, stderr or a file)")
	exclude   = astiflag.Strings{}
)

func main() {
	// Parse flags
	flag.Var(&exclude, "e", "Path that will be excluded from the process")
	flag.Parse()

	// Loop through paths
	for _, path := range flag.Args() {
		// Process path
		var tags asciitag.TAGs
		var err error
		if tags, err = asciitag.Extract(path, exclude...); err != nil {
			log.Fatal(err)
		}

		// Filter results for assignee
		if *assignees != "" {
			tags = tags.AssignedTo(strings.Split(*assignees, ",")...)
		}

		var writer io.Writer

		// Convert selected output into a writer
		switch *output {
		case "stdout":
			writer = os.Stdout
		case "stderr":
			writer = os.Stderr
		default:
			if writer, err = os.OpenFile(*output, os.O_WRONLY|os.O_CREATE, 0666); err != nil {
				log.Fatal(err)
			}

			defer writer.(*os.File).Close()
		}

		// Handle selected format
		switch *format {
		case "text":
			if err = tags.WriteText(writer); err != nil {
				log.Fatal(err)
			}
		case "csv":
			if err = tags.WriteCSV(writer); err != nil {
				log.Fatal(err)
			}
		case "json":
			if err = tags.WriteJSON(writer); err != nil {
				log.Fatal(err)
			}
		default:
			log.Fatalf("unsupported format: %s", *format)
		}
	}
}
