package controllers

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sut63/team08/ent"
	"github.com/sut63/team08/ent/nurse"
)

// NurseController defines the struct for the nurse controller
type NurseController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateNurse handles POST requests for adding nurse entities
// @Summary Create nurse
// @Description Create nurse
// @ID create-nurse
// @Accept   json
// @Produce  json
// @Param nurse body ent.Nurse true "Nurse entity"
// @Success 200 {object} ent.Nurse
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /nurses [post]
func (ctl *NurseController) CreateNurse(c *gin.Context) {
	obj := ent.Nurse{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "nurse binding failed",
		})
		return
	}

	r, err := ctl.client.Nurse.
		Create().
		SetNurseName(obj.NurseName).
		SetNurseEmail(obj.NurseEmail).
		SetNursePassword(obj.NursePassword).
		SetNurseTel(obj.NurseTel).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, r)
}

// GetNurse handles GET requests to retrieve a nurse entity
// @Summary Get a nurse entity by ID
// @Description get nurse by ID
// @ID get-nurse
// @Produce  json
// @Param id path int true "Nurse ID"
// @Success 200 {object} ent.Nurse
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /nurses/{id} [get]
func (ctl *NurseController) GetNurse(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	r, err := ctl.client.Nurse.
		Query().
		Where(nurse.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, r)
}

// ListNurse handles request to get a list of nurse entities
// @Summary List nurse entities
// @Description list nurse entities
// @ID list-nurse
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Nurse
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /nurses [get]
func (ctl *NurseController) ListNurse(c *gin.Context) {
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

	nurses, err := ctl.client.Nurse.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, nurses)
}

// NewNurseController creates and registers handles for the nurse controller
func NewNurseController(router gin.IRouter, client *ent.Client) *NurseController {
	rc := &NurseController{
		client: client,
		router: router,
	}

	rc.register()

	return rc

}

func (ctl *NurseController) register() {
	nurses := ctl.router.Group("/nurses")
	nurses.GET("", ctl.ListNurse)

	nurses.POST("", ctl.CreateNurse)
	nurses.GET(":id", ctl.GetNurse)

}
