package models

type ProjectMin struct{
	Id 	    int 	`json:"id"`
	Name    string 	`json:"name"`
}

type Project struct{
	Id 	    int 	`json:"id"`
	Name    string 	`json:"name"`
    Purpose string  `json:"purpose"`
	Relevance string `json:"relevance"`
	Result string `json:"result"`
	Specializations []*Specialization `json:"specializations"`
	Thematics []*Thematic `json:"thematics"`
}