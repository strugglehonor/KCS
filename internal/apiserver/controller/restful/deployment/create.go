package deployment

import (
	"net/http"

	"github.com/strugglehonor/KCS/internal/model"
	"github.com/strugglehonor/KCS/pkg/code"

	"github.com/gin-gonic/gin"
	"github.com/marmotedu/errors"
	log "github.com/sirupsen/logrus"
)

func (deploymentController *DeploymentController) Create(c *gin.Context) {
	log.Info("create cluster function called")
	
	// request ?
	var request model.Deployment
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, errors.WithCode(code.ErrBind, err.Error(), nil))
		return 
	}

	err := deploymentController.srv.Deployment().Create(c, &request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.WithCode(code.ErrDatabase, err.Error(), nil))
	}
}