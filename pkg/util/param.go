package util

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetIdFromParam(c *gin.Context) (int, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return 0, err
	}
	return id, nil
}

func GetDNIFromParam(c *gin.Context) string {
	return c.Request.URL.Query().Get("dni")
}

func GetDniMatriculaFromParam(c *gin.Context) (string, string) {
	return c.Param("dni"), c.Param("matricula")
}
