package controllers

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sut63/team08/ent"
	"github.com/sut63/team08/ent/schemetype"
)

// SchemeTypeController defines the struct for the schemeType controller
type SchemeTypeController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateSchemeType handles POST requests for adding schemeType entities
// @Summary Create schemeType
// @Description Create schemeType
// @ID create-schemeType
// @Accept   json
// @Produce  json
// @Param schemeType body ent.SchemeType true "SchemeType entity"
// @Success 200 {object} ent.SchemeType
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /schemeTypes [post]
func (ctl *SchemeTypeController) CreateSchemeType(c *gin.Context) {
	obj := ent.SchemeType{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "SchemeType binding failed",
		})
		return
	}

	me, err := ctl.client.SchemeType.
		Create().
		SetSchemeTypeName(obj.SchemeTypeName).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, me)
}

// GetSchemeType handles GET requests to retrieve a schemeType entity
// @Summary Get a schemeType entity by ID
// @Description get schemeType by ID
// @ID get-schemeType
// @Produce  json
// @Param id path int true "SchemeType ID"
// @Success 200 {object} ent.SchemeType
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /schemeTypes/{id} [get]
func (ctl *SchemeTypeController) GetSchemeType(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	me, err := ctl.client.SchemeType.
		Query().
		Where(schemetype.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, me)
}

// ListSchemeType handles request to get a list of schemeType entities
// @Summary List schemeType entities
// @Description list schemeType entities
// @ID list-schemeType
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.SchemeType
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /schemeTypes [get]
func (ctl *SchemeTypeController) ListSchemeType(c *gin.Context) {
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

	schemeTypes, err := ctl.client.SchemeType.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, schemeTypes)
}

// NewSchemeTypeController creates and registers handles for the schemeType controller
func NewSchemeTypeController(router gin.IRouter, client *ent.Client) *SchemeTypeController {
	uc := &SchemeTypeController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

// InitSchemeTypeeController registers routes to the main engine
func (ctl *SchemeTypeController) register() {
	schemeTypes := ctl.router.Group("/schemeTypes")

	schemeTypes.GET("", ctl.ListSchemeType)

	// CRUD
	schemeTypes.POST("", ctl.CreateSchemeType)
	schemeTypes.GET(":id", ctl.GetSchemeType)
}
