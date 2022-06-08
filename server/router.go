package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Route struct {
	// Name is the name of this Route.
	Name string
	// Method is the string for the HTTP method. ex) GET, POST etc..
	Method string
	// Pattern is the pattern of the URI.
	Pattern string
	// HandlerFunc is the handler function of this route.
	HandlerFunc gin.HandlerFunc
}

// Routes is the list of the generated Route.
type Routes []Route



func AddService(engine *gin.Engine) *gin.RouterGroup {
	group := engine.Group("/test")

	for _, route := range routes {
		switch route.Method {
		case "GET":
			group.GET(route.Pattern, route.HandlerFunc)
		case "POST":
			group.POST(route.Pattern, route.HandlerFunc)
		case "PUT":
			group.POST(route.Pattern, route.HandlerFunc)
		case "DELETE":
			group.POST(route.Pattern, route.HandlerFunc)
		}

	}
	return group
}

var routes = Routes{
	{
		"Index",
		"GET",
		"/",
		Index,
	},
	{
		"Test1 ",
		"POST",
		"/test1",
		Test1,
	},

}

// Index is the index handler.
func Index(c *gin.Context) {
	c.String(http.StatusOK, "Hello World!")
}

func Test1(c *gin.Context) {

	c.Status(http.StatusNoContent)
}
