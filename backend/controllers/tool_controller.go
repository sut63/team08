package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/sut63/team08/ent"
	"github.com/sut63/team08/ent/tool"
	"github.com/gin-gonic/gin"
)

// ToolController defines the struct for the tool controller
type ToolController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateTool handles POST requests for adding tool entities
// @Summary Create tool
// @Description Create tool
// @ID create-tool
// @Accept   json
// @Produce  json
// @Param tool body ent.Tool true "Tool entity"
// @Success 200 {object} ent.Tool
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /tools [post]
func (ctl *ToolController) CreateTool(c *gin.Context) {
	obj := ent.Tool{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Tool binding failed",
		})
		return
	}

	f, err := ctl.client.Tool.
		Create().
		SetToolName(obj.ToolName).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, f)
}

// GetTool handles GET requests to retrieve a tool entity
// @Summary Get a tool entity by ID
// @Description get tool by ID
// @ID get-tool
// @Produce  json
// @Param id path int true "Tool ID"
// @Success 200 {object} ent.Tool
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /tools/{id} [get]
func (ctl *ToolController) GetTool(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	f, err := ctl.client.Tool.
		Query().
		Where(tool.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, f)
}

// ListTool handles request to get a list of tool entities
// @Summary List tool entities
// @Description list tool entities
// @ID list-tool
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Tool
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /tools [get]
func (ctl *ToolController) ListTool(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 10
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 10, 64)
		if err == nil {
			limit = int(limit64)
		}
	}

	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {
			offset = int(offset64)
		}
	}

	tools, err := ctl.client.Tool.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, tools)
}

// DeleteTool handles DELETE requests to delete a tool entity
// @Summary Delete a tool entity by ID
// @Description get tool by ID
// @ID delete-tool
// @Produce  json
// @Param id path int true "tool ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /tools/{id} [delete]
func (ctl *ToolController) DeleteTool(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Tool.
		DeleteOneID(int(id)).
		Exec(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"result": fmt.Sprintf("ok deleted %v", id)})
}

// UpdateTool handles PUT requests to update a Tool entity
// @Summary Update a tool entity by ID
// @Description update tool by ID
// @ID update-tool
// @Accept   json
// @Produce  json
// @Param id path int true "tool ID"
// @Param tool body ent.Tool true "Tool entity"
// @Success 200 {object} ent.Tool
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /tools/{id} [put]
func (ctl *ToolController) UpdateTool(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Tool{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Tool binding failed",
		})
		return
	}
	obj.ID = int(id)
	fmt.Println(obj.ID)
	f, err := ctl.client.Tool.
		UpdateOne(&obj).
		SetToolName(obj.ToolName).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, f)
}

// NewToolController creates and registers handles for the fund controller
func NewToolController(router gin.IRouter, client *ent.Client) *ToolController {
	uc := &ToolController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

// InitToolController registers routes to the main engine
func (ctl *ToolController) register() {
	tools := ctl.router.Group("/tools")

	tools.GET("", ctl.ListTool)
	// CRUD
	tools.POST("", ctl.CreateTool)
	tools.GET(":id", ctl.GetTool)
	tools.PUT(":id", ctl.UpdateTool)
	tools.DELETE(":id", ctl.DeleteTool)
}
