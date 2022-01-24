{{- /*gotype:github.com/Necoro/engarde-importer.EngardeConfig*/ -}}
{{ range .Clubs -}}
    {[classe club] [nom "{{.Name}}"] [modifie vrai] [date_oed "332"] [cle {{.Id}}]}
{{ end }}