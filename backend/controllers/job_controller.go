package controllers

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/panupong/app/ent"
	"github.com/panupong/app/ent/job"
)

// JobController defines the struct for the job controller
type JobController struct {
	client *ent.Client
	router gin.IRouter
}

// GetJob handles GET requests to retrieve a Job entity
// @Summary Get a job entity by ID
// @Description get job by ID
// @ID get-job
// @Produce  json
// @Param id path int true "job ID"
// @Success 200 {object} ent.Job
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /jobs/{id} [get]
func (ctl *JobController) GetJob(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	jo, err := ctl.client.Job.
		Query().
		Where(job.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, jo)
}

// ListJob handles request to get a list of job entities
// @Summary List job entities
// @Description list job entities
// @ID list-job
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Job
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /jobs [get]
func (ctl *JobController) ListJob(c *gin.Context) {
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

	jobs, err := ctl.client.Job.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, jobs)
}

// NewJobController creates and registers handles for the job controller
func NewJobController(router gin.IRouter, client *ent.Client) *JobController {
	uc := &JobController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

// InitJobController registers routes to the main engine
func (ctl *JobController) register() {
	
	jobs := ctl.router.Group("/jobs")
	jobs.GET("", ctl.ListJob)
	// CRUD
	jobs.GET(":id", ctl.GetJob)

}