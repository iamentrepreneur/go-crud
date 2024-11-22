package handler

import (
	"github.com/gin-gonic/gin"
	go_crud "go-crud"
)

func (h *Handler) signUp(c *gin.Context) {
	var input go_crud.User

	if err := c.BindJSON(&input); err != nil {
	}
}

func (h *Handler) signIn(c *gin.Context) {

}
