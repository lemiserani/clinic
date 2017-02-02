package main

import "github.com/lemiserani/clinic/model"

var clinics model.Clinics

//Give us some seed data
func init() {
	RepoCreateClinic(model.Clinic{ID: 1, Name: "Dentistas de Três Pontas", Local: "Rua Marques de Abrantes, 342, Centro, CEP 37190-000, Três Pontas, MG", Phone: "(35) 3265-2928"})
	RepoCreateClinic(model.Clinic{ID: 2, Name: "Dentistas da Lapa", Local: "Afonso Sardinha, número 199, Lapa, São Paulo, SP", Phone: "(11) 3644-7777"})
}

//RepoCreateClinic create new Clinic
func RepoCreateClinic(cli model.Clinic) model.Clinic {
	clinics = append(clinics, cli)
	return cli
}

//RepoFindClinic find Clinic
func RepoFindClinic(id int) model.Clinic {
	for _, cli := range clinics {
		if cli.ID == id {
			return cli
		}
	}
	return model.Clinic{}
}
