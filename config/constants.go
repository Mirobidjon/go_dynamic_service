package config

import (
	"time"

	"github.com/mirobidjon/go_dynamic_service/api/models"
)

const (
	DateLayout              = "2006-01-02"
	TimeLayout              = "2006-01-02T15:04:05Z07:00"
	DatabaseQueryTimeLayout = `'YYYY-MM-DD"T"HH24:MI:SS"."MS"Z"TZ'`
	// DatabaseTimeLayout
	DatabaseTimeLayout string = time.RFC3339
	// AccessTokenExpiresInTime ...
	AccessTokenExpiresInTime time.Duration = 1 * 24 * 60 * time.Minute
	// RefreshTokenExpiresInTime ...
	RefreshTokenExpiresInTime time.Duration = 30 * 24 * 60 * time.Minute

	TimeStampLayout = "2006-01-02 15:04:05"

	ConsumerGroupID string = "mongo_document_service"
	ProtoType       string = "proto"
	StructType      string = "struct"
)

var (
	FieldTypeConfigurations = []models.Configuration{
		{
			Name:        "text",
			Description: "Textoviy maydon, (regex uchun uzunligini tekshirish `.{min,max}`)",
			Value:       "text",
			FieldType:   "string",
		},
		{
			Name:        "number",
			Description: "Raqamli maydon Min Max bo'yicha tekshiriladi",
			Value:       "number",
			FieldType:   "string",
		},
		{
			Name:        "float",
			Description: "Float maydon Min Max bo'yicha tekshiriladi",
			Value:       "float",
			FieldType:   "string",
		},
		{
			Name:        "file",
			Description: "Fayl maydon, field uchun link yuboriladi",
			Value:       "file",
			FieldType:   "string",
		},
		{
			Name:        "bool",
			Description: "Boolean maydon, true yoki false qiymatlarini qabul qiladi, ",
			Value:       "bool",
			FieldType:   "string",
		},
		{
			Name:        "date",
			Description: "Sana maydon, `2006-01-02` formatida qabul qiladi (validatsiya uchun `date`)",
			Value:       "date",
			FieldType:   "string",
		},
		{
			Name:        "datetime",
			Description: "Sana va vaqt maydon, `2006-01-02 15:04:05` formatida qabul qiladi (validatsiya uchun `time`)",
			Value:       "datetime",
			FieldType:   "string",
		},
		{
			Name:        "uuid",
			Description: "UUID maydon, `uuid` formatida qabul qiladi (validatsiya uchun `uuid`)",
			Value:       "uuid",
			FieldType:   "string",
		},
		{
			Name:        "object_id",
			Description: "Object ID maydon, `object_id` formatida qabul qiladi (validatsiya uchun `objectId`)",
			Value:       "object_id",
			FieldType:   "string",
		},
		{
			Name:        "Point",
			Description: "Geo Point maydon, `point` formatida qabul qiladi (validatsiya uchun `point` '{\"type\": \"Point\", \"coordinates\": [100.0, 0.0]}' ) https://www.mongodb.com/docs/v5.0/geospatial-queries/",
			Value:       "point",
			FieldType:   "string",
		},
		{
			Name:        "Polygon",
			Description: "Geo Polygon maydon, `polygon` formatida qabul qiladi (validatsiya uchun `polygon` '{\"type\": \"Polygon\", \"coordinates\": [[[100.0, 0.0], [101.0, 0.0], [101.0, 1.0], [100.0, 1.0], [100.0, 0.0]]]}' ) https://www.mongodb.com/docs/v5.0/geospatial-queries/",
			Value:       "polygon",
			FieldType:   "string",
		},
	}

	GroupTypeConfigurations = []models.Configuration{
		{
			Name:        "Object",
			Description: "Object",
			Value:       "1",
			FieldType:   "number",
		},
		{
			Name:        "Array",
			Description: "Array",
			Value:       "2",
			FieldType:   "number",
		},
	}

	DefaultValuesConfiguration = []models.Configuration{
		{
			Name:        "Generate UUID",
			Value:       "GENERATE",
			FieldType:   "string",
			Description: "Generate UUID or ObjectID if empty or null value",
		},
		{
			Name:        "TIME NOW",
			Value:       "TIME_NOW",
			FieldType:   "string",
			Description: "hozirgi vaqtni default qiymat sifatida qabul qiladi (sana va vaqt uchun)",
		},
		{
			Name:        "Null",
			Value:       "null",
			FieldType:   "string",
			Description: "Null value",
		},
	}

	ValidationFunctionConfiguration = []models.Configuration{
		{
			Name:        "Phone",
			Description: "Telefon raqamni tekshirish uchun funksiya",
			Value:       "phone",
			FieldType:   "string",
		},
		{
			Name:        "Email",
			Description: "Emailni tekshirish uchun funksiya",
			Value:       "email",
			FieldType:   "string",
		},
		{
			Name:        "Uuid",
			Description: "Uuidni tekshirish uchun funksiya",
			Value:       "uuid",
			FieldType:   "string",
		},
		{
			Name:        "ObjectId",
			Description: "ObjectIdni tekshirish uchun funksiya",
			Value:       "objectId",
			FieldType:   "string",
		},
		{
			Name:        "Date",
			Description: "Sana formatini tekshirish uchun funksiya",
			Value:       "date",
			FieldType:   "string",
		},
		{
			Name:        "Time",
			Description: "Vaqt formatini tekshirish uchun funksiya",
			Value:       "time",
			FieldType:   "string",
		},
	}
)
