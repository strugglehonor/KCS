package pod

import (
	"net/http"

	"github.com/marmotedu/errors"
	log "github.com/sirupsen/logrus"
	"github.com/strugglehonor/KCS/internal/model"
	"github.com/strugglehonor/KCS/pkg/code"

	"github.com/gin-gonic/gin"
)



func (PodController *PodController) Create(c *gin.Context) {
	log.Info("create pod function called")

	var request model.Pod
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, errors.WithCode(code.ErrBind, err.Error(), nil))
		return 
	}

	err := PodController.srv.Pod().Create(c,  &request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.WithCode(code.ErrDatabase, err.Error(), nil))
	}
}