package v1

import "github.com/gin-gonic/gin"

// @Summary demo，测试get
// @Produce  json
// @Param param1 query string false "some params named param1"
// @Param param2 query string false "some params named param2"
// @Success 200 {string} json ""
// @Router /api/v1/demo/get [get]
func DemoGet(c *gin.Context) {
	param1 := c.Query("param1")
	param2 := c.Query("param2")
	c.JSON(200, gin.H{
		"method": "get",
		"param1": param1,
		"param2": param2,
	})
}

type Params struct {
	Param1 string `json:"param1"`
	Param2 string `json:"param2"`
}

// @Summary demo，测试post
// @Produce  json
// @Param params body Params false "some params json"
// @Success 200 {string} json ""
// @Router /api/v1/demo/post [post]
func DemoPost(c *gin.Context) {
	params := Params{}
	c.BindJSON(&params)
	c.JSON(200, gin.H{
		"method":    "post",
		"post_body": params,
	})
}
