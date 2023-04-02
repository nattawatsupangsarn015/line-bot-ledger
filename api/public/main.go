package public

import (
	"example/line-bot-ledger/request"
	"example/line-bot-ledger/utils"

	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine) {
	public := route.Group("/public")
	public.POST("/register", func(c *gin.Context) {
		var user request.Register
		if err := c.BindJSON(&user); err != nil {
			utils.HandleBadRequest(c, err)
			return
		}

		result, err := CreateUser(user)
		utils.HandleResponse(c, result, err, 201, 500)
		return
	})

	// public.POST("/login", func(c *gin.Context) {
	// 	var user request.Login
	// 	if err := c.BindJSON(&user); err != nil {
	// 		utils.HandleBadRequest(c, err)
	// 		return
	// 	}

	// 	result, err := CreateUser(user)
	// 	utils.HandleResponse(c, result, err, 201, 500)
	// 	return
	// })
}
