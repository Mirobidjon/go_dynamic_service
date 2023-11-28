package handlers

import (
	"kassa360/kassa360_go_dynamic_service/api/http"
	"kassa360/kassa360_go_dynamic_service/api/models"
	"kassa360/kassa360_go_dynamic_service/genproto/dynamic_service"
	"kassa360/kassa360_go_dynamic_service/pkg/helper"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/spf13/cast"
)

// CreateGroup godoc
// @Security ApiKeyAuth
// @ID create_group
// @Router /client-api/dynamic/group [POST]
// @Summary Create Group
// @Description Create Group
// @Accept json
// @Produce json
// @Tags Group
// @Param group body models.Group true "Create Group"
// @Success 201 {object} http.Response{data=models.Group} "Created"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreateGroup(c *fiber.Ctx) error {
	var group dynamic_service.Group

	if err := helper.ByteToProto(&group, c.Body()); err != nil {
		return h.handleResponse(c, http.BadRequest, nil, "byte_to_proto", err.Error())
	}

	resp, err := h.services.DynamicService().CreateGroup(c.Context(), &group)
	if err != nil {
		return h.handleResponse(c, http.InternalServerError, nil, "create_group", err.Error())
	}

	var data models.DynamicGroup
	if err := helper.ProtoToStruct(&data, resp); err != nil {
		return h.handleResponse(c, http.InternalServerError, nil, "proto_to_struct", err.Error())
	}

	return h.handleResponse(c, http.OK, data, "create_group", "")
}

// GetGroup godoc
// @Security ApiKeyAuth
// @ID get_group
// @Router /client-api/dynamic/group/{id} [GET]
// @Summary Get Group
// @Description Get Group
// @Accept json
// @Produce json
// @Tags Group
// @Param id path string true "Group ID"
// @Success 200 {object} http.Response{data=models.Group} "OK"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetGroup(c *fiber.Ctx) error {
	id := c.Params("id")
	resp, err := h.services.DynamicService().GetGroupById(c.Context(), &dynamic_service.GetByIdRequest{XId: id})
	if err != nil {
		return h.handleResponse(c, http.InternalServerError, nil, "get_group", err.Error())
	}

	var data models.DynamicGroup
	if err := helper.ProtoToStruct(&data, resp); err != nil {
		return h.handleResponse(c, http.InternalServerError, nil, "proto_to_struct", err.Error())
	}

	return h.handleResponse(c, http.OK, data, "get_group", "")
}

// UpdateGroup godoc
// @Security ApiKeyAuth
// @ID update_group
// @Router /client-api/dynamic/group/{id} [PUT]
// @Summary Update Group
// @Description Update Group
// @Accept json
// @Produce json
// @Tags Group
// @Param id path string true "Group ID"
// @Param group body models.Group true "Update Group"
// @Success 200 {object} http.Response{data=string} "OK"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateGroup(c *fiber.Ctx) error {
	id := c.Params("id")

	var group dynamic_service.Group
	if err := helper.ByteToProto(&group, c.Body()); err != nil {
		return h.handleResponse(c, http.BadRequest, nil, "byte_to_proto", err.Error())
	}

	group.XId = id
	resp, err := h.services.DynamicService().UpdateGroup(c.Context(), &group)
	if err != nil {
		return h.handleResponse(c, http.InternalServerError, nil, "update_group", err.Error())
	}

	var data models.DynamicGroup
	if err := helper.ProtoToStruct(&data, resp); err != nil {
		return h.handleResponse(c, http.InternalServerError, nil, "proto_to_struct", err.Error())
	}

	return h.handleResponse(c, http.OK, data, "update_group", "")
}

// DeleteGroup godoc
// @Security ApiKeyAuth
// @ID delete_group
// @Router /client-api/dynamic/group/{id} [DELETE]
// @Summary Delete Group
// @Description Delete Group
// @Accept json
// @Produce json
// @Tags Group
// @Param id path string true "Group ID"
// @Success 200 {object} string "OK"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeleteGroup(c *fiber.Ctx) error {
	id := c.Params("id")

	_, err := h.services.DynamicService().DeleteGroup(c.Context(), &dynamic_service.GetByIdRequest{XId: id})
	if err != nil {
		return h.handleResponse(c, http.InternalServerError, nil, "delete_group", err.Error())
	}

	return h.handleResponse(c, http.OK, "Successfully deleted", "delete_group", "")
}

// GetAllGroup godoc
// @Security ApiKeyAuth
// @ID get_all_group
// @Router /client-api/dynamic/group [GET]
// @Summary Get All Group
// @Description Get All Group
// @Accept json
// @Produce json
// @Tags Group
// @Param offset query int false "Offset"
// @Param limit query int false "Limit"
// @Param sort query string false "Sort"
// @Param order query string false "Order"
// @Param group_type query int false "Group Type"
// @Param slug query string false "Slug"
// @Param search query string false "Search"
// @Param parent_id query string false "Parent ID"
// @Success 200 {object} http.Response{data=models.GetAllGroupResponse} "OK"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetAllGroup(c *fiber.Ctx) error {
	offset, err := h.getOffsetParam(c)
	if err != nil {
		return h.handleResponse(c, http.BadRequest, nil, "get_offset", err.Error())
	}

	limit, err := h.getLimitParam(c)
	if err != nil {
		return h.handleResponse(c, http.BadRequest, nil, "get_limit", err.Error())
	}

	sort := c.Query("sort", "created_at")
	order := c.Query("order", "")
	groupType := cast.ToInt32(c.Query("group_type", ""))

	resp, err := h.services.DynamicService().GetAllGroup(c.Context(), &dynamic_service.GetAllGroupRequest{
		Offset:    int32(offset),
		Limit:     int32(limit),
		Sort:      sort,
		Order:     order,
		Search:    c.Query("search", ""),
		ParentId:  c.Query("parent_id", ""),
		GroupType: groupType,
		Slug:      c.Query("slug", ""),
	})
	if err != nil {
		return h.handleResponse(c, http.InternalServerError, nil, "get_group_list", err.Error())
	}

	var group models.GetAllDynamicGroupResponse
	if err := helper.ProtoToStruct(&group, resp); err != nil {
		return h.handleResponse(c, http.InternalServerError, nil, "proto_to_struct", err.Error())
	}

	return h.handleResponse(c, http.OK, group, "get_group_list", "")
}

// CreateField godoc
// @Security ApiKeyAuth
// @ID create_field
// @Router /client-api/dynamic/field [POST]
// @Summary Create Field
// @Description Create Field
// @Accept json
// @Produce json
// @Tags Field
// @Param field body dynamic_service.Field true "Create Field"
// @Success 201 {object} http.Response{data=dynamic_service.Field} "Created"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreateField(c *fiber.Ctx) error {
	var field dynamic_service.Field
	if err := c.BodyParser(&field); err != nil {
		return h.handleResponse(c, http.BadRequest, nil, "body_parse", err.Error())
	}

	resp, err := h.services.DynamicService().CreateField(c.Context(), &field)
	if err != nil {
		return h.handleResponse(c, http.InternalServerError, nil, "create_field", err.Error())
	}

	return h.handleResponse(c, http.OK, resp, "create_field", "")
}

// UpdateField godoc
// @Security ApiKeyAuth
// @ID update_field
// @Router /client-api/dynamic/field/{id} [PUT]
// @Summary Update Field
// @Description Update Field
// @Accept json
// @Produce json
// @Tags Field
// @Param id path string true "Field ID"
// @Param field body dynamic_service.Field true "Update Field"
// @Success 200 {object} http.Response{data=dynamic_service.Field} "OK"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateField(c *fiber.Ctx) error {
	id := c.Params("id")

	var field dynamic_service.Field
	if err := c.BodyParser(&field); err != nil {
		return h.handleResponse(c, http.BadRequest, nil, "body_parse", err.Error())
	}

	field.XId = id
	resp, err := h.services.DynamicService().UpdateField(c.Context(), &field)
	if err != nil {
		return h.handleResponse(c, http.InternalServerError, nil, "update_field", err.Error())
	}

	return h.handleResponse(c, http.OK, resp, "update_field", "")
}

// DeleteField godoc
// @Security ApiKeyAuth
// @ID delete_field
// @Router /client-api/dynamic/field/{id} [DELETE]
// @Summary Delete Field
// @Description Delete Field
// @Accept json
// @Produce json
// @Tags Field
// @Param id path string true "Field ID"
// @Success 200 {object} http.Response{data=string} "OK"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeleteField(c *fiber.Ctx) error {
	id := c.Params("id")

	_, err := h.services.DynamicService().DeleteField(c.Context(), &dynamic_service.GetByIdRequest{XId: id})
	if err != nil {
		return h.handleResponse(c, http.InternalServerError, nil, "delete_field", err.Error())
	}

	return h.handleResponse(c, http.OK, "Successfully deleted", "delete_field", "")
}

// GetAllField godoc
// @Security ApiKeyAuth
// @ID get_all_field
// @Router /client-api/dynamic/field [GET]
// @Summary Get All Field
// @Description Get All Field
// @Accept json
// @Produce json
// @Tags Field
// @Param offset query int false "Offset"
// @Param limit query int false "Limit"
// @Param sort query string false "Sort"
// @Param order query string false "Order"
// @Param group_id query string false "Group ID"
// @Param search query string false "Search"
// @Param slug query string false "Slug"
// @Success 200 {object} http.Response{data=dynamic_service.GetAllFieldResponse} "OK"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetAllField(c *fiber.Ctx) error {
	offset, err := h.getOffsetParam(c)
	if err != nil {
		return h.handleResponse(c, http.BadRequest, nil, "get_offset", err.Error())
	}

	limit, err := h.getLimitParam(c)
	if err != nil {
		return h.handleResponse(c, http.BadRequest, nil, "get_limit", err.Error())
	}

	sort := c.Query("sort", "created_at")
	order := c.Query("order", "")

	resp, err := h.services.DynamicService().GetAllField(c.Context(), &dynamic_service.GetAllFieldRequest{
		Offset:  int32(offset),
		Limit:   int32(limit),
		Sort:    sort,
		Order:   order,
		GroupId: c.Query("group_id", ""),
		Search:  c.Query("search", ""),
		Slug:    c.Query("slug", ""),
	})
	if err != nil {
		return h.handleResponse(c, http.InternalServerError, nil, "get_field_list", err.Error())
	}

	return h.handleResponse(c, http.OK, resp, "get_field_list", "")
}

// GetFieldById godoc
// @Security ApiKeyAuth
// @ID get_field_by_id
// @Router /client-api/dynamic/field/{id} [GET]
// @Summary Get Field By ID
// @Description Get Field By ID
// @Accept json
// @Produce json
// @Tags Field
// @Param id path string true "Field ID"
// @Success 200 {object} http.Response{data=dynamic_service.Field} "OK"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetFieldById(c *fiber.Ctx) error {
	id := c.Params("id")

	resp, err := h.services.DynamicService().GetFieldById(c.Context(), &dynamic_service.GetByIdRequest{XId: id})
	if err != nil {
		return h.handleResponse(c, http.InternalServerError, nil, "get_field", err.Error())
	}

	return h.handleResponse(c, http.OK, resp, "get_field", "")
}

// GetFullGroup godoc
// @Security ApiKeyAuth
// @ID get_full_group
// @Router /client-api/dynamic/group/{slug}/full [GET]
// @Summary Get Full Group
// @Description Get Full Group
// @Accept json
// @Produce json
// @Tags Group
// @Param slug path string true "Group Slug Or ID"
// @Success 200 {object} http.Response{data=models.Group} "OK"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetFullGroup(c *fiber.Ctx) error {
	var group models.DynamicGroup

	id := c.Params("slug")

	if id == "" {
		return h.handleResponse(c, http.BadRequest, nil, "invalid_id", "ID is required")
	}

	resp, err := h.services.DynamicService().GetFullGroup(c.Context(), &dynamic_service.GetByIdRequest{XId: id})
	if err != nil {
		return h.handleResponse(c, http.InternalServerError, nil, "get_full_group", err.Error())
	}

	if err := helper.ProtoToStruct(&group, resp); err != nil {
		return h.handleResponse(c, http.InternalServerError, nil, "proto_to_struct", err.Error())
	}

	return h.handleResponse(c, http.OK, group, "get_full_group", "")
}
