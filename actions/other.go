package actions

import "github.com/gobuffalo/buffalo"

// otherHandler is a default handler to serve up
// a home page.
func otherHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("other.html"))
}
