package lib

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type FetchUrlRequest struct {
	Urls []string `form:"urls" json:"urls"`
}

type FetchUrlResponse struct {
	Websites	[]Website `json:"websites"`
}

func FetchUrlHandler(c *gin.Context) {
	var req FetchUrlRequest
	var res FetchUrlResponse
	c.Bind(&req)

	ch := Fetch(req.Urls)
	for i := 0; i < len(req.Urls); i++ {
		website := <-ch
		res.Websites = append(res.Websites, website)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": res,
		"message": "ok",
	})
}