package utils

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetID(c *gin.Context) (uint, bool) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return uint(id), false
	}

	return uint(id), true
}
