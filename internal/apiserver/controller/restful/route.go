package restful

import (
	"github.com/strugglehonor/KCS/internal/apiserver/controller/restful/cluster"
	"github.com/strugglehonor/KCS/internal/apiserver/controller/restful/deployment"
	"github.com/strugglehonor/KCS/pkg/log"

	// "github.com/strugglehonor/KCS/internal/apiserver/store"
	"github.com/strugglehonor/KCS/internal/apiserver/store/mysql"

	"github.com/gin-gonic/gin"
)

func InitRouter(g *gin.Engine) {
	log.INFO.Println("init gin router")
	installMiddleWare(g)
	installController(g)
}

func installMiddleWare(g *gin.Engine) {
	log.INFO.Println("install gin middleware")
}

func installController(g *gin.Engine) *gin.Engine {
	log.INFO.Println("install gin controller")
	v1 := g.Group("/v1") 
	{
		// clusterController := cluster.newClusterController()
		storeIns, _ := mysql.GetMySQLFactoryOr()
		
		clusterv1 := v1.Group("/cluster")
		{
			clusterController := cluster.NewClusterController(storeIns)

			clusterv1.POST("", clusterController.Create)
		}

		deploymentv1 := v1.Group("/deployment")
		{
			deploymentController := deployment.NewDeploymentController(storeIns)

			deploymentv1.POST("", deploymentController.Create)
		}
	}

	return g
}

