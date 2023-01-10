package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"goblog.com/goblog/internal/goblog/mapper"
	"goblog.com/goblog/internal/goblog/service"
	"goblog.com/goblog/internal/goblog/util"
)

type CommentHandler interface {
	GetAllForPost(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
}

type CommentHandlerImpl struct {
	Service service.CommentService
}

func NewCommentHandler(service service.CommentService) CommentHandler {
	return &CommentHandlerImpl{service}
}

func (h *CommentHandlerImpl) GetAllForPost(c *gin.Context) {
	comments, err := h.Service.GetAllForPost(util.GetPathID(c))
	util.CreateResponseIfError(c, err)
	response := util.FromModelToDtoList(comments, mapper.FromCommentToDto)
	c.JSON(http.StatusOK, response)

}

func (h *CommentHandlerImpl) Create(c *gin.Context) {
	body := util.GetBodyAndMapToModel(c, mapper.FromDtoToComment)
	body.UserID = util.GetUserId(c)
	body.PostID = util.GetPathID(c)
	comment, err := h.Service.Create(body)
	util.CreateResponseIfError(c, err)
	c.JSON(http.StatusOK, mapper.FromCommentToDto(comment))

}

func (h *CommentHandlerImpl) Update(c *gin.Context) {
	body := util.GetBodyAndMapToModel(c, mapper.FromDtoToComment)
	body.UserID = util.GetUserId(c)
	body.ID = util.GetNamedID(c, "commentId")
	comment, err := h.Service.Update(body)
	util.CreateResponseIfError(c, err)
	c.JSON(http.StatusOK, mapper.FromCommentToDto(comment))

}
