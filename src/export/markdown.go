package export

var Markdown = `# {{ .Name }}
{{ .Description }}
{{ range $slKey, $slValues := .Storylines }}
- [{{ $slKey }}](#{{ $slKey | replace " " "-" }}){{ range $questKey, $questVal := $slValues.Quests }}
  - [{{ $questKey }}](#{{ $questKey | replace " " "-" }}){{ end }}{{end}}
{{ range $slKey, $slValues := .Storylines }}
## {{ $slKey }}
{{ range $questKey, $questValues :=  $slValues.Quests }}
### {{ $questKey }}
{{ range $task := $questValues.Tasks }}- [ ] {{ $task.Name }}
  - {{ $task.Description }}
  - {{ $task.Action }}
{{ end }}{{ end }}{{ end }}

# Resources
{{ range $resource := .Resources }}
## [{{ $resource.Name }}]({{$resource.URL}})
{{ $resource.Description }}
{{ end }}
`
