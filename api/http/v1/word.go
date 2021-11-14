package v1

import (
	"net/http"

	sz "acgfate/serializer"
	"acgfate/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// WordPost 文字投稿接口
// @Summary 文字投稿接口
// @Description 文字投稿接口
// @Tags Word
// @Accept application/json
// @Produce application/json
// @Param word body service.WordPostService true "投稿信息"
// @Success 0 {object} serializer.Response "成功"
// @Failure 40000 {object} serializer.Response "操作错误"
// @Failure 50000 {object} serializer.Response "服务器错误"
// @Router /word/post [post]
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

// WordDetail 文字详情接口
// @Summary 文字详情接口，展示文字的相关信息
// @Description 文字详情接口，展示文字的相关信息
// @Tags Word
// @Produce application/json
// @Param wid path int64 true "文字ID"
// @Success 0 {array} []serializer.Word "成功"
// @Failure 40000 {object} serializer.Response "操作错误"
// @Failure 50000 {object} serializer.Response "服务器错误"
// @Router /word/{wid} [get]
func WordDetail(c *gin.Context) {
	serv := new(service.WordDetailService)
	res := serv.Detail(c)
	c.JSON(http.StatusOK, res)
}

// WordList 文字列表接口
// @Summary 文字列表接口，展示多个文字的相关信息
// @Description 文字列表接口，展示多个文字的相关信息，支持分页查询，按用户ID、分区ID查询
// @Tags Word
// @Accept application/json
// @Produce application/json
// @Param param body service.WordListService false "分页/用户/分区"
// @Success 0 {array} []serializer.Word "成功"
// @Failure 40000 {object} serializer.Response "操作错误"
// @Failure 50000 {object} serializer.Response "服务器错误"
// @Router /word/list [get]
func WordList(c *gin.Context) {
	form := new(service.WordListService)
	if err := c.ShouldBindQuery(form); err == nil {
		res := form.List(c)
		c.JSON(http.StatusOK, res)
	} else {
		zap.S().Info(err)
		c.JSON(http.StatusOK, sz.ParamError())
	}
}

// WordDelete 文字删除接口
// @Summary 文字删除接口，删除一个文字
// @Description 文字删除接口，删除一个文字
// @Tags Word
// @Produce application/json
// @Param wid path int64 true "文字ID"
// @Success 0 {object} serializer.Response "成功"
// @Failure 40000 {object} serializer.Response "操作错误"
// @Failure 50000 {object} serializer.Response "服务器错误"
// @Router /word/{wid}/delete [delete]
func WordDelete(c *gin.Context) {
	serv := new(service.WordDeleteService)
	res := serv.Delete(c)
	c.JSON(http.StatusOK, res)
}

// WordLike 文字点赞接口
// @Summary 文字点赞接口
// @Description 文字点赞接口
// @Tags Word
// @Produce application/json
// @Param wid path int64 true "文字ID"
// @Success 0 {object} serializer.Response "成功"
// @Failure 40000 {object} serializer.Response "操作错误"
// @Failure 50000 {object} serializer.Response "服务器错误"
// @Router /word/{wid}/like [put]
func WordLike(c *gin.Context) {
	serv := new(service.WordLikeService)
	res := serv.Like(c)
	c.JSON(http.StatusOK, res)
}

// WordUnlike 文字取消点赞接口
// @Summary 文字取消点赞接口
// @Description 文字取消点赞接口
// @Tags Word
// @Produce application/json
// @Param wid path int64 true "文字ID"
// @Success 0 {object} serializer.Response "成功"
// @Failure 40000 {object} serializer.Response "操作错误"
// @Failure 50000 {object} serializer.Response "服务器错误"
// @Router /word/{wid}/like [delete]
func WordUnlike(c *gin.Context) {
	serv := new(service.WordLikeService)
	res := serv.Unlike(c)
	c.JSON(http.StatusOK, res)
}

// WordTrend 文字热度趋势接口
// @Summary 文字热度趋势接口
// @Description 文字热度趋势接口
// @Tags Word
// @Produce application/json
// @Success 0 {object} serializer.Trend "成功"
// @Failure 40000 {object} serializer.Response "操作错误"
// @Failure 50000 {object} serializer.Response "服务器错误"
// @Router /word/trend [get]
func WordTrend(c *gin.Context) {
	serv := new(service.WordLikeService)
	res := serv.Trend(c)
	c.JSON(http.StatusOK, res)
}
