package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"goblog.com/goblog/internal/goblog/mapper"
	"goblog.com/goblog/internal/goblog/service"
	"goblog.com/goblog/internal/goblog/util"
)

type PostHandler interface {
	GetAll(c *gin.Context)
	GetOne(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	AddImage(c *gin.Context)
	Delete(c *gin.Context)
}

type PostHandlerImpl struct {
	Service      service.PostService
	StoreService service.StoreService
}

func NewPostHandler(service service.PostService, StoreService service.StoreService) PostHandler {
	return &PostHandlerImpl{Service: service, StoreService: StoreService}
}

func (h *PostHandlerImpl) GetAll(c *gin.Context) {
	posts, err := h.Service.GetAll()
	util.CreateResponseIfError(c, err)
	response := util.FromModelToDtoList(posts, mapper.FromPostToDto)
	c.JSON(http.StatusOK, response)
}

func (h *PostHandlerImpl) GetOne(c *gin.Context) {
	post, err := h.Service.GetOne(util.GetPathID(c))
	util.CreateResponseIfError(c, err)
	response := mapper.FromPostToDto(post)
	c.JSON(http.StatusOK, response)
}

func (h *PostHandlerImpl) Create(c *gin.Context) {
	body := util.GetBodyAndMapToModel(c, mapper.FromDtoToPost)
	body.UserID = util.GetUserId(c)
	post, err := h.Service.Create(body)
	util.CreateResponseIfError(c, err)
	c.JSON(http.StatusOK, mapper.FromPostToDto(post))
}

func (h *PostHandlerImpl) Update(c *gin.Context) {
	body := util.GetBodyAndMapToModel(c, mapper.FromDtoToPost)
	body.ID = util.GetPathID(c)
	post, err := h.Service.Update(body)
	util.CreateResponseIfError(c, err)
	c.JSON(http.StatusOK, mapper.FromPostToDto(post))
}

func (h *PostHandlerImpl) AddImage(c *gin.Context) {
	fileHeader, _ := c.FormFile("file")
	h.Service.AddImage(util.GetPathID(c), fileHeader)
	c.JSON(http.StatusNoContent, nil)
}

func (h *PostHandlerImpl) Delete(c *gin.Context) {
	id := util.GetPathID(c)
	post, err := h.Service.Delete(id)
	util.CreateResponseIfError(c, err)
	c.JSON(http.StatusOK, mapper.FromPostToDto(post))
}
