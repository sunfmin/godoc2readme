// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// # godoc to README.md
// godoc2readme converts godoc formatted package documentation into Markdown format that can be used directly as github README.md.
//
//
// Usage
//
//    godoc2readme github.com/sunfmin/godoc2readme > README.md
package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/build"
	"go/printer"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"text/template"
	"unicode"
	"unicode/utf8"

	"golang.org/x/tools/godoc"
	"golang.org/x/tools/godoc/vfs"
)

var (
	verbose = flag.Bool("v", false, "verbose mode")

	// file system roots
	// TODO(gri) consider the invariant that goroot always end in '/'
	goroot = flag.String("goroot", runtime.GOROOT(), "Go root directory")

	// layout control
	tabWidth       = flag.Int("tabwidth", 4, "tab width")
	showTimestamps = flag.Bool("timestamps", false, "show timestamps with directory listings")
	templateDir    = flag.String("templates", "", "directory containing alternate template files")
	showPlayground = flag.Bool("play", false, "enable playground in web interface")
	// showExamples   = flag.Bool("ex", false, "show examples in command line mode")
	declLinks = flag.Bool("links", true, "link identifiers to their declarations")
)

func usage() {
	fmt.Fprintf(os.Stderr,
		"usage: godoc2md package [name ...]\n")
	flag.PrintDefaults()
	os.Exit(2)
}

var (
	pres *godoc.Presentation
	fs   = vfs.NameSpace{}

	funcs = map[string]interface{}{
		"comment_md":     comment_mdFunc,
		"base":           path.Base,
		"md":             mdFunc,
		"pre":            preFunc,
		"title":          title,
		"example_blocks": example_blocks,
		"archor":         archor,
	}
)

const punchCardWidth = 80

func startsWithUppercase(s string) bool {
	r, _ := utf8.DecodeRuneInString(s)
	return unicode.IsUpper(r)
}

// stripExampleSuffix strips lowercase braz in Foo_braz or Foo_Bar_braz from name
// while keeping uppercase Braz in Foo_Braz.
func stripExampleSuffix(name string) string {
	if i := strings.LastIndex(name, "_"); i != -1 {
		if i < len(name)-1 && !startsWithUppercase(name[i+1:]) {
			name = name[:i]
		}
	}
	return name
}

func example_blocks(info *godoc.PageInfo, funcName string) string {

	var buf bytes.Buffer
	first := true
	for _, eg := range info.Examples {
		name := stripExampleSuffix(eg.Name)
		if name != funcName {
			continue
		}

		if !first {
			buf.WriteString("\n")
		}
		first = false

		// print code
		cnode := &printer.CommentedNode{Node: eg.Code, Comments: eg.Comments}
		var buf1 bytes.Buffer
		pres.WriteNode(&buf1, info.FSet, cnode)
		code := buf1.String()
		// Additional formatting if this is a function body.
		if n := len(code); n >= 2 && code[0] == '{' && code[n-1] == '}' {
			// remove surrounding braces
			code = code[1 : n-1]
			// unindent
			code = strings.Replace(code, "\n    ", "\n", -1)
		}
		code = strings.Trim(code, "\n")
		code = strings.Replace(code, "\n", "\n\t", -1)

		buf.WriteString(eg.Doc)
		buf.WriteString("```go\n\t")
		buf.WriteString(code)
		buf.WriteString("\n```\n")
	}
	return buf.String()
}

func title(funcName string) string {
	inputs := []rune(funcName)
	results := []rune{}
	var isPrevUpper = false
	for _, c := range inputs {
		if c == '*' {
			continue
		}
		if len(results) == 0 {
			results = append(results, c)
			isPrevUpper = unicode.IsUpper(c)
			continue
		}
		if unicode.IsUpper(c) && !isPrevUpper {
			results = append(results, ' ')
			results = append(results, c)
			isPrevUpper = true
			continue
		}
		results = append(results, c)
		isPrevUpper = false
	}
	return string(results)
}

func archor(funcNames ...string) string {
	rs := []string{}
	for _, funcName := range funcNames {
		rs = append(rs, strings.Replace(strings.ToLower(title(funcName)), " ", "-", -1))
	}
	return strings.Join(rs, "-")
}

func comment_mdFunc(comment string) string {
	var buf bytes.Buffer
	toMD(&buf, comment, nil)
	return buf.String()
}

func mdFunc(text string) string {
	text = strings.Replace(text, "*", "\\*", -1)
	text = strings.Replace(text, "_", "\\_", -1)
	return text
}

func preFunc(text string) string {
	return "``` go\n" + text + "\n```"
}

func readTemplate(name, data string) *template.Template {
	// be explicit with errors (for app engine use)
	t, err := template.New(name).Funcs(pres.FuncMap()).Funcs(funcs).Parse(string(data))
	if err != nil {
		log.Fatal("readTemplate: ", err)
	}
	return t
}

func readTemplates(p *godoc.Presentation, html bool) {
	p.PackageText = readTemplate("package.txt", pkgTemplate)
}

func main() {
	flag.Usage = usage
	flag.Parse()

	// Check usage
	if flag.NArg() == 0 {
		usage()
	}

	// use file system of underlying OS
	fs.Bind("/", vfs.OS(*goroot), "/", vfs.BindReplace)

	// Bind $GOPATH trees into Go root.
	for _, p := range filepath.SplitList(build.Default.GOPATH) {
		fs.Bind("/src/pkg", vfs.OS(p), "/src", vfs.BindAfter)
	}

	corpus := godoc.NewCorpus(fs)
	corpus.Verbose = *verbose

	pres = godoc.NewPresentation(corpus)
	pres.TabWidth = *tabWidth
	pres.ShowTimestamps = *showTimestamps
	pres.ShowPlayground = *showPlayground
	pres.ShowExamples = true
	pres.DeclLinks = *declLinks
	pres.SrcMode = false
	pres.HTMLMode = false

	readTemplates(pres, false)

	if err := godoc.CommandLine(os.Stdout, fs, pres, flag.Args()); err != nil {
		log.Print(err)
	}
}
