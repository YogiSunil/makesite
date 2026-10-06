package main

import (
	"bytes"
	"flag"
	"fmt"
	"html"
	"html/template"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/yuin/goldmark"
)

type Page struct {
	Content template.HTML
}

func outputName(filename string) string {
	return strings.TrimSuffix(filename, filepath.Ext(filename)) + ".html"
}

// convert turns a .txt or .md source into an HTML fragment.
func convert(filename string, data []byte) (template.HTML, error) {
	if strings.HasSuffix(filename, ".md") {
		var buf bytes.Buffer
		if err := goldmark.Convert(data, &buf); err != nil {
			return "", err
		}
		return template.HTML(buf.String()), nil
	}
	return template.HTML(`<pre class="whitespace-pre-wrap font-sans bg-transparent">` + html.EscapeString(string(data)) + `</pre>`), nil
}

// render writes the page for one source file and returns the size of the HTML written.
func render(filename string, echo bool) (int64, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return 0, err
	}
	content, err := convert(filename, data)
	if err != nil {
		return 0, err
	}
	page := Page{Content: content}
	t := template.Must(template.ParseFiles("template.tmpl"))

	if echo {
		if err := t.Execute(os.Stdout, page); err != nil {
			return 0, err
		}
	}

	f, err := os.Create(outputName(filename))
	if err != nil {
		return 0, err
	}
	defer f.Close()
	if err := t.Execute(f, page); err != nil {
		return 0, err
	}
	info, err := f.Stat()
	if err != nil {
		return 0, err
	}
	return info.Size(), nil
}

func findFiles(dir string) ([]string, error) {
	var files []string
	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		if ext := filepath.Ext(path); ext == ".txt" || ext == ".md" {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

func main() {
	start := time.Now()
	file := flag.String("file", "first-post.txt", "name of a .txt or .md file to render")
	dir := flag.String("dir", "", "directory to search (recursively) for .txt and .md files")
	flag.Parse()

	files := []string{*file}
	if *dir != "" {
		var err error
		files, err = findFiles(*dir)
		if err != nil {
			fmt.Println("error reading directory:", err)
			os.Exit(1)
		}
		fmt.Println("Found files:")
		for _, f := range files {
			fmt.Println(" ", f)
		}
	}

	var total int64
	for _, f := range files {
		size, err := render(f, *dir == "")
		if err != nil {
			fmt.Println("error processing", f+":", err)
			os.Exit(1)
		}
		total += size
	}

	fmt.Printf("\033[1;32mSuccess!\033[0m Generated \033[1m%d\033[0m pages (%.1fkB total) in %.2f seconds.\n",
		len(files), float64(total)/1000, time.Since(start).Seconds())
}
