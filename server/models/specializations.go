package models

type Specialization struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Description string `json:"description"`
}

type FilterSpecialization struct {
	Id    []int `json:"id"`
}