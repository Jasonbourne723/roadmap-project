package main

import (
	"net/http"
	"roadmap/url-shortening/internal/api"

	"github.com/gin-gonic/gin"
)

func main() {

	g := gin.Default()

	urlApi := &api.UrlApi{}
	g.POST("/url", urlApi.Add)
	g.GET("/*short", urlApi.Redirect)

	http.ListenAndServe(":8080", g)
}
