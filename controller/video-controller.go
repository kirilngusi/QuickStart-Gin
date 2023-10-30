package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/kirilngusi/QuickStart-Gin/entity"
	"github.com/kirilngusi/QuickStart-Gin/service"
	"github.com/kirilngusi/QuickStart-Gin/templates"
	"github.com/kirilngusi/QuickStart-Gin/validators"
	"html/template"
	"net/http"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(content *gin.Context) error
	ShowAll(content *gin.Context)
}

type controller struct {
	service service.VideoService
}

func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}

func (c *controller) Save(content *gin.Context) error {
	var video entity.Video
	err := content.ShouldBindJSON(&video)

	if err != nil {
		content.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"error":   "VALIDATEERR-1",
				"message": "Invalid inputs. Please check your inputs"})
		return nil
	}

	err = validate.Struct(video)

	if err != nil {
		return err
	}

	c.service.Save(video)
	return nil
}

var validate *validator.Validate

func New(service service.VideoService) VideoController {

	validate = validator.New()
	validate.RegisterValidation("is-cool", validators.ValidateCoolTitle)

	return &controller{
		service: service,
	}
}

func (c *controller) ShowAll(content *gin.Context) {
	videos := c.service.FindAll()

	fmt.Println("videos", videos)

	data := gin.H{
		"title":  "Video",
		"videos": videos,
	}

	fmt.Println(data)

	_, err := template.New("").Parse(templates.Home)
	if err != nil {
		// Handle the error
		panic(err)
	}

	//return t

	content.HTML(http.StatusOK, "index.html", data)
}
