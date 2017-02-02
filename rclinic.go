package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

//Index main route
func Index(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprint(w, "Clinicas!\n")
}

//AllClinics return all clinics
func AllClinics(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(clinics); err != nil {
		panic(err)
	}
}

//FindByIDClinic return specific clinic
func FindByIDClinic(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	vars := p.ByName("id")
	var id int
	var err error
	if id, err = strconv.Atoi(vars); err != nil {
		fmt.Fprintln(w, id)
		panic(err)
	}
	clinic := RepoFindClinic(id)
	if clinic.ID > 0 {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(clinic); err != nil {
			panic(err)
		}
		return
	}

	// If we didn't find it, 404
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found clinic"}); err != nil {
		panic(err)
	}
}

//KindsReceipt return specific receipt
func KindsReceipt(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	args := []string{"Dinheiro", "Cartão Débito", "Cartão de Crédito", "Cheque"}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(args); err != nil {
		panic(err)
	}
}

//KindsCosts return specific cost
func KindsCosts(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	args := []string{"Salário Thiago", "Salário Sostenes", "Salário Luana", "Salário Recepcionista",
		"Água", "Luz", "Telefone", "Materiais Odontológico - (Dental)", "Protético", "Impostos", "Publicidade", "Equipamentos", "Coleta Lixo", "Contador", "Custo Máquina de crédito", "Outros Gastos"}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(args); err != nil {
		panic(err)
	}
}
