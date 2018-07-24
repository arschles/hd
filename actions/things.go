package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
)

func apiV1ThingsHandler(c buffalo.Context) error {
	things := []string{"thing1", "thing2", "thing3"}
	return c.Render(http.StatusOK, r.JSON(things))
}

func apiV2ThingsHandler(c buffalo.Context) error {
	things := []string{"thing4", "thing5", "thing6"}
	return c.Render(http.StatusOK, r.JSON(things))
}
