package cluster

import (
	"net/http"

	"github.com/strugglehonor/KCS/internal/model"
	"github.com/strugglehonor/KCS/pkg/code"

	"github.com/gin-gonic/gin"
	"github.com/marmotedu/errors"
	log "github.com/sirupsen/logrus"
)

func (clusterController *ClusterController) Create(c *gin.Context) {
	log.Info("create cluster function called")
	
	// request ?
	var request model.Cluster
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.WithCode(code.ErrBind, err.Error(), nil).Error()})
		return 
	}

	err := clusterController.srv.Cluster().Create(c, &request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errors.WithCode(code.ErrDatabase, err.Error(), nil).Error()})
	}
}