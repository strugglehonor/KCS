package apiserver

import (
	"github.com/gin-gonic/gin"
)

func initRouter(g *gin.Engine) {
	installMiddleWare(g)
	installController(g)
}

func installMiddleWare(g *gin.Engine) {

}

func installController(g *gin.Engine) {

}

