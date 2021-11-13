package v1

import (
	"net/http"

	sz "acgfate/serializer"
	"acgfate/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func WordPost(c *gin.Context) {
	form := new(service.WordPostService)
	if err := c.ShouldBind(form); err == nil {
		res := form.Post(c)
		c.JSON(http.StatusOK, res)
	} else {
		zap.S().Info(err)
		c.JSON(http.StatusOK, sz.ParamError())
	}
}

func WordDetail(c *gin.Context) {
	serv := new(service.WordDetailService)
	res := serv.Detail(c)
	c.JSON(http.StatusOK, res)
}

func WordList(c *gin.Context) {
	form := new(service.WordListService)
	if err := c.ShouldBind(form); err == nil {
		res := form.List(c)
		c.JSON(http.StatusOK, res)
	} else {
		zap.S().Info(err)
		c.JSON(http.StatusOK, sz.ParamError())
	}
}

func WordDelete(c *gin.Context) {
	serv := new(service.WordDeleteService)
	res := serv.Delete(c)
	c.JSON(http.StatusOK, res)
}
