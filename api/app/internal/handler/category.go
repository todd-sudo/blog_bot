package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/todd-sudo/blog_bot/api/internal/dto"
	"github.com/todd-sudo/blog_bot/api/internal/helper"
	"github.com/todd-sudo/blog_bot/api/internal/model"
)

func (c *Handler) AllCategory(ctx *gin.Context) {
	userTgId := ctx.GetHeader("tg_user_id")
	convUserTgId, err := strconv.Atoi(userTgId)
	if err != nil {
		c.log.Error(err)
	}
	categories, err := c.service.Category.All(ctx, convUserTgId)
	if err != nil {
		c.log.Errorf("get all categories error: %v", err)
	}
	res := helper.BuildResponse(true, "OK", categories)
	ctx.JSON(http.StatusOK, res)
}

func (c *Handler) InsertCategory(ctx *gin.Context) {
	var categoryCreateDTO dto.CreateCategoryDTO
	errDTO := ctx.ShouldBind(&categoryCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
	} else {
		category, err := c.service.Category.Insert(ctx, categoryCreateDTO)
		if err != nil {
			c.log.Errorf("insert category error: %v", err)
		}
		response := helper.BuildResponse(true, "OK", category)
		ctx.JSON(http.StatusCreated, response)
	}
}

func (c *Handler) DeleteCategory(ctx *gin.Context) {
	var category model.Category

	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed tou get id", "No param id were found", helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
	}
	category.ID = id
	userTgID := ctx.GetHeader("user_id")
	isAllowedToEdit, err := c.service.Category.IsAllowedToEdit(ctx, userTgID, category.ID)

	if err != nil {
		c.log.Errorf("is allowed to edit error: %v", err)
		response := helper.BuildErrorResponse("is allowed to edit error", err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusForbidden, response)
	}

	c.log.Info(isAllowedToEdit)

	if isAllowedToEdit {
		c.service.Category.Delete(ctx, category)
		res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
		ctx.JSON(http.StatusOK, res)
	} else {
		response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
		ctx.JSON(http.StatusForbidden, response)
	}

}
