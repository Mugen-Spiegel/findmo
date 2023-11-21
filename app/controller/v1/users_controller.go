package v1

import (
	"Mugen-Spiegel/findmo/app/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func (cont UserController) Create(r *gin.Context) {
	var task_create_params UserCreateParams
	if err := r.Bind(&task_create_params); err != nil {
		r.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		FirstName:  task_create_params.FirstName,
		LastName:   task_create_params.LastName,
		MiddleName: task_create_params.MiddleName,
		Email:      task_create_params.Email,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	r.JSON(http.StatusOK, gin.H{"data": user.Create()})
}

func (cont UserController) Index(r *gin.Context) {
	id, _ := strconv.Atoi(r.Param("id"))
	user := models.User{ID: id}
	r.JSON(http.StatusOK, gin.H{"data": user.Find()})
}
