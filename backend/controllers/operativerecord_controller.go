package controllers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/sut63/team08/ent"
	"github.com/sut63/team08/ent/nurse"
	"github.com/sut63/team08/ent/operative"
	"github.com/sut63/team08/ent/operativerecord"
	"github.com/sut63/team08/ent/tool"
	"github.com/sut63/team08/ent/examinationroom"
	"github.com/gin-gonic/gin"
)

// OperativerecordController defines the struct for the Operativerecord controller
type OperativerecordController struct {
	client *ent.Client
	router gin.IRouter
}

//Operativerecord struct
type Operativerecord struct {
	Nurse           int
	Examinationroom int
	Tool            int
	Operative       int
	Added           string
}

// CreateOperativerecord handles POST requests for adding Operativerecord entities
// @Summary Create operativerecord
// @Description Create operativerecord
// @ID create-operativerecord
// @Accept   json
// @Produce  json
// @Param operativerecord body Operativerecord true "Operativerecord entity"
// @Success 200 {object} ent.Operativerecord
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /operativerecords [post]
func (ctl *OperativerecordController) CreateOperativerecord(c *gin.Context) {
	obj := Operativerecord{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Operativerecord binding failed",
		})
		return
	}
	p, err := ctl.client.Nurse.
		Query().
		Where(nurse.IDEQ(int(obj.Nurse))).
		Only(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Nurse not found",
		})
		return
	}
	st, err := ctl.client.Tool.
		Query().
		Where(tool.IDEQ(int(obj.Tool))).
		Only(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Tool not found",
		})
		return
	}
	f, err := ctl.client.Examinationroom.
		Query().
		Where(examinationroom.IDEQ(int(obj.Examinationroom))).
		Only(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Examinationroom not found",
		})
		return
	}
	ce, err := ctl.client.Operative.
		Query().
		Where(operative.IDEQ(int(obj.Operative))).
		Only(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Operative not found",
		})
		return
	}
	time, err := time.Parse(time.RFC3339, obj.Added)
	cp, err := ctl.client.Operativerecord.
		Create().
		SetNurse(p).
		SetTool(st).
		SetExaminationroom(f).
		SetOperative(ce).
		SetOperativeTime(time).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}
	c.JSON(200, gin.H{
		"status": true,
		"data":   cp,
	})
}

// GetOperativerecord handles GET requests to retrieve a Operativerecord entity
// @Summary Get a operativerecord entity by ID
// @Description get operativerecord by ID
// @ID get-operativerecord
// @Produce  json
// @Param id path int true "Operativerecord ID"
// @Success 200 {object} ent.Operativerecord
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /operativerecords/{id} [get]
func (ctl *OperativerecordController) GetOperativerecord(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	b, err := ctl.client.Operativerecord.
		Query().
		Where(operativerecord.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, b)
}

// ListOperativerecord handles request to get a list of Operativerecord entities
// @Summary List operativerecord entities
// @Description list operativerecord entities
// @ID list-operativerecord
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Operativerecord
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /operativerecords [get]
func (ctl *OperativerecordController) ListOperativerecord(c *gin.Context) {
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

	operativerecords, err := ctl.client.Operativerecord.
		Query().
		WithNurse().
		WithExaminationroom().
		WithTool().
		WithOperative().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, operativerecords)

}

// DeleteOperativerecord handles DELETE requests to delete a Operativerecord entity
// @Summary Delete a operativerecord entity by ID
// @Description get operativerecord by ID
// @ID delete-operativerecord
// @Produce  json
// @Param id path int true "Operativerecord ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /operativerecords/{id} [delete]
func (ctl *OperativerecordController) DeleteOperativerecord(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = ctl.client.Operativerecord.
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

// UpdateOperativerecordhandles PUT requests to update a Operativerecord entity
// @Summary Update a operativerecord entity by ID
// @Description update operativerecord by ID
// @ID update-operativerecord
// @Accept   json
// @Produce  json
// @Param id path int true "Operativerecord ID"
// @Param operativerecord body ent.Operativerecord true "Operativerecord entity"
// @Success 200 {object} ent.Operativerecord
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /operativerecords/{id} [put]
func (ctl *OperativerecordController) UpdateOperativerecord(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	obj := ent.Operativerecord{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Operativerecord binding failed",
		})
		return
	}
	obj.ID = int(id)
	b, err := ctl.client.Operativerecord.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}
	c.JSON(200, b)
}

// NewOperativerecordController creates and registers handles for the operativerecord controller
func NewOperativerecordController(router gin.IRouter, client *ent.Client) *OperativerecordController {
	da := &OperativerecordController{
		client: client,
		router: router,
	}

	da.register()

	return da

}

func (ctl *OperativerecordController) register() {
	operativerecords := ctl.router.Group("/operativerecords")

	operativerecords.POST("", ctl.CreateOperativerecord)
	operativerecords.GET("", ctl.ListOperativerecord)
	operativerecords.DELETE(":id", ctl.DeleteOperativerecord)

}
