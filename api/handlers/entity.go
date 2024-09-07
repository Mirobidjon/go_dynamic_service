package handlers

import (
	"fmt"
	"strings"

	"github.com/mirobidjon/go_dynamic_service/api/http"
	"github.com/mirobidjon/go_dynamic_service/api/models"
	"github.com/mirobidjon/go_dynamic_service/genproto/dynamic_service"
	"github.com/mirobidjon/go_dynamic_service/pkg/helper"

	fiber "github.com/gofiber/fiber/v2"
	"google.golang.org/protobuf/types/known/structpb"
)

// CreateEntity godoc
// @Security ApiKeyAuth
// @ID create_entity
// @Router /client-api/entity/{slug}/create [POST]
// @Param location header string false "Location" default("Asia/Tashkent")
// @Summary Create Entity
// @Description Create Entity
// @Accept json
// @Produce json
// @Tags Entity
// @Param slug path string true "Entity Slug"
// @Param entity body models.Entity true "Create Entity"
// @Success 201 {object} http.Response{data=models.Entity} "Created"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreateEntity(c *fiber.Ctx) error {
	slug := c.Params("slug")
	if slug == "" {
		return h.handleResponse(c, http.BadRequest, nil, "invalid_slug", "slug is required")
	}

	resp, err := h.createEntity(c, c.Body(), slug)
	if err != nil {
		return h.handleResponse(c, http.InternalServerError, nil, "create_entity", err.Error())
	}

	response := resp.GetData().AsMap()

	return h.handleResponse(c, http.Created, response, "create_entity", "")
}

// UpdateEntity godoc
// @Security ApiKeyAuth
// @ID update_entity
// @Router /client-api/entity/{slug}/update/{id} [PUT]
// @Summary Update Entity
// @Description Update Entity
// @Accept json
// @Produce json
// @Tags Entity
// @Param slug path string true "Entity Slug"
// @Param id path string true "Entity ID"
// @Param entity body models.Entity true "Update Entity"
// @Success 200 {object} http.Response{data=models.Entity}
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateEntity(c *fiber.Ctx) error {
	slug := c.Params("slug")
	if slug == "" {
		return h.handleResponse(c, http.BadRequest, nil, "invalid_slug", "slug is required")
	}

	var entity = dynamic_service.Entity{
		Data: &structpb.Struct{},
	}
	entity.Slug = slug
	entity.XId = c.Params("id")
	entity.Location = c.Get("location")

	if entity.XId == "" {
		return h.handleResponse(c, http.BadRequest, nil, "invalid_id", "id is required")
	}

	if err := entity.Data.UnmarshalJSON(c.Body()); err != nil {
		return h.handleResponse(c, http.BadRequest, nil, "unmarshal_entity", err.Error())
	}

	resp, err := h.services.EntityService().Update(c.Context(), &entity)
	if err != nil {
		return h.handleResponse(c, http.InternalServerError, nil, "update_entity", err.Error())
	}

	response := resp.Data.AsMap()

	return h.handleResponse(c, http.OK, response, "update_entity", "")
}

// PatchUpdateEntity godoc
// @Security ApiKeyAuth
// @ID update_entity
// @Router /client-api/entity/{slug}/update/{id} [PATCH]
// @Summary Patch Update Entity
// @Description Patch Update Entity
// @Accept json
// @Produce json
// @Tags Entity
// @Param slug path string true "Entity Slug"
// @Param id path string true "Entity ID"
// @Param entity body models.Entity true "Update Entity"
// @Success 200 {object} http.Response{data=models.Entity}
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) PatchUpdateEntity(c *fiber.Ctx) error {
	slug := c.Params("slug")
	if slug == "" {
		return h.handleResponse(c, http.BadRequest, nil, "invalid_slug", "slug is required")
	}

	resp, err := h.patchUpdateEntity(c, c.Body(), slug, c.Params("id"))
	if err != nil {
		return h.handleResponse(c, http.InternalServerError, nil, "update_entity", err.Error())
	}

	response := resp.Data.AsMap()

	return h.handleResponse(c, http.OK, response, "update_entity", "")
}

// DeleteEntity godoc
// @Security ApiKeyAuth
// @ID delete_entity
// @Router /client-api/entity/{slug}/delete/{id} [DELETE]
// @Param location header string false "Location" default("Asia/Tashkent")
// @Summary Delete Entity
// @Description Delete Entity
// @Accept json
// @Produce json
// @Tags Entity
// @Param slug path string true "Entity Slug"
// @Param id path string true "Entity ID"
// @Success 200 {object} http.Response{data=string} "Deleted"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeleteEntity(c *fiber.Ctx) error {
	slug := c.Params("slug")
	if slug == "" {
		return h.handleResponse(c, http.BadRequest, nil, "invalid_slug", "slug is required")
	}

	var entity dynamic_service.Entity
	entity.Slug = slug
	entity.XId = c.Params("id")
	entity.Location = c.Get("location")

	if entity.XId == "" {
		return h.handleResponse(c, http.BadRequest, nil, "invalid_id", "id is required")
	}

	if _, err := h.services.EntityService().Delete(c.Context(), &entity); err != nil {
		return h.handleResponse(c, http.InternalServerError, nil, "delete_entity", err.Error())
	}

	return h.handleResponse(c, http.OK, "Successfully deleted", "delete_entity", "")
}

// GetEntity godoc
// @Security ApiKeyAuth
// @ID get_entity
// @Router /client-api/entity/{slug}/get/{id} [GET]
// @Param location header string false "Location" default("Asia/Tashkent")
// @Summary Get Entity
// @Description Get Entity
// @Accept json
// @Produce json
// @Tags Entity
// @Param slug path string true "Entity Slug"
// @Param id path string true "Entity ID"
// @Success 200 {object} http.Response{data=models.Entity} "OK"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetEntity(c *fiber.Ctx) error {
	slug := c.Params("slug")
	if slug == "" {
		return h.handleResponse(c, http.BadRequest, nil, "invalid_slug", "slug is required")
	}

	var entity dynamic_service.GetByPk
	entity.Slug = slug
	entity.XId = c.Params("id")
	entity.Location = c.Get("location")

	if entity.XId == "" {
		return h.handleResponse(c, http.BadRequest, nil, "invalid_id", "id is required")
	}

	resp, err := h.services.EntityService().GetById(c.Context(), &entity)
	if err != nil {
		return h.handleResponse(c, http.InternalServerError, nil, "get_entity", err.Error())
	}

	response := resp.Data.AsMap()

	return h.handleResponse(c, http.OK, response, "get_entity", "")
}

// GetAll godoc
// @Security ApiKeyAuth
// @ID get_all_post
// @Router /client-api/entity/{slug}/get-all [POST]
// @Param location header string false "Location" default("Asia/Tashkent")
// @Summary Get All Entities
// @Description Get All Entities
// @Accept json
// @Produce json
// @Tags Entity
// @Param slug path string true "Entity Slug"
// @Param offset query int false "Offset"
// @Param limit query int false "Limit"
// @Param sort query string false "Sort"
// @Param order query string false "Order"
// @Param search query string false "Search"
// @Param filter body models.Entity false "Filter Entity"
// @Success 200 {object} http.Response{data=models.Entity} "Entities"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetAllEntityPost(c *fiber.Ctx) error {
	slug := c.Params("slug")
	if slug == "" {
		return h.handleResponse(c, http.BadRequest, nil, "invalid_slug", "slug is required")
	}

	offset, err := h.getOffsetParam(c)
	if err != nil {
		return h.handleResponse(c, http.BadRequest, nil, "get_offset", err.Error())
	}

	limit, err := h.getLimitParam(c)
	if err != nil {
		return h.handleResponse(c, http.BadRequest, nil, "get_limit", err.Error())
	}

	var filter = dynamic_service.GetAllRequest{
		Data: &structpb.Struct{},
	}
	filter.Slug = slug
	filter.Offset = int32(offset)
	filter.Limit = int32(limit)
	filter.Sort = c.Query("sort")
	filter.Order = c.Query("order")
	filter.Search = c.Query("search")
	filter.Location = c.Get("location")

	if len(c.Body()) > 5 {
		if err := filter.Data.UnmarshalJSON(c.Body()); err != nil {
			return h.handleResponse(c, http.BadRequest, nil, "unmarshal_entity", err.Error())
		}
	}

	resp, err := h.services.EntityService().GetAll(c.Context(), &filter)
	if err != nil {
		return h.handleResponse(c, http.InternalServerError, nil, "get_entity_list", err.Error())
	}

	var response models.GetAllEntityResponse

	for _, entity := range resp.Entities {
		response.Entities = append(response.Entities, entity.AsMap())
	}

	response.Count = resp.Count

	return h.handleResponse(c, http.OK, response, "get_entity_list", "")
}

// GetAllEntityGet godoc
// @Security ApiKeyAuth
// @ID get_all_get
// @Router /client-api/entity/{slug}/get-all [GET]
// @Param location header string false "Location" default("Asia/Tashkent")
// @Summary Get All Entities
// @Description Get All Entities
// @Accept json
// @Produce json
// @Tags Entity
// @Param slug path string true "Entity Slug"
// @Param offset query int false "Offset"
// @Param limit query int false "Limit"
// @Param sort query string false "Sort"
// @Param order query string false "Order"
// @Param search query string false "Search"
// @Success 200 {object} http.Response{data=models.Entity} "Entities"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetAllEntityGet(c *fiber.Ctx) error {
	slug := c.Params("slug")
	if slug == "" {
		return h.handleResponse(c, http.BadRequest, nil, "invalid_slug", "slug is required")
	}

	offset, err := h.getOffsetParam(c)
	if err != nil {
		return h.handleResponse(c, http.BadRequest, nil, "get_offset", err.Error())
	}

	limit, err := h.getLimitParam(c)
	if err != nil {
		return h.handleResponse(c, http.BadRequest, nil, "get_limit", err.Error())
	}

	var filter = dynamic_service.GetAllRequest{
		Data: &structpb.Struct{},
	}
	filter.Slug = slug
	filter.Offset = int32(offset)
	filter.Limit = int32(limit)
	filter.Sort = c.Query("sort")
	filter.Order = c.Query("order")
	filter.Search = c.Query("search")
	filter.Location = c.Get("location")

	resp, err := h.services.EntityService().GetAll(c.Context(), &filter)
	if err != nil {
		return h.handleResponse(c, http.InternalServerError, nil, "get_entity_list", err.Error())
	}

	var response models.GetAllEntityResponse

	for _, entity := range resp.Entities {
		response.Entities = append(response.Entities, entity.AsMap())
	}

	response.Count = resp.Count

	return h.handleResponse(c, http.OK, response, "get_entity_list", "")
}

// GetJoin godoc
// @Security ApiKeyAuth
// @ID get_join
// @Router /client-api/entity/{slug}/get-join [POST]
// @Param location header string false "Location" default("Asia/Tashkent")
// @Summary Get Join Entities
// @Description Get Join Entities
// @Accept json
// @Produce json
// @Tags Entity
// @Param slug path string true "Entity Slug"
// @Param offset query int false "Offset"
// @Param limit query int false "Limit"
// @Param sort query string false "Sort"
// @Param order query string false "Order"
// @Param search query string false "Search"
// @Param filter body models.JoinGroupRequest false "Filter Entity"
func (h *Handler) GetJoinEntity(c *fiber.Ctx) error {
	slug := c.Params("slug")
	if slug == "" {
		return h.handleResponse(c, http.BadRequest, nil, "invalid_slug", "slug is required")
	}

	offset, err := h.getOffsetParam(c)
	if err != nil {
		return h.handleResponse(c, http.BadRequest, nil, "get_offset", err.Error())
	}

	limit, err := h.getLimitParam(c)
	if err != nil {
		return h.handleResponse(c, http.BadRequest, nil, "get_limit", err.Error())
	}

	var filter = dynamic_service.GetJoinRequest{
		Data: &structpb.Struct{},
	}

	if err := c.BodyParser(&filter); err != nil {
		return h.handleResponse(c, http.BadRequest, nil, "body_parse", err.Error())
	}

	filter.Slug = slug
	filter.Offset = int32(offset)
	filter.Limit = int32(limit)
	filter.Sort = c.Query("sort")
	filter.Order = c.Query("order")
	filter.Search = c.Query("search")
	filter.Location = c.Get("location")

	if len(c.Body()) > 5 {
		if err := filter.Data.UnmarshalJSON(c.Body()); err != nil {
			return h.handleResponse(c, http.BadRequest, nil, "unmarshal_entity", err.Error())
		}

		delete(filter.Data.Fields, "lookups")
		delete(filter.Data.Fields, "aggregate")
	}

	resp, err := h.services.EntityService().GetJoin(c.Context(), &filter)
	if err != nil {
		return h.handleResponse(c, http.InternalServerError, nil, "join_entity", err.Error())
	}

	var response models.GetAllEntityResponse

	for _, entity := range resp.Entities {
		response.Entities = append(response.Entities, entity.AsMap())
	}

	response.Count = resp.Count

	return h.handleResponse(c, http.OK, response, "join_entity", "")
}

// GetAllService godoc
// @Security ApiKeyAuth
// @ID get_all_service
// @Router /client-api/service/{slug}/get-all [GET]
// @Param location header string false "Location" default("Asia/Tashkent")
// @Summary Get All Services
// @Description Get All Services
// @Accept json
// @Produce json
// @Tags Entity
// @Param slug path string true "Entity Slug"
// @Param offset query int false "Offset"
// @Param limit query int false "Limit"
// @Param sort query string false "Sort"
// @Param order query string false "Order"
// @Param search query string false "Search"
// @Success 200 {object} http.Response{data=models.Entity} "Services"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetAllServices(c *fiber.Ctx) error {
	slug := c.Params("slug")
	if slug == "" {
		return h.handleResponse(c, http.BadRequest, nil, "invalid_slug", "slug is required")
	}

	offset, err := h.getOffsetParam(c)
	if err != nil {
		return h.handleResponse(c, http.BadRequest, nil, "get_offset", err.Error())
	}

	limit, err := h.getLimitParam(c)
	if err != nil {
		return h.handleResponse(c, http.BadRequest, nil, "get_limit", err.Error())
	}

	group, err := h.services.DynamicService().GetFullGroup(
		c.Context(),
		&dynamic_service.GetByIdRequest{
			XId: slug,
		},
	)

	if err != nil {
		return h.handleResponse(c, http.NotFound, nil, "get_full_group", err.Error())
	}

	if group == nil || group.Status != 2 {
		return h.handleResponse(c, http.NotFound, nil, "get_full_group", "valid group not found")
	}

	var filter = dynamic_service.GetAllRequest{
		Data: &structpb.Struct{},
	}
	filter.Slug = slug
	filter.Offset = int32(offset)
	filter.Limit = int32(limit)
	filter.Sort = c.Query("sort")
	filter.Order = c.Query("order")
	filter.Search = c.Query("search")
	filter.Location = c.Get("location")

	resp, err := h.services.EntityService().GetAll(c.Context(), &filter)
	if err != nil {
		return h.handleResponse(c, http.InternalServerError, nil, "get_entity_list", err.Error())
	}

	var response models.GetAllEntityResponse

	for _, entity := range resp.Entities {
		response.Entities = append(response.Entities, entity.AsMap())
	}

	response.Count = resp.Count

	return h.handleResponse(c, http.OK, response, "get_entity_list", "")
}

// GetEntityService godoc
// @Security ApiKeyAuth
// @ID get_entity_service
// @Router /client-api/service/{slug}/get/{id} [GET]
// @Param location header string false "Location" default("Asia/Tashkent")
// @Summary Get Entity
// @Description Get Entity
// @Accept json
// @Produce json
// @Tags Entity
// @Param slug path string true "Entity Slug"
// @Param id path string true "Entity ID"
// @Success 200 {object} http.Response{data=models.Entity} "OK"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetEntityService(c *fiber.Ctx) error {
	slug := c.Params("slug")
	if slug == "" {
		return h.handleResponse(c, http.BadRequest, nil, "invalid_slug", "slug is required")
	}

	var entity dynamic_service.GetByPk
	entity.Slug = slug
	entity.XId = c.Params("id")
	entity.Location = c.Get("location")

	if entity.XId == "" {
		return h.handleResponse(c, http.BadRequest, nil, "invalid_id", "id is required")
	}

	resp, err := h.services.EntityService().GetById(c.Context(), &entity)
	if err != nil {
		return h.handleResponse(c, http.InternalServerError, nil, "get_entity", err.Error())
	}

	response := resp.Data.AsMap()

	return h.handleResponse(c, http.OK, response, "get_entity", "")
}

// GetExampleEntity godoc
// @Security ApiKeyAuth
// @ID get_example_entity
// @Router /client-api/entity/{slug}/example [GET]
// @Summary Get Example Entity
// @Description Get Example Entity
// @Accept json
// @Produce json
// @Tags Entity
// @Param slug path string true "Entity Slug"
// @Success 200 {object} http.Response{data=models.Entity} "OK"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetExampleEntity(c *fiber.Ctx) error {
	var group models.DynamicGroup

	slug := c.Params("slug")
	if slug == "" {
		return h.handleResponse(c, http.BadRequest, nil, "invalid_slug", "slug is required")
	}

	resp, err := h.services.DynamicService().GetFullGroup(c.Context(), &dynamic_service.GetByIdRequest{XId: slug})
	if err != nil {
		return h.handleResponse(c, http.InternalServerError, nil, "get_full_group", err.Error())
	}

	if err := helper.ProtoToStruct(&group, resp); err != nil {
		return h.handleResponse(c, http.InternalServerError, nil, "proto_to_struct", err.Error())
	}

	return h.handleResponse(c, http.OK, group, "default_success_message", "")
}

func (h *Handler) createEntity(c *fiber.Ctx, body any, slug string) (*dynamic_service.Entity, error) {
	var (
		err    error
		entity dynamic_service.Entity
		resp   *dynamic_service.Entity
	)

	if slug == "" {
		return nil, fmt.Errorf("slug is required")
	}

	entity.Slug = slug
	entity.Location = c.Get("location")

	switch actualBody := body.(type) {
	case string:
		entity.Data, err = helper.StringToStructPb(actualBody)
	case []byte:
		entity.Data, err = helper.ByteToStructPb(actualBody)
	case map[string]interface{}:
		entity.Data, err = helper.ToProtoStruct(actualBody)
	case map[string]string:
		entity.Data, err = helper.ToProtoStruct(actualBody)
	default:
		return nil, fmt.Errorf("unsupported body type: %T", actualBody)
	}

	if err != nil {
		return nil, err
	}

	resp, err = h.services.EntityService().Create(c.Context(), &entity)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h *Handler) patchUpdateEntity(c *fiber.Ctx, body any, slug string, xIds ...string) (*dynamic_service.Entity, error) {
	var (
		err    error
		entity dynamic_service.Entity
		resp   *dynamic_service.Entity
	)

	if slug == "" {
		return nil, fmt.Errorf("slug is required")
	}

	entity.Slug = slug
	entity.XId = strings.Join(xIds, ",")
	entity.Location = c.Get("location")

	switch actualBody := body.(type) {
	case string:
		entity.Data, err = helper.StringToStructPb(actualBody)
	case []byte:
		entity.Data, err = helper.ByteToStructPb(actualBody)
	case map[string]interface{}:
		entity.Data, err = helper.ToProtoStruct(actualBody)
	case map[string]string:
		entity.Data, err = helper.ToProtoStruct(actualBody)
	default:
		return nil, fmt.Errorf("unsupported body type: %T", actualBody)
	}

	if err != nil {
		return nil, err
	}

	resp, err = h.services.EntityService().UpdatePatch(c.Context(), &entity)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
