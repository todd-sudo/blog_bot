package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/todd-sudo/blog_bot/api/internal/dto"
	"github.com/todd-sudo/blog_bot/api/internal/helper"
	"github.com/todd-sudo/blog_bot/api/internal/model"
)

func (c *Handler) AllPost(ctx *gin.Context) {
	userIdString := ctx.GetHeader("user_id")
	userId, err := strconv.Atoi(userIdString)
	if err != nil {
		c.log.Error(err)
		response := helper.BuildErrorResponse("Error", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	posts, err := c.service.Post.All(ctx, userId)
	if err != nil {
		c.log.Errorf("get all posts error: %v", err)
		response := helper.BuildErrorResponse("get all posts error", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	res := helper.BuildResponse(true, "OK", posts)
	ctx.JSON(http.StatusOK, res)
}

// Создание Post
func (c *Handler) InsertPost(ctx *gin.Context) {
	var postCreateDTO dto.PostCreateDTO
	errDTO := ctx.ShouldBind(&postCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse(
			"Failed to process request",
			errDTO.Error(),
			helper.EmptyObj{},
		)
		ctx.JSON(http.StatusBadRequest, res)
		return
	} else {
		post, err := c.service.Post.Insert(ctx, postCreateDTO)
		if err != nil {
			c.log.Errorf("insert post error: %v", err)
			res := helper.BuildErrorResponse("insert post error", err.Error(), helper.EmptyObj{})
			ctx.JSON(http.StatusBadRequest, res)
			return
		}
		response := helper.BuildResponse(true, "OK", post)
		ctx.JSON(http.StatusCreated, response)
	}
}

// Удаление Post
func (c *Handler) DeletePost(ctx *gin.Context) {
	var post model.Post
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse(
			"Failed tou get id",
			"No param id were found",
			helper.EmptyObj{},
		)
		ctx.JSON(http.StatusBadRequest, response)
	}
	post.ID = id
	userID := ctx.GetHeader("user_id")

	isAllowedToEdit, err := c.service.Post.IsAllowedToEdit(ctx, userID, post.ID)
	c.log.Info(isAllowedToEdit)
	if err != nil {
		c.log.Errorf("is allowed to edit error: %v", err)
		response := helper.BuildErrorResponse(
			"is allowed to edit error",
			err.Error(),
			helper.EmptyObj{},
		)
		ctx.JSON(http.StatusForbidden, response)
	}

	if isAllowedToEdit {
		c.service.Post.Delete(ctx, post)
		res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
		ctx.JSON(http.StatusOK, res)
	} else {
		response := helper.BuildErrorResponse(
			"You dont have permission",
			"You are not the owner",
			helper.EmptyObj{},
		)
		ctx.JSON(http.StatusForbidden, response)
		return
	}
}
