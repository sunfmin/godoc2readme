package main

var pkgTemplate = `{{with .PDoc}}

{{comment_md .Doc}}

{{range .Funcs}}## {{title .Name}}
{{node $ .Decl | pre}}
{{comment_md .Doc}}
{{example_blocks $ .Name}}
{{callgraph_html $ "" .Name}}{{end}}

{{range .Types}}{{$tname := .Name}}## Type {{title .Name}}
{{node $ .Decl | pre}}
{{comment_md .Doc}}{{range .Consts}}
{{node $ .Decl | pre }}
{{comment_md .Doc}}{{end}}{{range .Vars}}
{{node $ .Decl | pre }}
{{comment_md .Doc}}{{end}}

{{example_blocks $ $tname}}
{{implements_html $ $tname}}
{{methodset_html $ $tname}}

{{range .Funcs}}{{$name_html := html .Name}}### func {{$name_html}}
{{node $ .Decl | pre}}
{{comment_md .Doc}}
{{example_blocks $ .Name }}{{end}}
{{callgraph_html $ "" .Name}}

{{range .Methods}}### {{title .Recv|md}}'s {{title .Name}}
{{node $ .Decl | pre}}
{{comment_md .Doc}}
{{$name := printf "%s_%s" $tname .Name}}{{example_blocks $ $name}}
{{callgraph_html $ .Recv .Name}}
{{end}}{{end}}{{end}}

`
