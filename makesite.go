package main

import (
	"flag"
	"fmt"
	"html/template"
	"os"
	"strings"
)

type Page struct {
	Content string
}

func save(filename string, page Page) error {
	t := template.Must(template.ParseFiles("template.tmpl"))
	out := strings.TrimSuffix(filename, ".txt") + ".html"
	f, err := os.Create(out)
	if err != nil {
		return err
	}
	defer f.Close()
	return t.Execute(f, page)
}

func main() {
	file := flag.String("file", "first-post.txt", "name of a .txt file to render")
	flag.Parse()

	data, err := os.ReadFile(*file)
	if err != nil {
		fmt.Println("error reading file:", err)
		os.Exit(1)
	}
	page := Page{Content: string(data)}

	t := template.Must(template.ParseFiles("template.tmpl"))
	if err := t.Execute(os.Stdout, page); err != nil {
		fmt.Println("error rendering template:", err)
		os.Exit(1)
	}

	if err := save(*file, page); err != nil {
		fmt.Println("error saving file:", err)
		os.Exit(1)
	}
}
