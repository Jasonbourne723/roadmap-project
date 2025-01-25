package api

import (
	"roadmap/url-shortening/internal/api/requests"
	"roadmap/url-shortening/internal/services"

	"github.com/gin-gonic/gin"
)

type UrlApi struct{}

func (u *UrlApi) Add(c *gin.Context) {
	var req requests.CreateUrl
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message": "参数错误",
		})
		return
	}
	item, err := services.UrlSvc.Add(req.Url)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, item)
}

func (u *UrlApi) Del(c *gin.Context) {

}

func (u *UrlApi) Redirect(c *gin.Context) {
	p := c.Param("short")
	if len(p) <= 1 {
		c.JSON(404, gin.H{
			"message": "url is not exist",
		})
	}
	p = p[1:]

	item, err := services.UrlSvc.Get(p)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.Redirect(301, item.Long)
}
