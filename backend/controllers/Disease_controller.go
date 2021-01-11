package controllers
 
import (
   "context"
   "fmt"
   "strconv"
   "github.com/sut63/team08/ent"
   "github.com/sut63/team08/ent/disease"
   "github.com/gin-gonic/gin"
)
 
// DiseaseController defines the struct for the disease controller
type DiseaseController struct {
   client *ent.Client
   router gin.IRouter
}
// CreateDisease handles POST requests for adding disease entities
// @Summary Create disease
// @Description Create disease
// @ID create-disease
// @Accept   json
// @Produce  json
// @Param disease body ent.Disease true "Disease entity"
// @Success 200 {object} ent.Disease
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /diseases [post]
func (ctl *DiseaseController) CreateDisease(c *gin.Context) {
	obj := ent.Disease{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Disease binding failed",
		})
		return
	}
  
	di, err := ctl.client.Disease.
		Create().
		SetDiseaseName(obj.DiseaseName).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}
  
	c.JSON(200, di)
 }
 // GetDisease handles GET requests to retrieve a disease entity
// @Summary Get a disease entity by ID
// @Description get disease by ID
// @ID get-disease
// @Produce  json
// @Param id path int true "Disease ID"
// @Success 200 {object} ent.Disease
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /diseases/{id} [get]
func (ctl *DiseaseController) GetDisease(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	di, err := ctl.client.Disease.
		Query().
		Where(disease.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	c.JSON(200, di)
 }
 // ListDisease handles request to get a list of disease entities
// @Summary List disease entities
// @Description list disease entities
// @ID list-disease
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Disease
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /diseases [get]
func (ctl *DiseaseController) ListDisease(c *gin.Context) {
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
  
	diseases, err := ctl.client.Disease.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
		if err != nil {
		c.JSON(400, gin.H{"error": err.Error(),})
		return
	}
  
	c.JSON(200, diseases)
 }
 // DeleteDisease handles DELETE requests to delete a disease entity
// @Summary Delete a disease entity by ID
// @Description get disease by ID
// @ID delete-disease
// @Produce  json
// @Param id path int true "Disease ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /diseases/{id} [delete]
func (ctl *DiseaseController) DeleteDisease(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	err = ctl.client.Disease.
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
 // UpdateDisease handles PUT requests to update a disease entity
// @Summary Update a disease entity by ID
// @Description update disease by ID
// @ID update-disease
// @Accept   json
// @Produce  json
// @Param id path int true "Disease ID"
// @Param disease body ent.Disease true "Disease entity"
// @Success 200 {object} ent.Disease
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /diseases/{id} [put]
func (ctl *DiseaseController) UpdateDisease(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	obj := ent.Disease{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Disease binding failed",
		})
		return
	}
	obj.ID = int(id)
	fmt.Println(obj.ID)
	di, err := ctl.client.Disease.
		UpdateOne(&obj).
		SetDiseaseName(obj.DiseaseName).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed",})
		return
	}
  
	c.JSON(200, di)
 }
 // NewDiseaseController creates and registers handles for the disease controller
func NewDiseaseController(router gin.IRouter, client *ent.Client) *DiseaseController {
	uc := &DiseaseController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
 }
  
 // InitDiseaseController registers routes to the main engine
 func (ctl *DiseaseController) register() {
	diseases := ctl.router.Group("/diseases")
  
	diseases.GET("", ctl.ListDisease)
  
	// CRUD
	diseases.POST("", ctl.CreateDisease)
	diseases.GET(":id", ctl.GetDisease)
	diseases.PUT(":id", ctl.UpdateDisease)
	diseases.DELETE(":id", ctl.DeleteDisease)
 }