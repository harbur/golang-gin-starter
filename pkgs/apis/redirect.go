package apis

import "github.com/gin-gonic/gin"

// RedirectTo returns a function that redirects to a specific URL
func RedirectTo(code int, url string) func(*gin.Context) {
	return func(c *gin.Context) {
		c.Redirect(code, url)
	}
}
