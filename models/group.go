package models

const (
	FieldTypeNumber   = "number"
	FieldTypeFloat    = "float"
	FieldTypeText     = "text"
	FieldTypeRadio    = "radio"
	FieldTypeSelect   = "select"
	FieldTypeFile     = "file"
	FieldTypeBool     = "bool"
	FieldTypeDate     = "date"
	FieldTypeDateTime = "datetime"
	FieldTypeUuid     = "uuid"
	FieldTypeObjectID = "object_id"
	FieldTypeJsonb    = "jsonb"
	FieldTypePolygon  = "polygon"
)

type SelectTypes struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Value       string `json:"value"`
}

const (
	GroupTypeObject int32 = 1
	GroupTypeArray  int32 = 2
)

type Field struct {
	ID              string      `json:"id"`
	Name            string      `json:"name"`
	Description     string      `json:"description"`
	Slug            string      `json:"slug"`
	OrderNumber     int         `json:"order_number"`
	Status          int         `json:"status"`
	IsRequired      bool        `json:"is_required"`
	IsSearchable    int32       `json:"is_searchable"`
	Placeholder     string      `json:"placeholder"`
	FieldType       string      `json:"field_type"`
	SelectType      SelectTypes `json:"select_type"`
	ValidationRegex string      `json:"validation_regex"`
	ValidationFunc  string      `json:"validation_func"`
	GroupID         string      `json:"group_id"`
	Min             int32       `json:"min"`
	Max             int32       `json:"max"`
	DefaultValue    string      `json:"default_value"`
	CreatedAt       string      `json:"created_at"`
	UpdatedAt       string      `json:"updated_at"`
}

type Group struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Slug        string   `json:"slug"`
	Status      int      `json:"status"`
	OrderNumber int      `json:"order_number"`
	ParentID    *string  `json:"parent_id"`
	GroupType   int32    `json:"group_type"`
	CreatedAt   string   `json:"created_at"`
	UpdatedAt   string   `json:"updated_at"`
	Fields      []*Field `json:"fields"`
	Children    []*Group `json:"children"`
}
