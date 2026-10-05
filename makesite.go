package main

import (
	"flag"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
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

func render(file string) error {
	data, err := os.ReadFile(file)
	if err != nil {
		return err
	}
	page := Page{Content: string(data)}

	t := template.Must(template.ParseFiles("template.tmpl"))
	if err := t.Execute(os.Stdout, page); err != nil {
		return err
	}
	return save(file, page)
}

func main() {
	file := flag.String("file", "first-post.txt", "name of a .txt file to render")
	dir := flag.String("dir", "", "directory to search for .txt files")
	flag.Parse()

	files := []string{*file}
	if *dir != "" {
		entries, err := os.ReadDir(*dir)
		if err != nil {
			fmt.Println("error reading directory:", err)
			os.Exit(1)
		}
		files = nil
		for _, e := range entries {
			if !e.IsDir() && strings.HasSuffix(e.Name(), ".txt") {
				files = append(files, filepath.Join(*dir, e.Name()))
			}
		}
		fmt.Println("Found .txt files:")
		for _, f := range files {
			fmt.Println(" ", f)
		}
	}

	for _, f := range files {
		if err := render(f); err != nil {
			fmt.Println("error processing", f+":", err)
			os.Exit(1)
		}
	}
}
