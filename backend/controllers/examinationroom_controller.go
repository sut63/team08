package controllers
 
import (
   "context"
   "fmt"
   "strconv"
   "github.com/sut63/team08/ent"
   "github.com/sut63/team08/ent/examinationroom"
   "github.com/gin-gonic/gin"
)
 
// ExaminationroomController defines the struct for the examinationroom controller
type ExaminationroomController struct {
   client *ent.Client
   router gin.IRouter
}
// CreateExaminationroom handles POST requests for adding examinationroom entities
// @Summary Create examinationroom
// @Description Create examinationroom
// @ID create-examinationroom
// @Accept   json
// @Produce  json
// @Param examinationroom body ent.Examinationroom true "Examinationroom entity"
// @Success 200 {object} ent.Examinationroom
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /examinationrooms [post]
func (ctl *ExaminationroomController) CreateExaminationroom(c *gin.Context) {
	obj := ent.Examinationroom{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Examinationroom binding failed",
		})
		return
	}
  
	f, err := ctl.client.Examinationroom.
		Create().
		SetExaminationroomName(obj.ExaminationroomName).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}
  
	c.JSON(200, f)
 }
 // GetExaminationroom handles GET requests to retrieve a examinationroomund entity
// @Summary Get a examinationroom entity by ID
// @Description get examinationroom by ID
// @ID get-examinationroom
// @Produce  json
// @Param id path int true "Examinationroom ID"
// @Success 200 {object} ent.Examinationroom
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /examinationrooms/{id} [get]
func (ctl *ExaminationroomController) GetExaminationroom(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	f, err := ctl.client.Examinationroom.
		Query().
		Where(examinationroom.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	c.JSON(200, f)
 }
 // ListExaminationroom handles request to get a list of examinationroom entities
// @Summary List examinationroom entities
// @Description list examinationroom entities
// @ID list-examinationroom
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Examinationroom
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /examinationrooms [get]
func (ctl *ExaminationroomController) ListExaminationroom(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 10
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 10, 64)
		if err == nil {limit = int(limit64)}
	}
  
	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {offset = int(offset64)}
	}
  
	examinationrooms, err := ctl.client.Examinationroom.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
		if err != nil {
		c.JSON(400, gin.H{"error": err.Error(),})
		return
	}
  
	c.JSON(200, examinationrooms)
 }
 // DeleteExaminationroom handles DELETE requests to delete a examinationroom entity
// @Summary Delete a examinationroom entity by ID
// @Description get examinationroom by ID
// @ID delete-examinationroom
// @Produce  json
// @Param id path int true "examinationroom ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /examinationrooms/{id} [delete]
func (ctl *ExaminationroomController) DeleteExaminationroom(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	err = ctl.client.Examinationroom.
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
 // UpdateExaminationroom handles PUT requests to update a examinationroom entity
// @Summary Update a examinationroom entity by ID
// @Description update examinationroom by ID
// @ID update-examinationroom
// @Accept   json
// @Produce  json
// @Param id path int true "examinationroom ID"
// @Param examinationroom body ent.Examinationroom true "examinationroom entity"
// @Success 200 {object} ent.Examinationroom
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /examinationrooms/{id} [put]
func (ctl *ExaminationroomController) UpdateExaminationroom(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	obj := ent.Examinationroom{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Examinationroom binding failed",
		})
		return
	}
	obj.ID = int(id)
	fmt.Println(obj.ID)
	f, err := ctl.client.Examinationroom.
		UpdateOne(&obj).
		SetExaminationroomName(obj.ExaminationroomName).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed",})
		return
	}
  
	c.JSON(200, f)
 }
 // NewExaminationroomController creates and registers handles for the fund controller
func NewExaminationroomController(router gin.IRouter, client *ent.Client) *ExaminationroomController {
	uc := &ExaminationroomController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
 }
  
 // InitExaminationroomController registers routes to the main engine
 func (ctl *ExaminationroomController) register() {
	examinationrooms := ctl.router.Group("/examinationrooms")
  
	examinationrooms.GET("", ctl.ListExaminationroom)
	// CRUD
	examinationrooms.POST("", ctl.CreateExaminationroom)
	examinationrooms.GET(":id", ctl.GetExaminationroom)
	examinationrooms.PUT(":id", ctl.UpdateExaminationroom)
	examinationrooms.DELETE(":id", ctl.DeleteExaminationroom)
 }
 
 