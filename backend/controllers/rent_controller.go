package controllers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/sut63/team08/ent"
	"github.com/sut63/team08/ent/nurse"
	"github.com/sut63/team08/ent/patient"
	"github.com/sut63/team08/ent/rent"
	"github.com/sut63/team08/ent/room"
	"github.com/gin-gonic/gin"
)

// RentController defines the struct for the rent controller
type RentController struct {
	client *ent.Client
	router gin.IRouter
}

//Rent struct
type Rent struct {
	Added   string
	Room    int
	Patient int
	Nurse   int
}

// CreateRent handles POST requests for adding rent entities
// @Summary Create rent
// @Description Create rent
// @ID create-rent
// @Accept   json
// @Produce  json
// @Param rent body Rent true "Rent entity"
// @Success 200 {object} ent.Rent
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /rents [post]
func (ctl *RentController) CreateRent(c *gin.Context) {
	obj := Rent{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "rent binding failed",
		})
		return
	}
	ro, err := ctl.client.Room.
		Query().
		Where(room.IDEQ(int(obj.Room))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "room not found",
		})
		return
	}

	p, err := ctl.client.Patient.
		Query().
		Where(patient.IDEQ(int(obj.Patient))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "patient not found",
		})
		return
	}

	n, err := ctl.client.Nurse.
		Query().
		Where(nurse.IDEQ(int(obj.Nurse))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "nurse not found",
		})
		return
	}


	datetime, err := time.Parse(time.RFC3339, obj.Added)

	r, err := ctl.client.Rent.
		Create().
		SetAddedTime(datetime).
		SetRoom(ro).
		SetPatient(p).
		SetNurse(n).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, gin.H{
		"status": true,
		"data":   r,
	})
}

// GetRent handles GET requests to retrieve a rent entity
// @Summary Get a rent entity by ID
// @Description get rent by ID
// @ID get-rent
// @Produce  json
// @Param id path int true "Rent ID"
// @Success 200 {object} ent.Rent
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /rents/{id} [get]
func (ctl *RentController) GetRent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	b, err := ctl.client.Rent.
		Query().
		Where(rent.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, b)
}

// ListRent handles request to get a list of rent entities
// @Summary List rent entities
// @Description list rent entities
// @ID list-rent
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Rent
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /rents [get]
func (ctl *RentController) ListRent(c *gin.Context) {
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

	rents, err := ctl.client.Rent.
		Query().
		WithNurse().
		WithPatient().
		WithRoom().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, rents)
}

// DeleteRent handles DELETE requests to delete a rent entity
// @Summary Delete a rent entity by ID
// @Description get rent by ID
// @ID delete-rent
// @Produce  json
// @Param id path int true "Rent ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /rents/{id} [delete]
func (ctl *RentController) DeleteRent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Rent.
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

// UpdateRent handles PUT requests to update a Rent entity
// @Summary Update a rent entity by ID
// @Description update rent by ID
// @ID update-rent
// @Accept   json
// @Produce  json
// @Param id path int true "Rent ID"
// @Param rent body ent.Rent true "Rent entity"
// @Success 200 {object} ent.Rent
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /rents/{id} [put]
func (ctl *RentController) UpdateRent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Rent{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "rent binding failed",
		})
		return
	}
	obj.ID = int(id)
	b, err := ctl.client.Rent.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, b)
}

// NewRentController creates and registers handles for the rent controller
func NewRentController(router gin.IRouter, client *ent.Client) *RentController {
	uc := &RentController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

// RentController registers routes to the main engine
func (ctl *RentController) register() {
	rents := ctl.router.Group("/rents")

	rents.GET("", ctl.ListRent)

	// CRUD
	rents.POST("", ctl.CreateRent)
	rents.GET(":id", ctl.GetRent)
	rents.PUT(":id", ctl.UpdateRent)
	rents.DELETE(":id", ctl.DeleteRent)
}
