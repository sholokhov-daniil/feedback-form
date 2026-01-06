package dto

type Form struct {
    ID   string `json:"id"`
    Name string `json:"name"`
	Fields []FormField `json:"fields"`
}

type FormField struct {
	ID string `json:"id"`
	Code string `json:code`
	Name string `json:"name"`
	Type string `json:"type"`
}