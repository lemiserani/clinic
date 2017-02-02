package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"

	"clinic/model"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type (
	// ClosureController create Session to mongo
	ClosureController struct {
		session *mgo.Session
	}
)

// NewClosureController provides a reference to a ClosureController with provided mongo session
func NewClosureController(s *mgo.Session) *ClosureController {
	return &ClosureController{s}
}

// GetClosure retrieves an individual closure resource
func (cc ClosureController) GetClosure(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	if !bson.IsObjectIdHex(id) {
		codeHTTP(w, 206, "Id do Documento Inválido!"+id)
		return
	}
	oid := bson.ObjectIdHex(id)
	u := model.Closure{}

	if err := cc.session.DB("api_clinic").C("closures").FindId(oid).One(&u); err != nil {
		codeHTTP(w, 404, "Não encontrado registro!"+id)
	} else {
		uj, _ := json.Marshal(u)
		w.Header().Set("Content-Type", "application/json")
		codeHTTP(w, 200, string(uj))
	}
}

// CreateClosure creates a new closure resource
func (cc ClosureController) CreateClosure(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	closure := model.Closure{}

	json.NewDecoder(r.Body).Decode(&closure)
	closure.ID = bson.NewObjectId()
	closure.Created = time.Now()

	cc.session.DB("api_clinic").C("closures").Insert(closure)
	uj, _ := json.Marshal(closure)

	w.Header().Set("Content-Type", "application/json")
	codeHTTP(w, 201, "Criado com sucesso! "+string(uj))
}

// UpdateClosure creates a new closure resource
func (cc ClosureController) UpdateClosure(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	closure := model.Closure{}
	id := p.ByName("id")
	json.NewDecoder(r.Body).Decode(&closure)

	cc.session.DB("api_clinic").C("closures").UpdateId(bson.ObjectIdHex(id), bson.M{"$set": bson.M{"name": closure.Name,
		"value": closure.Value}})
	codeHTTP(w, 201, "Atualizado com sucesso! "+id)
}

// RemoveClosure removes an existing closure resource
func (cc ClosureController) RemoveClosure(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	if !bson.IsObjectIdHex(id) {
		codeHTTP(w, 206, "Parâmetro Inválido! "+id)
		return
	}
	oid := bson.ObjectIdHex(id)

	if err := cc.session.DB("api_clinic").C("closures").RemoveId(oid); err != nil {
		codeHTTP(w, 404, "Não encontrado registro! "+id)
	} else {
		codeHTTP(w, 200, "Deletado com sucesso! "+id)
	}
}

// SearchClosures retrieves closure resource
func (cc ClosureController) SearchClosures(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	month, err := strconv.Atoi(r.URL.Query().Get("month"))
	year, err := strconv.Atoi(r.URL.Query().Get("year"))
	clinic := r.URL.Query().Get("clinic")
	if err != nil {
		print(error.Error)
		i := "Parâmetros Inválidos! month: " + r.URL.Query().Get("month") + " year: " + r.URL.Query().Get("year") + " clinic: " + clinic
		codeHTTP(w, 206, i)
		// return
	}
	u := model.Closures{}
	if err := cc.session.DB("api_clinic").C("closures").Find(bson.M{"month": month, "year": year, "clinic": clinic}).All(&u); err != nil {
		i := "Não encontrado registros! month: " + r.URL.Query().Get("month") + " year: " + r.URL.Query().Get("year") + " clinic: " + clinic
		codeHTTP(w, 404, i)
	} else {
		uj, _ := json.Marshal(u)
		codeHTTP(w, 200, string(uj))
	}
}

func codeHTTP(w http.ResponseWriter, code int, i interface{}) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%+v", i)
}
