package routes

import (
	"github.com/bhaveshkh/ecommerce-yt/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/signup", controllers.SignUp())
	incomingRoutes.POST("/users/login", controllers.Login())
	incomingRoutes.POST("/admin/addproduct", controllers.ProductViewerAdmin())
	incomingRoutes.POST("/users/productview", controllers.SearchProduct())
	incomingRoutes.POST("/users/search", controllers.SearchProductByQuery())
}
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