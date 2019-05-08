package main

var pkgTemplate = `{{with .PDoc}}

{{comment_md .Doc}}

{{example_blocks $ }}

{{end}}
`
