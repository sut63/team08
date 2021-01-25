package controllers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sut63/team08/ent"
	"github.com/sut63/team08/ent/doctor"
	"github.com/sut63/team08/ent/drug"
	"github.com/sut63/team08/ent/nurse"
	"github.com/sut63/team08/ent/patient"
	"github.com/sut63/team08/ent/prescription"
)

// PrescriptionController defines the struct for the pre-scription controller
type PrescriptionController struct {
	client *ent.Client
	router gin.IRouter
}

//Prescription struct
type Prescription struct {
	Doctor  int
	Patient int
	Nurse   int
	Drug    int
	Note    string
	Added   string
	Number  string
	Issue   string
}

// CreatePrescription handles POST requests for adding Prescription entities
// @Summary Create pre-scription
// @Description Create pre-scription
// @ID create-pre-scription
// @Accept   json
// @Produce  json
// @Param prescription body ent.Prescription true "Prescription entity"
// @Success 200 {object} ent.Prescription
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /prescriptions [post]
func (ctl *PrescriptionController) CreatePrescription(c *gin.Context) {
	obj := Prescription{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Prescription binding failed",
		})
		return
	}

	d, err := ctl.client.Doctor.
		Query().
		Where(doctor.IDEQ(int(obj.Doctor))).
		Only(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Doctor not found",
		})
		return
	}

	p, err := ctl.client.Patient.
		Query().
		Where(patient.IDEQ(int(obj.Patient))).
		Only(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Patient not found",
		})
		return
	}

	n, err := ctl.client.Nurse.
		Query().
		Where(nurse.IDEQ(int(obj.Nurse))).
		Only(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Nurse not found",
		})
		return
	}

	dr, err := ctl.client.Drug.
		Query().
		Where(drug.IDEQ(int(obj.Drug))).
		Only(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Drug not found",
		})
		return
	}

	time, err := time.Parse(time.RFC3339, obj.Added)
	if err != nil {
		c.JSON(400, gin.H{
			"status": false,
			"error":  "added time wrong",
		})
		return
	}

	da, err := ctl.client.Prescription.
		Create().
		SetDoctor(d).
		SetPatient(p).
		SetNurse(n).
		SetDrug(dr).
		SetPrescripNumber(obj.Number).
		SetPrescripDateTime(time).
		SetPrescripNote(obj.Note).
		SetPrescripIssue(obj.Issue).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"status": false,
			"error":  err,
		})
		return
	}
	c.JSON(200, gin.H{
		"status": true,
		"data":   da,
	})
}

// GetPrescription handles GET requests to retrieve a Prescription entity
// @Summary Get a Prescription entity by Patient PatientName
// @Description get Prescription by Patient PatientName
// @ID get-pre-scription
// @Produce  json
// @Param name path string true "Patient PatientName"
// @Success 200 {array} ent.Prescription
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /prescriptions/{name} [get]
func (ctl *PrescriptionController) GetPrescription(c *gin.Context) {
	PatientName := string(c.Param("name"))

	da, err := ctl.client.Prescription.
		Query().
		Where(prescription.HasPatientWith(patient.PatientNameEQ(string(PatientName)))).
		WithPatient().
		WithDoctor().
		WithNurse().
		WithDrug().
		All(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error":  err.Error(),
			"status": false,
		})
		return
	}

	if len(da) != 0 {
		c.JSON(200, da)
		return
	} else {
		c.JSON(404, gin.H{
			"error":  "patient not found",
			"status": false,
		})
		return
	}

}

// ListPrescription handles request to get a list of Prescription entities
// @Summary List Prescription entities
// @Description list Prescription entities
// @ID list-pre-scription
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Prescription
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /prescriptions [get]
func (ctl *PrescriptionController) ListPrescription(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 20
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

	Prescriptions, err := ctl.client.Prescription.
		Query().
		WithPatient().
		WithNurse().
		WithDrug().
		WithDoctor().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, Prescriptions)
}

// DeletePrescription handles DELETE requests to delete a Prescription entity
// @Summary Delete a Prescription entity by ID
// @Description get Prescription by ID
// @ID delete-pre-scription
// @Produce  json
// @Param id path int true "Prescription ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /prescriptions/{id} [delete]
func (ctl *PrescriptionController) DeletePrescription(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Prescription.
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

// UpdatePrescription handles PUT requests to update a Prescription entity
// @Summary Update a Prescription entity by ID
// @Description update Prescription by ID
// @ID update-pre-scription
// @Accept   json
// @Produce  json
// @Param id path int true "Prescription ID"
// @Param prescription body ent.Prescription true "Prescription entity"
// @Success 200 {object} ent.Prescription
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /prescriptions/{id} [put]
func (ctl *PrescriptionController) UpdatePrescription(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Prescription{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "prescription binding failed",
		})
		return
	}
	obj.ID = int(id)
	u, err := ctl.client.Prescription.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, u)
}

// NewPrescriptionController creates and registers handles for the Prescription controller
func NewPrescriptionController(router gin.IRouter, client *ent.Client) *PrescriptionController {
	da := &PrescriptionController{
		client: client,
		router: router,
	}
	da.register()
	return da
}

// InitPrescriptionController registers routes to the main engine
func (ctl *PrescriptionController) register() {
	prescriptions := ctl.router.Group("/prescriptions")

	prescriptions.GET("", ctl.ListPrescription)

	// CRUD
	prescriptions.POST("", ctl.CreatePrescription)
	prescriptions.GET(":name", ctl.GetPrescription)
	prescriptions.PUT(":id", ctl.UpdatePrescription)
	prescriptions.DELETE(":id", ctl.DeletePrescription)
}
