package models

type Thematic struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Description string `json:"description"`
}

type FilterThematic struct {
	Id    []int `json:"id"`
}