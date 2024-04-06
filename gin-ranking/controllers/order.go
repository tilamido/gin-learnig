package controllers

import "github.com/gin-gonic/gin"

type OrderController struct{}

func (o OrderController) GetInfo(c *gin.Context) {
	ReturnSucess(c, 0, "订单信息", "已下单", 1)
}

type Body struct {
	Cid  int    `json:"cid"`
	Name string `json:"name"`
}

func (o OrderController) GetList(c *gin.Context) {
	//post from形式
	// cid := c.PostForm("cid")
	// name := c.DefaultPostForm("name", "wangwu")

	//post json 格式
	// param := make(map[string]interface{})
	// if err := c.BindJSON(&param); err == nil {
	// 	ReturnSucess(c, 0, param[cid], param[name], 1)
	// } else {
	// 	ReturnError(c, 401, gin.H{"err": err})
	// }

	//结构体模式
	param := &Body{}

	if err := c.BindJSON(&param); err == nil {
		ReturnSucess(c, 0, param.Cid, param.Name, 1)
	} else {
		ReturnError(c, 401, gin.H{"err": err})
	}
}
