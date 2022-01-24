{{- /*gotype:github.com/Necoro/engarde-importer.EngardeConfig*/ -}}
(def competition
  type competition
  titre_reduit {[type chaine] [taillecar 16] [nomchamp titre_reduit] [texte titre_reduit]}
  titre_ligne {[type chaine] [taillecar 40] [nomchamp titre_ligne] [texte titre_ligne]}
  organisateur {[type chaine] [taillecar 16] [nomchamp organisateur] [texte organisateur]}
  federation {[type chaine] [taillecar 16] [nomchamp federation] [texte federation]}
  domaine_compe {[type domaine_compe] [taillecar 15] [defaut international] [nomchamp domaine_compe]
 [texte domaine_compe] [enumere (("i" international "international") ("n" national
 "national") ("r" regional "regional") ("l" local "local"))]}
  championnat {[type chaine] [taillecar 16] [nomchamp championnat] [texte championnat]}
  id {[type chaine] [taillecar 8] [nomchamp id] [texte id]}
  annee {[type chaine] [taillecar 10] [nomchamp annee] [texte annee]}
  arme {[type arme] [taillecar 10] [nomchamp arme] [texte arme] [enumere (("F" fleuret "Florett")
 ("D" epee "Degen") ("S" sabre "Säbel"))]}
  categorie {[type categorie] [taillecar 10] [nomchamp categorie] [texte categorie] [enumere
 (("S" veteran "Senioren") ("A" senior "Aktive") ("J" junior "Junioren") ("K" cadet
 "Kadetten") ("u" minime "Jugend B aelter") ("g" benjamin "Jugend B juenger") ("c"
 pupille "Schueler aelter") ("h" poussin "Schueler juenger"))]}
  sexe {[type sexe] [taillecar 10] [defaut masculin] [nomchamp sexe] [texte sexe] [enumere
 (("H" masculin "Herren") ("D" feminin "Damen"))]}
  date {[type date] [taillecar 10] [nomchamp date] [texte date]}
  titre1 {[type chaine] [taillecar 40] [nomchamp titre1] [texte titre1]}
  titre2 {[type chaine] [taillecar 40] [nomchamp titre2] [texte titre2]}
  titre3 {[type chaine] [taillecar 40] [nomchamp titre3] [texte titre3]}
  titre4 {[type chaine] [taillecar 40] [nomchamp titre4] [texte titre4]}
  entites_premieres (tireur arbitre)
  tables (tireur arbitre club ligue nation)
  typetype entite
  champs_affiches (titre_reduit titre_ligne organisateur federation domaine_compe championnat id annee
 arme categorie sexe date titre1 titre2 titre3 titre4)
  tous_les_champs (titre_reduit titre_ligne organisateur federation domaine_compe championnat id annee
 arme categorie sexe date titre1 titre2 titre3 titre4)
  texte_ligue etat
  modifie vrai
)

(def ma_competition
  classe competition
  domaine_compe international
  sexe {{.Gender.Engarde}}
  type_compe individuelle
  contexte {[federation ""] [texte_ligue etat] [info_categorie ((veteran 41 120) (senior 21
 40) (junior 18 20) (cadet 16 17) (minime 14 15) (benjamin 12 13) (pupille 10 11)
 (poussin 8 9))] [fichiers (ficnations)] [ficnations {[type nation] [chemin "nations.txt"]
 [info_champs ((nom 1 3) (nom_etendu 5 20))]}]}
  federation "HFV"
  version "                ENGARDE Version 8.16 - 11/11/2005"
  version_fic 80
  titre_reduit "{{.Name}}"
  titre_ligne "{{.Description}}"
  organisateur "ETV"
  annee "{{.Date.Year}}"
  categorie {{.AgeGroup.Engarde}}
  date "~{{.Date.Format "02/01/2006"}}"
  arme {{.Weapon.Engarde}}
  tireur ()
  arbitre ()
  club ()
  ligue ()
  nation ()
)

(def formule
  type formule
  nombre_entites {[type entier] [taille 32] [taillecar 3] [defaut 100] [nomchamp nombre_entites] [texte
 nombre_entites]}
  clasmt_origine {[type serie_points] [taille 50] [taillecar 7] [defaut serie] [nomchamp clasmt_origine]
 [texte clasmt_origine] [enumere (("S" serie "Systemnr.") ("P" points "Punkte"))]}
  entree_progressive {[type oui_non] [taille 40] [taillecar 3] [defaut non] [nomchamp entree_progressive]
 [texte entree_progressive] [enumere (("j" oui "ja") ("n" non "nein"))]}
  nombre_touches_poules {[type entier] [defaut 5] [nomchamp nombre_touches_poules] [texte nombre_touches_poules]}
  qual_fin_poules_sur {[type entier] [taille 32] [taillecar 3] [defaut 1] [nomchamp qual_fin_poules_sur]
 [texte qual_fin_poules_sur]}
  clas_fin_poules_sur {[type entier] [taille 32] [taillecar 3] [defaut 1] [nomchamp clas_fin_poules_sur]
 [texte clas_fin_poules_sur]}
  nombre_touches_tableaux {[type entier] [defaut 15] [nomchamp nombre_touches_tableaux] [texte nombre_touches_tableaux]}
  typetype entite
  champs_affiches (nombre_entites clasmt_origine entree_progressive nombre_touches_poules qual_fin_poules_sur
 clas_fin_poules_sur nombre_touches_tableaux)
  tous_les_champs (nombre_entites clasmt_origine entree_progressive nombre_touches_poules qual_fin_poules_sur
 clas_fin_poules_sur nombre_touches_tableaux)
  modifie vrai
)

(def ma_formule
  classe formule
  nombre_entites 64
  clasmt_origine points
  entree_progressive non
  nombre_touches_poules 5
  qual_fin_poules_sur 1
  clas_fin_poules_sur 1
  nombre_touches_tableaux 15
  modifie vrai
  etat poules
  etattour constitution
  nutour 1
)

(def tour_de_poules
  type tour_de_poules
  verif_item verifier_tour_de_poules
  verif_ajout_item verifier_ajout_tour_de_poules
  sans_navigation vrai
  avec_saisie_suppression vrai
  numero {[type entier] [taille 24] [taillecar 3] [saisie_bloquee vrai] [nomchamp numero]
 [texte numero] [no_col 1001] [taille_doc 36] }
  entites_depart {[type entier] [taille 60] [taillecar 3] [nomchamp entites_depart] [texte entites_depart]
 [no_col 1002] [taille_doc 90] }
  entites_exemptees {[type entier] [taille 40] [taillecar 3] [defaut 0] [nomchamp entites_exemptees]
 [texte entites_exemptees] [no_col 1003] [taille_doc 60] }
  entites_dans_poules {[type entier] [taille 60] [taillecar 3] [nomchamp entites_dans_poules] [texte entites_dans_poules]
 [no_col 1004] [taille_doc 90] }
  nombre_poules {[type entier] [taille 60] [taillecar 3] [defaut 0] [nomchamp nombre_poules] [texte
 nombre_poules] [no_col 1005] [taille_doc 90] }
  entites_par_poule {[type entier] [taille 60] [taillecar 3] [defaut 6] [nomchamp entites_par_poule]
 [texte entites_par_poule] [no_col 1006] [taille_doc 90] }
  critere_decal_1 {[type affiliation] [taille 50] [taillecar 10] [nomchamp critere_decal_1] [texte
 critere_decal_1] [no_col 1007] [taille_doc 75] [enumere (("V" club "Verein") ("B"
 ligue "Bundesland") ("N" nation "Nation"))] }
  critere_decal_2 {[type affiliation] [taille 50] [taillecar 10] [nomchamp critere_decal_2] [texte
 critere_decal_2] [no_col 1008] [taille_doc 75] [enumere (("V" club "Verein") ("B"
 ligue "Bundesland") ("N" nation "Nation"))] }
  critere_decal_3 {[type affiliation] [taille 50] [taillecar 10] [nomchamp critere_decal_3] [texte
 critere_decal_3] [no_col 1009] [taille_doc 75] [enumere (("V" club "Verein") ("B"
 ligue "Bundesland") ("N" nation "Nation"))] }
  critere_placement {[type affiliation] [taille 50] [taillecar 10] [nomchamp critere_placement] [texte
 critere_placement] [no_col 1010] [taille_doc 75] [enumere (("V" club "Verein") ("B"
 ligue "Bundesland") ("N" nation "Nation"))] [defaut nation] }
  limite_decal {[type entier] [taille 60] [taillecar 3] [sans_saisie vrai] [invisible vrai] [nomchamp
 limite_decal] [texte limite_decal]}
  qualifies_par_poule {[type entier] [taille 80] [taillecar 3] [defaut 0] [sans_saisie vrai] [invisible
 vrai] [nomchamp qualifies_par_poule] [texte qualifies_par_poule]}
  qualifies_indice {[type entier] [taille 70] [taillecar 3] [defaut 0] [nomchamp qualifies_indice] [texte
 qualifies_indice] [no_col 1011] [taille_doc 105] }
  qualifies {[type entier] [taille 50] [taillecar 3] [defaut 0] [sans_saisie vrai] [invisible
 vrai] [nomchamp qualifies] [texte qualifies]}
  entites_fin {[type entier] [taille 60] [taillecar 3] [defaut 0] [nomchamp entites_fin] [texte
 entites_fin] [no_col 1012] [taille_doc 90] }
  typetype entite
  champs_affiches (numero entites_depart entites_exemptees entites_dans_poules nombre_poules entites_par_poule
 critere_decal_1 critere_decal_2 critere_decal_3 critere_placement qualifies_indice
 entites_fin)
  tous_les_champs (numero entites_depart entites_exemptees entites_dans_poules nombre_poules entites_par_poule
 critere_decal_1 critere_decal_2 critere_decal_3 critere_placement limite_decal qualifies_par_poule
 qualifies_indice qualifies entites_fin)
  genre masculin
  nombre 1
  modifie vrai
  fiche_table ()
)

(def suite_tableaux
  type suite_tableaux
  verif_item verifier_suite_tableaux
  verif_ajout_item verifier_ajout_suite_tableaux
  sans_navigation vrai
  avec_saisie_suppression vrai
  avec_description vrai
  modeles_possibles vrai
  ordre_alpha (nom croissant)
  nom {[type chaine] [texte nom_une_lettre] [taille 24] [taillecar 3] [style majuscule]
 [obligatoire vrai] [nomchamp nom] [no_col 1001] [taille_doc 36] }
  nom_etendu {[type chaine] [taille 100] [taillecar 40] [nomchamp nom_etendu] [texte nom_etendu]
 [no_col 1002] [taille_doc 150] }
  nom_tableaux {[type chaine] [taille 100] [taillecar 40] [nomchamp nom_tableaux] [texte nom_tableaux]
 [no_col 1003] [taille_doc 150] }
  origine1 {[type chaine] [taille 100] [taillecar 40] [nomchamp origine1] [texte origine1] [no_col
 1004] [taille_doc 150] }
  origine2 {[type chaine] [taille 100] [taillecar 40] [nomchamp origine2] [texte origine2] [no_col
 1005] [taille_doc 150] }
  origine3 {[type chaine] [taille 100] [taillecar 40] [nomchamp origine3] [texte origine3] [no_col
 1006] [taille_doc 150] }
  origine4 {[type chaine] [taille 100] [taillecar 40] [nomchamp origine4] [texte origine4] [no_col
 1007] [taille_doc 150] }
  origine5 {[type chaine] [taille 100] [taillecar 40] [nomchamp origine5] [texte origine5] [no_col
 1008] [taille_doc 150] }
  origine6 {[type chaine] [taille 100] [taillecar 40] [nomchamp origine6] [texte origine6] [no_col
 1009] [taille_doc 150] }
  critere_constitution {[type critere_constitution] [taille 100] [taillecar 40] [defaut classement_initial]
 [nomchamp critere_constitution] [texte critere_constitution] [no_col 1010] [taille_doc
 150] [enumere (("n" progression_naturelle "normale Progression") ("u" classement_initial
 "urspruengliche Platzierung") ("r" classement_initial_par_groupes "urspruengliche Platzierung in Gruppen"))]
 }
  hasard_par_2 {[type oui_non] [taille 24] [taillecar 10] [defaut non] [nomchamp hasard_par_2] [texte
 hasard_par_2] [no_col 1011] [taille_doc 36] [enumere (("j" oui "ja") ("n" non "nein"))]
 }
  critere_decal_1 {[type affiliation] [taille 50] [taillecar 10] [nomchamp critere_decal_1] [texte
 critere_decal_1] [no_col 1012] [taille_doc 75] [enumere (("V" club "Verein") ("B"
 ligue "Bundesland") ("N" nation "Nation"))] }
  critere_decal_2 {[type affiliation] [taille 50] [taillecar 10] [nomchamp critere_decal_2] [texte
 critere_decal_2] [no_col 1013] [taille_doc 75] [enumere (("V" club "Verein") ("B"
 ligue "Bundesland") ("N" nation "Nation"))] }
  critere_decal_3 {[type affiliation] [taille 50] [taillecar 10] [nomchamp critere_decal_3] [texte
 critere_decal_3] [no_col 1014] [taille_doc 75] [enumere (("V" club "Verein") ("B"
 ligue "Bundesland") ("N" nation "Nation"))] }
  limite_decal {[type entier] [taille 60] [taillecar 3] [sans_saisie vrai] [invisible vrai] [nomchamp
 limite_decal] [texte limite_decal]}
  protege_debut {[type oui_non] [taille 40] [taillecar 5] [defaut oui] [sans_saisie vrai] [invisible
 vrai] [nomchamp protege_debut] [texte protege_debut]}
  entites_dans_tableau {[type entier] [taille 60] [taillecar 3] [sans_saisie vrai] [invisible vrai] [nomchamp
 entites_dans_tableau] [texte entites_dans_tableau]}
  qualifies {[type entier] [taille 60] [taillecar 3] [defaut 0] [nomchamp qualifies] [texte qualifies]
 [no_col 1015] [taille_doc 90] }
  typetype entite
  champs_affiches (nom nom_etendu nom_tableaux origine1 origine2 origine3 origine4 origine5 origine6
 critere_constitution hasard_par_2 critere_decal_1 critere_decal_2 critere_decal_3
 qualifies)
  tous_les_champs (nom nom_etendu nom_tableaux origine1 origine2 origine3 origine4 origine5 origine6
 critere_constitution hasard_par_2 critere_decal_1 critere_decal_2 critere_decal_3
 limite_decal protege_debut entites_dans_tableau qualifies)
  genre masculin
  nombre 1
  modifie vrai
  fiche_table ()
)

(def description_tableau
  type description_tableau
  verif_item verifier_description_tableau
  verif_ajout_item verifier_ajout_description_tableau
  sans_navigation vrai
  ordre_alpha (suite croissant taille decroissant)
  serie {[type chaine] [style majuscule] [sans_saisie vrai] [invisible vrai] [nomchamp serie]
 [texte serie]}
  nom {[type chaine] [taille 40] [taillecar 5] [style majuscule] [saisie_bloquee vrai]
 [nomchamp nom] [texte nom] [no_col 1001] [taille_doc 60] }
  nom_etendu {[type chaine] [taille 140] [taillecar 40] [nomchamp nom_etendu] [texte nom_etendu]
 [no_col 1002] [taille_doc 210] }
  nombre_entites {[type entier] [taille 50] [taillecar 3] [sans_saisie vrai] [invisible vrai] [nomchamp
 nombre_entites] [texte nombre_entites]}
  classe_apres {[type chaine] [taille 100] [taillecar 20] [sans_saisie vrai] [invisible vrai] [nomchamp
 classe_apres] [texte classe_apres]}
  destination_vainqueurs {[type chaine] [taille 80] [taillecar 5] [nomchamp destination_vainqueurs] [texte
 destination_vainqueurs] [no_col 1003] [taille_doc 120] }
  destination_battus {[type chaine] [taille 80] [taillecar 5] [nomchamp destination_battus] [texte destination_battus]
 [no_col 1004] [taille_doc 120] }
  groupe_clasmt_vainqueur {[type entier] [taille 90] [taillecar 3] [nomchamp groupe_clasmt_vainqueur] [texte
 groupe_clasmt_vainqueur] [no_col 1005] [taille_doc 135] }
  groupe_clasmt_battus {[type entier] [taille 90] [taillecar 3] [nomchamp groupe_clasmt_battus] [texte groupe_clasmt_battus]
 [no_col 1006] [taille_doc 135] }
  rang_premier_vainqueur {[type entier] [taille 80] [taillecar 3] [sans_saisie vrai] [nomchamp rang_premier_vainqueur]
 [texte rang_premier_vainqueur] [no_col 1007] [taille_doc 120]}
  rang_premier_battu {[type entier] [taille 80] [taillecar 3] [sans_saisie vrai] [nomchamp rang_premier_battu]
 [texte rang_premier_battu] [no_col 1008] [taille_doc 120]}
  typetype entite
  champs_affiches (nom nom_etendu destination_vainqueurs destination_battus groupe_clasmt_vainqueur
 groupe_clasmt_battus rang_premier_vainqueur rang_premier_battu)
  tous_les_champs (serie nom nom_etendu nombre_entites classe_apres destination_vainqueurs destination_battus
 groupe_clasmt_vainqueur groupe_clasmt_battus rang_premier_vainqueur rang_premier_battu)
  genre masculin
  modifie vrai
  nombre 10
  fiche_table ()
)

(def criteres_arbitres
  type criteres_arbitres
  critere_decal_1 {[type affiliation] [taille 50] [taillecar 10] [nomchamp critere_decal_1] [texte
 critere_decal_1]}
  critere_decal_2 {[type affiliation] [taille 50] [taillecar 10] [nomchamp critere_decal_2] [texte
 critere_decal_2]}
  critere_decal_3 {[type affiliation] [taille 50] [taillecar 10] [nomchamp critere_decal_3] [texte
 critere_decal_3]}
  typetype entite
  champs_affiches (critere_decal_1 critere_decal_2 critere_decal_3)
  tous_les_champs (critere_decal_1 critere_decal_2 critere_decal_3)
  modifie vrai
)

(def tireur
  type tireur
  affichage (nom prenom serie nation1 ligue1 club1)
  affichage_red (nom prenom nation1 club1)
  de_base vrai
  affiliations (club1 ligue1 nation1)
  ordre_alpha (nom croissant prenom croissant)
  ordre_structure ((nation1 nom) croissant (ligue1 nom) croissant (club1 nom) croissant nom croissant
 prenom croissant)
  ordre_serie (serie croissant nom croissant prenom croissant)
  ordre_rang (rang croissant nom croissant prenom croissant)
  ordre_points (points decroissant nom croissant prenom croissant)
  champs_export (nom prenom date_nais sexe club1 ligue1 nation1 licence licence_fie serie points
 categorie dossard presence rang status paiement mode)
  champs_import (nom prenom date_nais sexe club1 ligue1 nation1 licence licence_fie serie points)
  ordres_doc (ordre_alpha ordre_structure ordre_serie ordre_points)
  selections_doc (tous present absent)
  champs_doc (marge_gche nomprenom nom prenom sexe presence serie club1 ligue1 nation1 date_nais
 licence licence_fie points dossard categorie status paiement mode marge1 marge2
 marge_dte)
  champs_doc_select (marge_gche nomprenom presence serie club1 nation1 date_nais licence marge_dte)
  champs_doc_diapo (marge_gche nomprenom nom prenom sexe presence serie club1 ligue1 nation1 date_nais
 licence licence_fie points dossard categorie status marge1 marge2 marge_dte)
  champs_doc_diapo_select (marge_gche serie marge1 nomprenom nation1)
  champs_doc_web (serie nom prenom nation1 club1)
  ordres_clas (ordre_rang ordre_alpha ordre_structure)
  champs_clas (marge_gche rang nomprenom nom prenom sexe club1 ligue1 nation1 serie points date_nais
 licence licence_fie dossard categorie marge1 marge2 marge_dte)
  champs_clas_select (marge_gche rang nomprenom club1 nation1 marge1 marge2 marge_dte)
  champs_clas_diapo (marge_gche rang nom prenom sexe club1 ligue1 nation1 serie points date_nais licence
 licence_fie dossard categorie marge1 marge2 marge_dte)
  champs_clas_diapo_select (marge_gche rang marge1 nomprenom nation1 marge_dte)
  champs_clas_web (rang nom prenom nation1 club1)
  champs_clas_tab (marge_gche rang nomprenom nom prenom club1 ligue1 nation1 groupe serie points date_nais
 licence licence_fie dossard categorie marge1 marge2 marge_dte)
  champs_clas_tab_select (marge_gche rang nomprenom club1 nation1 marge1 marge_dte)
  champs_clas_tab_diapo (marge_gche rang nomprenom nom prenom club1 ligue1 nation1 groupe serie points date_nais
 licence licence_fie dossard categorie marge1 marge2 marge_dte)
  champs_clas_tab_diapo_select (marge_gche rang nomprenom nation1 marge_dte)
  champs_clas_tab_web (rang nom prenom nation1 club1)
  champs_clas_pou (marge_gche rang nomprenom nom prenom sexe club1 ligue1 nation1 groupe rangpoule
 vic_match indice td serie points date_nais licence licence_fie dossard categorie
 marge1 marge2 marge_dte)
  champs_clas_pou_select (marge_gche rang nomprenom nation1 vic_match indice td groupe marge1 marge_dte)
  champs_clas_pou_web (rang nom prenom nation1 club1 vic_match indice td groupe)
  champs_tir_pou_piste (marge_gche nomprenom nom prenom club1 ligue1 nation1 date_nais licence licence_fie
 dossard categorie poule piste marge1 marge_dte)
  champs_tir_pou_piste_select (marge_gche nomprenom club1 nation1 poule piste marge1 marge_dte)
  champs_pou_const (marge_gche nomprenom nom prenom sexe club1 ligue1 nation1 rang serie dossard marge1
 marge2)
  champs_pou_const_select (marge_gche nomprenom club1 nation1 rang serie marge1 marge2)
  champs_info_pou (marge_gche nomprenom nom prenom sexe club1 ligue1 nation1 grille_feuille marge1
 vic_match indice td rang dossard marge2)
  champs_info_pou_select (marge_gche nomprenom nation1 grille_feuille marge1 vic_match indice td rang marge2)
  champs_info_tab (marge_gche num nomprenom nom prenom sexe club1 ligue1 nation1 dossard)
  champs_info_tab_select (marge_gche num nomprenom nation1)
  champs_info_tab_arbi (marge_gche num titre nomprenom nom prenom sexe club1 ligue1 nation1 dossard)
  champs_info_tab_arbi_select (marge_gche num titre nomprenom nation1)
  champs_feuilles_pou (marge_gche nomprenom nom prenom sexe club1 ligue1 nation1 num grille_feuille signature
 vic_match indice td rang dossard marge1 marge2)
  champs_feuilles_pou_select (marge_gche nomprenom nation1 num grille_feuille signature marge1 marge2)
  champs_feuilles_match (marge_gche nomprenom nom prenom sexe club1 ligue1 nation1 num grille_match signature
 dossard marge1 marge2)
  champs_feuilles_match_select (marge_gche nomprenom nation1 grille_match signature marge1 marge2)
  num {[type entier] [texte numero] [taille 16] [sans_saisie vrai] [invisible vrai] [nomchamp
 num]}
  cle {[type entier] [taille 16] [sans_saisie vrai] [invisible vrai] [nomchamp cle] [texte
 cle]}
  rang {[type entier] [taille 32] [sans_saisie vrai] [invisible vrai] [nomchamp rang] [texte
 rang]}
  rangpoule {[type entier] [taille 32] [sans_saisie vrai] [invisible vrai] [nomchamp rangpoule]
 [texte rangpoule]}
  groupe {[type groupe] [taille 50] [sans_saisie vrai] [invisible vrai] [nomchamp groupe]
 [texte groupe]}
  vic_match {[type chaine] [taille 50] [sans_saisie vrai] [invisible vrai] [nomchamp vic_match]
 [texte vic_match]}
  indice {[type entier] [taille 40] [sans_saisie vrai] [invisible vrai] [nomchamp indice]
 [texte indice]}
  td {[type entier] [taille 40] [sans_saisie vrai] [invisible vrai] [nomchamp td] [texte
 td]}
  signature {[type chaine] [taille 120] [sans_saisie vrai] [invisible vrai] [nomchamp signature]
 [texte signature]}
  poule {[type entier] [sans_saisie vrai] [invisible vrai] [nomchamp poule] [texte poule]}
  piste {[type chaine] [sans_saisie vrai] [invisible vrai] [nomchamp piste] [texte piste]}
  grille_feuille {[type chaine] [texte grille] [taille 200] [sans_saisie vrai] [invisible vrai] [nomchamp
 grille_feuille]}
  grille_match {[type chaine] [texte grille] [taille 200] [sans_saisie vrai] [invisible vrai] [nomchamp
 grille_match]}
  marge_gche {[type chaine] [texte marge] [sans_saisie vrai] [invisible vrai] [nomchamp marge_gche]}
  nom {[type chaine] [taille 100] [taillecar 20] [style majuscule] [obligatoire vrai] [nomchamp
 nom] [texte nom] [no_col 1001] [taille_doc 150] }
  prenom {[type chaine] [taille 80] [taillecar 20] [style titre] [nomchamp prenom] [texte
 prenom] [no_col 1002] [taille_doc 120] }
  presence {[type presence] [taille 24] [taillecar 10] [defaut present] [nomchamp presence]
 [texte presence] [no_col 1003] [taille_doc 36] [enumere (("a" present "anwesend")
 ("n" absent "nicht anwesend"))] }
  serie {[type entier] [taille 32] [taillecar 3] [nomchamp serie] [texte serie] [no_col 1004]
 [taille_doc 48] }
  club1 {[type club] [texte club] [taille 100] [taillecar 20] [mere ligue1] [si_modifie (verif_club
 maj_affiliation)] [nomchamp club1] [no_col 1005] [taille_doc 150] }
  ligue1 {[type ligue] [texte ligue] [taille 100] [taillecar 20] [fille club1] [mere nation1]
 [si_modifie (verif_club maj_affiliation)] [nomchamp ligue1] [no_col 1006] [taille_doc
 150] }
  nation1 {[type nation] [texte nation] [taille 42] [taillecar 5] [fille ligue1] [maxcar 3]
 [si_modifie (verif_club maj_affiliation)] [nomchamp nation1] [no_col 1007] [taille_doc
 63] }
  date_nais {[type date] [taille 70] [taillecar 12] [maxcar 12] [nomchamp date_nais] [texte date_nais]
 [no_col 1008] [taille_doc 105] }
  licence {[type chaine] [taille 80] [taillecar 12] [maxcar 12] [si_modifie (verif_club)] [nomchamp
 licence] [texte licence] [no_col 1009] [taille_doc 120] }
  licence_fie {[type chaine] [taille 80] [taillecar 16] [nomchamp licence_fie] [texte licence_fie]
 [no_col 1010] [taille_doc 120] }
  points {[type decimal] [taille 40] [taillecar 6] [nomchamp points] [texte points] [no_col
 1011] [taille_doc 60] }
  dossard {[type entier] [taille 32] [taillecar 4] [nomchamp dossard] [texte dossard] [no_col
 1012] [taille_doc 48] }
  categorie {[type categorie] [taille 24] [taillecar 10] [sans_saisie vrai] [si_besoin calcul_categorie]
 [nomchamp categorie] [texte categorie] [no_col 1013] [taille_doc 36]}
  sexe {[type sexe] [taille 24] [taillecar 10] [nomchamp sexe] [texte sexe] [no_col 1014]
 [taille_doc 36] [enumere (("H" masculin "Herren") ("D" feminin "Damen"))] [defaut
 masculin] }
  status {[type status] [taille 45] [taillecar 10] [defaut normal] [nomchamp status] [texte
 status] [no_col 1015] [taille_doc 67] [enumere (("n" normal "normal") ("A" abandon
 "Aufgabe") ("S" exclusion "Schwarze Karte") ("N" forfait "Nichtantreten"))] }
  paiement {[type decimal] [texte paiement] [taille 40] [taillecar 6] [nomchamp paiement] [no_col
 1016] [taille_doc 60] }
  mode {[type chaine] [texte mode] [taille 50] [taillecar 10] [nomchamp mode] [no_col 1017]
 [taille_doc 75] }
  nomprenom {[type chaine] [taille 120] [sans_saisie vrai] [invisible vrai] [si_besoin calcul_nomprenom]
 [nomchamp nomprenom] [texte nomprenom]}
  marge1 {[type chaine] [texte marge] [sans_saisie vrai] [invisible vrai] [nomchamp marge1]}
  marge2 {[type chaine] [texte marge] [sans_saisie vrai] [invisible vrai] [nomchamp marge2]}
  marge_dte {[type chaine] [texte marge] [sans_saisie vrai] [invisible vrai] [nomchamp marge_dte]}
  typetype entite
  champs_affiches (nom prenom presence serie club1 ligue1 nation1 date_nais licence licence_fie points
 dossard categorie sexe status paiement mode)
  tous_les_champs (num cle rang rangpoule groupe vic_match indice td signature poule piste grille_feuille
 grille_match marge_gche nom prenom presence serie club1 ligue1 nation1 date_nais
 licence licence_fie points dossard categorie sexe status paiement mode nomprenom
 marge1 marge2 marge_dte)
  genre masculin
  nombre 42
  fiche_table ()
)

(def arbitre
  type arbitre
  affichage (nom prenom nation1 ligue1 club1)
  affichage_red (nom prenom nation1 club1)
  de_base vrai
  affiliations (club1 ligue1 nation1 club2 ligue2 nation2)
  ordre_alpha (nom croissant prenom croissant)
  ordre_structure ((nation1 nom) croissant (ligue1 nom) croissant (club1 nom) croissant nom croissant
 prenom croissant)
  ordres_doc (ordre_alpha ordre_structure)
  selections_doc (tous present absent)
  champs_doc (marge_gche nomprenom nom prenom presence sexe categorie club1 ligue1 nation1 club2
 ligue2 nation2 marge1 marge2 marge_dte)
  champs_doc_select (marge_gche nomprenom categorie nation1 marge_dte)
  champs_doc_diapo (marge_gche nomprenom nom prenom presence sexe categorie club1 ligue1 nation1 club2
 ligue2 nation2 marge1 marge2 marge_dte)
  champs_doc_diapo_select (marge_gche nomprenom nation1)
  champs_doc_web (nom prenom nation1 categorie poules matches finales)
  champs_activ (marge_gche nomprenom nom prenom presence sexe categorie club1 ligue1 nation1 club2
 ligue2 nation2 marge1 marge2 poules matches finales marge_dte)
  champs_activ_select (marge_gche nomprenom categorie nation1 poules matches finales marge_dte)
  champs_export (nom prenom sexe categorie club1 ligue1 nation1)
  champs_import (nom prenom sexe categorie club1 ligue1 nation1)
  marge_gche {[type chaine] [texte marge] [sans_saisie vrai] [invisible vrai] [nomchamp marge_gche]}
  titre {[type chaine] [sans_saisie vrai] [invisible vrai] [defaut &arbitre] [nomchamp titre]
 [texte titre]}
  nom {[type chaine] [taille 100] [taillecar 16] [style majuscule] [obligatoire vrai] [nomchamp
 nom] [texte nom] [no_col 1001] [taille_doc 150] }
  prenom {[type chaine] [taille 80] [taillecar 12] [style titre] [nomchamp prenom] [texte
 prenom] [no_col 1002] [taille_doc 120] }
  presence {[type presence] [taille 24] [taillecar 10] [defaut present] [nomchamp presence]
 [texte presence] [no_col 1003] [taille_doc 36] [enumere (("a" present "anwesend")
 ("n" absent "nicht anwesend"))] }
  sexe {[type sexe] [taille 24] [taillecar 10] [nomchamp sexe] [texte sexe] [no_col 1004]
 [taille_doc 36] [enumere (("H" masculin "Herren") ("D" feminin "Damen"))] }
  categorie {[type chaine] [taille 24] [taillecar 10] [style majuscule] [nomchamp categorie]
 [texte categorie] [no_col 1005] [taille_doc 36] }
  date_nais {[type date] [taille 70] [taillecar 12] [maxcar 12] [nomchamp date_nais] [texte date_nais]
 [no_col 1006] [taille_doc 105] }
  licence_fie {[type chaine] [taille 80] [taillecar 16] [nomchamp licence_fie] [texte licence_fie]
 [no_col 1007] [taille_doc 120] }
  club1 {[type club] [texte club] [taille 100] [taillecar 12] [mere ligue1] [nomchamp club1]
 [no_col 1008] [taille_doc 150] }
  ligue1 {[type ligue] [texte ligue] [taille 100] [taillecar 12] [fille club1] [mere nation1]
 [nomchamp ligue1] [no_col 1009] [taille_doc 150] }
  nation1 {[type nation] [texte nation] [taille 42] [taillecar 5] [fille ligue1] [maxcar 3]
 [nomchamp nation1] [no_col 1010] [taille_doc 63] }
  nomprenom {[type chaine] [taille 120] [sans_saisie vrai] [invisible vrai] [si_besoin calcul_nomprenom]
 [nomchamp nomprenom] [texte nomprenom]}
  poules {[type chaine] [texte nb_poules] [sans_saisie vrai] [invisible vrai] [nomchamp poules]}
  matches {[type chaine] [texte nb_matches] [sans_saisie vrai] [invisible vrai] [nomchamp matches]}
  finales {[type chaine] [texte nb_finales] [sans_saisie vrai] [invisible vrai] [nomchamp finales]}
  marge1 {[type chaine] [texte marge] [sans_saisie vrai] [invisible vrai] [nomchamp marge1]}
  marge2 {[type chaine] [texte marge] [sans_saisie vrai] [invisible vrai] [nomchamp marge2]}
  marge_dte {[type chaine] [texte marge] [sans_saisie vrai] [invisible vrai] [nomchamp marge_dte]}
  typetype entite
  champs_affiches (nom prenom presence sexe categorie date_nais licence_fie club1 ligue1 nation1)
  tous_les_champs (marge_gche titre nom prenom presence sexe categorie date_nais licence_fie club1
 ligue1 nation1 nomprenom poules matches finales marge1 marge2 marge_dte)
  genre masculin
  nombre 0
  fiche_table ()
)

(def club
  type club
  affichage (nom ligue1 nation1)
  affiliations (ligue1 nation1)
  ordre_alpha (nom croissant)
  ordre_structure ((nation1 nom) croissant (ligue1 nom) croissant nom croissant)
  utilise_par (tireur arbitre)
  ordres_doc (ordre_alpha ordre_structure)
  selections_doc (tous concerne_tir concerne_tir_pres concerne_arbi concerne_arbi_pres)
  champs_doc (marge_gche nom ligue1 nation1 effectif marge1 marge_dte)
  champs_doc_diapo (marge_gche nom ligue1 nation1 effectif marge1 marge_dte)
  champs_doc_web (nom nation1 effectif)
  champs_export (nom ligue1 nation1)
  marge_gche {[type chaine] [texte marge] [sans_saisie vrai] [invisible vrai] [nomchamp marge_gche]}
  nom {[type chaine] [taille 100] [taillecar 12] [style majuscule] [obligatoire vrai] [nomchamp
 nom] [texte nom] [no_col 1001] [taille_doc 150] }
  ligue1 {[type ligue] [taille 100] [taillecar 12] [texte ligue] [mere nation1] [nomchamp
 ligue1] [no_col 1002] [taille_doc 150] }
  nation1 {[type nation] [taille 42] [taillecar 5] [texte nation] [maxcar 3] [fille ligue1]
 [nomchamp nation1] [no_col 1003] [taille_doc 63] }
  effectif {[type entier] [sans_saisie vrai] [invisible vrai] [nomchamp effectif] [texte effectif]}
  marge1 {[type chaine] [texte marge] [sans_saisie vrai] [invisible vrai] [nomchamp marge1]}
  marge_dte {[type chaine] [texte marge] [sans_saisie vrai] [invisible vrai] [nomchamp marge_dte]}
  typetype entite
  champs_affiches (nom ligue1 nation1)
  tous_les_champs (marge_gche nom ligue1 nation1 effectif marge1 marge_dte)
  genre masculin
  nombre 42
  fiche_table ()
)

(def ligue
  type ligue
  affichage (nom nation1)
  utilise_par (club tireur arbitre)
  affiliations (nation1)
  ordre_alpha (nom croissant)
  ordre_structure ((nation1 nom) croissant nom croissant)
  ordres_doc (ordre_alpha ordre_structure)
  selections_doc (tous concerne_tir concerne_tir_pres concerne_arbi concerne_arbi_pres)
  champs_doc (marge_gche nom nation1 effectif marge1 marge_dte)
  champs_doc_diapo (marge_gche nom nation1 effectif marge1 marge_dte)
  champs_doc_web (nom nation1 effectif)
  champs_export (nom nation1)
  marge_gche {[type chaine] [texte marge] [sans_saisie vrai] [invisible vrai] [nomchamp marge_gche]}
  nom {[type chaine] [taille 100] [taillecar 12] [style majuscule] [obligatoire vrai] [nomchamp
 nom] [texte nom] [no_col 1001] [taille_doc 150] }
  nation1 {[type nation] [taille 42] [taillecar 5] [texte nation] [maxcar 3] [nomchamp nation1]
 [no_col 1002] [taille_doc 63] }
  effectif {[type entier] [sans_saisie vrai] [invisible vrai] [nomchamp effectif] [texte effectif]}
  marge1 {[type chaine] [texte marge] [sans_saisie vrai] [invisible vrai] [nomchamp marge1]}
  marge_dte {[type chaine] [texte marge] [sans_saisie vrai] [invisible vrai] [nomchamp marge_dte]}
  typetype entite
  champs_affiches (nom nation1)
  tous_les_champs (marge_gche nom nation1 effectif marge1 marge_dte)
  genre masculin
  nombre 0
  fiche_table ()
)

(def nation
  type nation
  affichage (nom)
  ordre_alpha (nom croissant)
  ordre_structure (nom croissant)
  utilise_par (ligue club tireur arbitre)
  ordres_doc (ordre_alpha ordre_structure)
  selections_doc (tous concerne_tir concerne_tir_pres concerne_arbi concerne_arbi_pres)
  champs_doc (marge_gche nom nom_etendu effectif marge1 marge_dte)
  champs_doc_diapo (marge_gche nom nom_etendu effectif marge1 marge_dte)
  champs_doc_web (nom nom_etendu effectif)
  champs_export (nom nom_etendu)
  marge_gche {[type chaine] [texte marge] [sans_saisie vrai] [invisible vrai] [nomchamp marge_gche]}
  nom {[type chaine] [taille 42] [taillecar 5] [maxcar 3] [style majuscule] [obligatoire
 vrai] [nomchamp nom] [texte nom] [no_col 1001] [taille_doc 63] }
  nom_etendu {[type chaine] [taille 100] [taillecar 16] [style titre] [nomchamp nom_etendu] [texte
 nom_etendu] [no_col 1002] [taille_doc 150] }
  effectif {[type entier] [sans_saisie vrai] [invisible vrai] [nomchamp effectif] [texte effectif]}
  marge_dte {[type chaine] [texte marge] [sans_saisie vrai] [invisible vrai] [nomchamp marge_dte]}
  typetype entite
  champs_affiches (nom nom_etendu)
  tous_les_champs (marge_gche nom nom_etendu effectif marge_dte)
  genre masculin
  modifie vrai
  nombre 0
  fiche_table ()
)