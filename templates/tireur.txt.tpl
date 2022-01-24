{{- /*gotype:github.com/Necoro/engarde-importer.EngardeConfig*/ -}}
{{ range .Participants -}}
    {[classe tireur] [sexe {{.Gender.Engarde}}] [presence present] [carton_coach non] [status normal]
    [medical non] [lateralite droite] [nom " {{.LastName | upper}} "] [prenom " {{.FirstName}} "]
    [points 1.0] [date_oed "340"] [cle {{.Id}}] [club1 {{.ClubId}}]}
{{ end }}