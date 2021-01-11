package controllers
 
import (
   "context"
   "fmt"
   "strconv"
   "github.com/sut63/team08/ent"
   "github.com/sut63/team08/ent/operative"
   "github.com/gin-gonic/gin"
)
 
// OperativeController defines the struct for the operative controller
type OperativeController struct {
   client *ent.Client
   router gin.IRouter
}
// CreateOperative handles POST requests for adding operative entities
// @Summary Create operative
// @Description Create operative
// @ID create-operative
// @Accept   json
// @Produce  json
// @Param operative body ent.Operative true "Operative entity"
// @Success 200 {object} ent.Operative
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /operatives [post]
func (ctl *OperativeController) CreateOperative(c *gin.Context) {
	obj := ent.Operative{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Operative binding failed",
		})
		return
	}
  
	f, err := ctl.client.Operative.
		Create().
		SetOperativeName(obj.OperativeName).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}
  
	c.JSON(200, f)
 }
 // GetOperative handles GET requests to retrieve a operative entity
// @Summary Get a operative entity by ID
// @Description get operative by ID
// @ID get-operative
// @Produce  json
// @Param id path int true "Operative ID"
// @Success 200 {object} ent.Operative
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /operatives/{id} [get]
func (ctl *OperativeController) GetOperative(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	f, err := ctl.client.Operative.
		Query().
		Where(operative.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	c.JSON(200, f)
 }
 // ListOperative handles request to get a list of operative entities
// @Summary List operative entities
// @Description list operative entities
// @ID list-operative
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Operative
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /operatives [get]
func (ctl *OperativeController) ListOperative(c *gin.Context) {
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
  
	operatives, err := ctl.client.Operative.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
		if err != nil {
		c.JSON(400, gin.H{"error": err.Error(),})
		return
	}
  
	c.JSON(200, operatives)
 }
 // DeleteOperative handles DELETE requests to delete a operative entity
// @Summary Delete a operative entity by ID
// @Description get operative by ID
// @ID delete-operative
// @Produce  json
// @Param id path int true "operative ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /operatives/{id} [delete]
func (ctl *OperativeController) DeleteOperative(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	err = ctl.client.Operative.
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
 // UpdateOperative handles PUT requests to update a operative entity
// @Summary Update a operative entity by ID
// @Description update operative by ID
// @ID update-operative
// @Accept   json
// @Produce  json
// @Param id path int true "operative ID"
// @Param operative body ent.Operative true "Operative entity"
// @Success 200 {object} ent.Operative
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /operatives/{id} [put]
func (ctl *OperativeController) UpdateOperative(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	obj := ent.Operative{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Operative binding failed",
		})
		return
	}
	obj.ID = int(id)
	fmt.Println(obj.ID)
	f, err := ctl.client.Operative.
		UpdateOne(&obj).
		SetOperativeName(obj.OperativeName).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed",})
		return
	}
  
	c.JSON(200, f)
 }
 // NewOperativeController creates and registers handles for the fund controller
func NewOperativeController(router gin.IRouter, client *ent.Client) *OperativeController {
	uc := &OperativeController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
 }
  
 // InitOperativeController registers routes to the main engine
 func (ctl *OperativeController) register() {
	operatives := ctl.router.Group("/operatives")
  
	operatives.GET("", ctl.ListOperative)
	// CRUD
	operatives.POST("", ctl.CreateOperative)
	operatives.GET(":id", ctl.GetOperative)
	operatives.PUT(":id", ctl.UpdateOperative)
	operatives.DELETE(":id", ctl.DeleteOperative)
 }
 
 