package models

const (
	QueryTypeExec  = "exec"
	QueryTypeQuery = "query"
)

type Query struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Slug        string `json:"slug"`
	GroupSlug   string `json:"group_slug"`
	Query       string `json:"query"`
	QueryType   string `json:"query_type"`
	Status      int32  `json:"status"`
}

type QueryField struct {
	ID           string `json:"id"`
	QueryID      string `json:"query_id"`
	FieldID      string `json:"field_id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Slug         string `json:"slug"`
	OrderNumber  int32  `json:"order_number"`
	Status       int32  `json:"status"`
	IsRequired   bool   `json:"is_required"`
	IsSearchable int32  `json:"is_searchable"`
	FieldType    string `json:"field_type"`
}
