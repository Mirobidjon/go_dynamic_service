package handlers

import (
	"kassa360/kassa360_go_dynamic_service/api/http"
	"kassa360/kassa360_go_dynamic_service/config"
	"regexp"

	fiber "github.com/gofiber/fiber/v2"
)

// FieldTypeConfiguration godoc
// @Security ApiKeyAuth
// @Security ApiKeyAuth
// @ID field_type_configuration
// @Router /client-api/configuration/field_types [GET]
// @Summary Field Type Configuration
// @Description Field Type Configuration
// @Accept json
// @Produce json
// @Tags Configuration
// @Success 200 {object} http.Response{data=[]models.Configuration} "Success"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) FieldTypeConfiguration(c *fiber.Ctx) error {
	return h.handleResponse(c, http.OK, config.FieldTypeConfigurations, "get_field_type_configuration", "")
}

// GroupTypeConfiguration godoc
// @Security ApiKeyAuth
// @ID group_type_configuration
// @Router /client-api/configuration/group_types [GET]
// @Summary Group Type Configuration
// @Description Group Type Configuration
// @Accept json
// @Produce json
// @Tags Configuration
// @Success 200 {object} http.Response{data=[]models.Configuration}  "Success"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GroupTypeConfiguration(c *fiber.Ctx) error {
	return h.handleResponse(c, http.OK, config.GroupTypeConfigurations, "get_group_type_configuration", "")
}

// DefaultValuesConfiguration godoc
// @Security ApiKeyAuth
// @ID default_values_configuration
// @Router /client-api/configuration/default_values [GET]
// @Summary Default Values Configuration
// @Description Default Values Configuration
// @Accept json
// @Produce json
// @Tags Configuration
// @Success 200 {object} http.Response{data=[]models.Configuration}  "Success"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DefaultValuesConfiguration(c *fiber.Ctx) error {
	return h.handleResponse(c, http.OK, config.DefaultValuesConfiguration, "get_default_values_configuration", "")
}

// ValidationFunctionConfiguration godoc
// @Security ApiKeyAuth
// @ID validation_function_configuration
// @Router /client-api/configuration/validation_functions [GET]
// @Summary Validation Function Configuration
// @Description Validation Function Configuration
// @Accept json
// @Produce json
// @Tags Configuration
// @Success 200 {object} http.Response{data=[]models.Configuration}  "Success"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) ValidationFunctionConfiguration(c *fiber.Ctx) error {
	return h.handleResponse(c, http.OK, config.ValidationFunctionConfiguration, "get_validation_function_configuration", "")
}

// RegexConfiguration godoc
// @Security ApiKeyAuth
// @ID regex_configuration
// @Router /client-api/configuration/regex [GET]
// @Summary Regex Configuration
// @Description Regex Configuration
// @Accept json
// @Produce json
// @Tags Configuration
// @Param regex query string true "Regex"
// @Param value query string true "Value"
// @Success 200 {object} http.Response{data=[]models.Configuration}  "Success"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) RegexConfiguration(c *fiber.Ctx) error {
	r := regexp.MustCompile(c.Query("regex"))
	value := c.Query("value")
	resp := r.MatchString(value)
	return h.handleResponse(c, http.OK, resp, "get_regex_configuration", "")
}
