package model

//Clinic is objeto about clinic
type Clinic struct {
	ID    int    `json:"Id"`
	Name  string `json:"Name"`
	Local string `json:"Local"`
	Phone string `json:"Phone"`
}

//Clinics is Array of Clinic
type Clinics []Clinic
