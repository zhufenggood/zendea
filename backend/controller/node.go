package controller

import (
	"github.com/gin-gonic/gin"

	"zendea/builder"
	"zendea/service"
	"zendea/form"
)

type NodeController struct {
	BaseController
}

// List 节点列表
func (c *NodeController) List(ctx *gin.Context) {
	nodes := service.NodeService.GetNodes()
	c.Success(ctx, builder.BuildNodes(nodes))
}

// Show 显示单个节点
func (c *NodeController) Show(ctx *gin.Context) {
	var gDto form.GeneralGetDto
	if c.BindAndValidate(ctx, &gDto) {
		node := service.NodeService.Get(gDto.ID)
		c.Success(ctx, builder.BuildNode(node))
	}
}
