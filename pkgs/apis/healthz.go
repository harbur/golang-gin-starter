package apis

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/harbur/golang-gin-starter/pkgs/store"
)

// Healthz handles healthz
// @Summary healthz
// @Description healthz
// @Tags healthz
// @Accept json
// @Produce json
// @Success 204 {string} string	"ok"
// @Router /healthz [get]
func Healthz(c *gin.Context) {
	c.JSON(http.StatusOK, store.Healthz())
}
