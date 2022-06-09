package server

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func NewDownStreamServer(successRate float64) *gin.Engine {
	if successRate > 1 || successRate < 0 {
		panic("invalid input")
	}
	app := gin.Default()
	app.GET("/api/down/v1", func(c *gin.Context) {
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
		if !reject(successRate) {
			c.String(http.StatusInternalServerError, "reject from downstream")
			return
		}
		c.String(http.StatusOK, "success from downstream")
	})
	return app
}

func reject(successRate float64) bool {
	return rand.Float64() < successRate
}
