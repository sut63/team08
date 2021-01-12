package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/sut63/team08/ent"
	"github.com/sut63/team08/ent/bloodtype"
	"github.com/sut63/team08/ent/gender"
	"github.com/sut63/team08/ent/patient"
	"github.com/sut63/team08/ent/prefix"
	"github.com/gin-gonic/gin"
)

// PatientController defines the struct for the patient controller
type PatientController struct {
	client *ent.Client
	router gin.IRouter
}

//Patient struct
type Patient struct {
	Name      string
	Weight    string
	Height    string
	Age       string
	Bloodtype int
	Gender    int
	Prefix    int
}

// CreatePatient handles POST requests for adding patient entities
// @Summary Create patient
// @Description Create patient
// @ID create-patient
// @Accept   json
// @Produce  json
// @Param patient body Patient true "Patient entity"
// @Success 200 {object} ent.Patient
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /patients [post]
func (ctl *PatientController) CreatePatient(c *gin.Context) {
	obj := Patient{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "patient binding failed",
		})
		return
	}
	pf, err := ctl.client.Prefix.
		Query().
		Where(prefix.IDEQ(int(obj.Prefix))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "prefix not found",
		})
		return
	}

	g, err := ctl.client.Gender.
		Query().
		Where(gender.IDEQ(int(obj.Gender))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "gender not found",
		})
		return
	}
	bt, err := ctl.client.Bloodtype.
		Query().
		Where(bloodtype.IDEQ(int(obj.Bloodtype))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "bloodtype not found",
		})
		return
	}

	w, err := strconv.ParseFloat(obj.Weight, 64)
	h, err := strconv.ParseFloat(obj.Height, 64)
	a, err := strconv.Atoi(obj.Age)

	p, err := ctl.client.Patient.
		Create().
		SetPrefix(pf).
		SetPatientName(obj.Name).
		SetGender(g).
		SetPatientWeight(w).
		SetPatientHeight(h).
		SetPatientAge(a).
		SetBloodtype(bt).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, gin.H{
		"status": true,
		"data":   p,
	})
}

// GetPatient handles GET requests to retrieve a patient entity
// @Summary Get a patient entity by ID
// @Description get patient by ID
// @ID get-patient
// @Produce  json
// @Param id path int true "Patient ID"
// @Success 200 {object} ent.Patient
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /patients/{id} [get]
func (ctl *PatientController) GetPatient(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	b, err := ctl.client.Patient.
		Query().
		Where(patient.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, b)
}

// ListPatient handles request to get a list of patient entities
// @Summary List patient entities
// @Description list patient entities
// @ID list-patient
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Patient
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /patients [get]
func (ctl *PatientController) ListPatient(c *gin.Context) {
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

	patients, err := ctl.client.Patient.
		Query().
		WithBloodtype().
		WithGender().
		WithPrefix().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, patients)
}

// DeletePatient handles DELETE requests to delete a patient entity
// @Summary Delete a patient entity by ID
// @Description get patient by ID
// @ID delete-patient
// @Produce  json
// @Param id path int true "Patient ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /patients/{id} [delete]
func (ctl *PatientController) DeletePatient(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Patient.
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

// UpdatePatient handles PUT requests to update a Patient entity
// @Summary Update a patient entity by ID
// @Description update patient by ID
// @ID update-patient
// @Accept   json
// @Produce  json
// @Param id path int true "Patient ID"
// @Param patient body ent.Patient true "Patient entity"
// @Success 200 {object} ent.Patient
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /patients/{id} [put]
func (ctl *PatientController) UpdatePatient(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Patient{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "booking binding failed",
		})
		return
	}
	obj.ID = int(id)
	b, err := ctl.client.Patient.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, b)
}

// NewPatientController creates and registers handles for the patient controller
func NewPatientController(router gin.IRouter, client *ent.Client) *PatientController {
	uc := &PatientController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

// PatientController registers routes to the main engine
func (ctl *PatientController) register() {
	patients := ctl.router.Group("/patients")

	patients.GET("", ctl.ListPatient)

	// CRUD
	patients.POST("", ctl.CreatePatient)
	patients.GET(":id", ctl.GetPatient)
	patients.PUT(":id", ctl.UpdatePatient)
	patients.DELETE(":id", ctl.DeletePatient)
}
