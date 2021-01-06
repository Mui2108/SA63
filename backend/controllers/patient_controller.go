package controllers

import (
	"context"
	"fmt"
	"strconv"
	"time"
	
	"github.com/gin-gonic/gin"
	"github.com/panupong/app/ent"
	"github.com/panupong/app/ent/patient"
	"github.com/panupong/app/ent/title"
	"github.com/panupong/app/ent/gender"
	"github.com/panupong/app/ent/job"
	
)

// PatientController defines the struct for the patient controller
type PatientController struct {
	client *ent.Client
	router gin.IRouter
}

type Patient struct {
	Card_id string
	First_name string
	Last_name string
	Allergic string
	Address string
	Birthday string
	Age int
	Gender int
	Job int
	Title int
}
// CreatePatient handles POST requests for adding patient entities
// @Summary Create patient
// @Description Create patient
// @ID create-patient
// @Accept   json
// @Produce  json
// @Param patient body ent.Patient true "Patient entity"
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

	g, err := ctl.client.Gender.
		Query().
		Where(gender.IDEQ(int(obj.Gender))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "หา Gender ไม่เจอ",
		})
			return 
	}
	
	t, err := ctl.client.Title.
		Query().
		Where(title.IDEQ(int(obj.Title))).
		Only(context.Background())
		if err != nil {
			c.JSON(400, gin.H{
				"error": "title not found",
			})
			return 
		}


	j, err := ctl.client.Job.
		Query().
		Where(job.IDEQ(int(obj.Job))).
		Only(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "job not found",
		})
		return
	}
	ti, err := time.Parse(time.RFC3339, obj.Birthday+"T00:00:00Z")
	// ti := ti.String()[:9] + "T00:00:00Z"
	
	p, err := ctl.client.Patient.
		Create().
		SetCardID(obj.Card_id).
		SetFirstName(obj.First_name).
		SetLastName(obj.Last_name).
		SetAge(int(obj.Age)).
		SetAllergic(obj.Allergic).
		SetAddress(obj.Address).
		SetBirthday(ti).
		SetGender(g).
		SetTitle(t).
		SetJob(j).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, gin.H{
		"status": true,
		"data": p,
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

	p, err := ctl.client.Patient.
		Query().
		Where(patient.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, p)
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

	patients, err := ctl.client.Patient.
		Query().
		WithTitle().
		WithJob().
		WithGender().
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

// UpdatePatient handles PUT requests to update a patient entity
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
			"error": "patient binding failed",
		})
		return
	}
	obj.ID = int(id)
	u, err := ctl.client.Patient.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, u)
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

// InitPatientController registers routes to the main engine
func (ctl *PatientController) register() {
	patients := ctl.router.Group("/patients")

	patients.GET("", ctl.ListPatient)

	// CRUD
	patients.POST("", ctl.CreatePatient)
	patients.GET(":id", ctl.GetPatient)
	patients.PUT(":id", ctl.UpdatePatient)
	patients.DELETE(":id", ctl.DeletePatient)
}
