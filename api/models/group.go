package models

type Field struct {
	Name            string        `json:"name"`
	Description     string        `json:"description"`
	Slug            string        `json:"slug"`
	OrderNumber     int32         `json:"order_number"`
	Status          int32         `json:"status"`
	IsRequired      bool          `json:"is_required"`
	Placeholder     string        `json:"placeholder"`
	FieldType       string        `json:"field_type"`
	SelectType      []*SelectType `json:"select_type"`
	ValidationRegex string        `json:"validation_regex"`
	ValidationFunc  string        `json:"validation_func"`
	GroupId         string        `json:"group_id"`
	Min             int32         `json:"min"`
	Max             int32         `json:"max"`
	DefaultValue    string        `json:"default_value"`
	IsSearchable    int32         `json:"is_searchable"`
}

type SelectType struct {
	XId         string `json:"_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Value       string `json:"value"`
}

type Group struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Slug        string  `json:"slug"`
	Status      int32   `json:"status"`
	OrderNumber int32   `json:"order_number"`
	ParentId    *string `json:"parent_id"`
	GroupType   int32   `json:"group_type"`
}

type GetAllGroupResponse struct {
	Groups []*Group `json:"groups"`
	Count  int32    `json:"count"`
}

type Entity struct {
	XId         string `json:"_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Slug        string `json:"slug"`
	Status      int32  `json:"status"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type GetAllEntityResponse struct {
	Entities []map[string]interface{} `json:"entities"`
	Count    int32                    `json:"count"`
}

type DynamicGroup struct {
	XId         string          `json:"_id"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Slug        string          `json:"slug"`
	Status      int32           `json:"status"`
	CreatedAt   string          `json:"created_at"`
	OrderNumber int32           `json:"order_number"`
	ParentId    *string         `json:"parent_id"`
	Children    []*DynamicGroup `json:"children"`
	Fields      []*DynamicField `json:"fields"`
	UpdatedAt   string          `json:"updated_at"`
	GroupType   int32           `json:"group_type"`
}

type GetAllDynamicGroupResponse struct {
	DynamicGroups []*DynamicGroup `json:"groups"`
	Count         int32           `json:"count"`
}

type DynamicField struct {
	XId             string        `json:"_id"`
	Name            string        `json:"name"`
	Description     string        `json:"description"`
	Slug            string        `json:"slug"`
	OrderNumber     int32         `json:"order_number"`
	Status          int32         `json:"status"`
	CreatedAt       string        `json:"created_at"`
	IsRequired      bool          `json:"is_required"`
	Placeholder     string        `json:"placeholder"`
	FieldType       string        `json:"field_type"`
	SelectType      []*SelectType `json:"select_type"`
	ValidationRegex string        `json:"validation_regex"`
	ValidationFunc  string        `json:"validation_func"`
	GroupId         string        `json:"group_id"`
	UpdatedAt       string        `json:"update_at"`
	Min             int32         `json:"min"`
	Max             int32         `json:"max"`
	DefaultValue    string        `json:"default_value"`
	IsSearchable    int32         `json:"is_searchable"`
}

type Lookup struct {
	From         string `json:"from" required:"true"`
	LocalField   string `json:"localField" required:"true"`
	ForeignField string `json:"foreignField" required:"true"`
	As           string `json:"as" required:"true"`
}

type JoinGroupRequest struct {
	XId       string     `json:"_id"`
	Name      string     `json:"name"`
	FromDate  string     `json:"from_date"`
	ToDate    string     `json:"to_date"`
	Aggregate *Aggregate `json:"aggregate"`
}

type Aggregate struct {
	Lookups []*Lookup         `json:"lookups"`
	Group   *AggregationGroup `json:"group"`
}

type AggregationGroup struct {
	XId         string `json:"_id"`
	Field       string `json:"field"`
	Accumulator string `json:"accumulator"`
	Expression  string `json:"expression"`
}

type EntityIdRequest struct {
	Id string `json:"_id"`
}
