package models

type Configuration struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Filter      string `json:"filter"`
	Value       string `json:"value"`
	FieldType   string `json:"field_type"`
}
