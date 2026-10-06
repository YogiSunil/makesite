# 🔗 makesite

[![Go Report Card](https://goreportcard.com/badge/github.com/YogiSunil/makesite)](https://goreportcard.com/report/github.com/YogiSunil/makesite)

A small static site generator written in Go. It reads `.txt` and `.md` files and turns each one into an HTML page using a Go template styled with Tailwind CSS (dark theme).

## Features

- Render a single file with `--file`.
- Render every `.txt` and `.md` file in a directory, including subdirectories, with `--dir`.
- Markdown (`.md`) files are converted to HTML with [goldmark](https://github.com/yuin/goldmark).
- Prints a summary when done, for example: `Success! Generated 7 pages (11.9kB total) in 0.01 seconds.`

## Project Structure

```
📂 makesite
├── makesite.go      # the generator
├── template.tmpl    # HTML template (Tailwind via CDN)
├── first-post.txt   # sample posts
├── latest-post.txt
├── second-post.txt
├── third-post.txt
├── fourth-post.txt
├── fifth-post.md    # sample Markdown post
├── go.mod
└── go.sum
```

## Requirements

- [Go](https://go.dev/dl/) 1.22 or newer.
- Internet access when viewing the pages, because Tailwind is loaded from a CDN.

## How to Run

```bash
git clone https://github.com/YogiSunil/makesite.git
cd makesite
```

Generate one page (defaults to `first-post.txt`):

```bash
go run makesite.go
go run makesite.go --file=latest-post.txt
```

Generate pages for every `.txt` and `.md` file in a directory (recursive):

```bash
go run makesite.go --dir=.
```

Or build a binary first:

```bash
go build
./makesite --dir=.        # Windows: .\makesite.exe --dir=.
```

Each input file produces an HTML file next to it, for example `latest-post.txt` becomes `latest-post.html`.

## View the Site

Serve the folder and open http://localhost:8000/first-post.html in a browser:

```bash
python -m http.server 8000
```

Or just double-click any generated `.html` file.

## How It Works

1. Read the input file.
2. Convert it to an HTML fragment (`.txt` is escaped and wrapped in `<pre>`, `.md` goes through goldmark).
3. Render the fragment into `template.tmpl` with Go's `html/template`.
4. Write the result to `<name>.html`.

Edit `template.tmpl` to change the layout and styling, then run the command again to regenerate the pages.

## Progress

- [x] v1.0: read a file, render it with a Go template, save it as HTML, add the `--file` flag
- [x] v1.0 stretch: styled the template with Tailwind CSS
- [x] v1.1: add the `--dir` flag and generate a page for each file
- [x] v1.1 stretch: recursive search, colored success message, total size and elapsed time
- [x] v1.2: Go modules and a third-party library

### v1.2 library

I will use the goldmark library. The documentation is located at https://github.com/yuin/goldmark. My goal is to use it to convert Markdown (.md) files into HTML.
