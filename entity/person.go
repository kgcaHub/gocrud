package entity

//Entidad persona
type Person struct {
	Id   int64  `json:"id,string"`
	Name string `json:"name"`
}
