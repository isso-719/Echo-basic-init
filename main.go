package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
	e.Renderer = renderer

	e.GET("/", Index)
	e.GET("/hello/:name", Hello)
	e.GET("/form", FormGet)
	e.POST("/form", FormPost)

	// If you are not using port 1323, please change to your port number.
	e.Logger.Fatal(e.Start(":1323"))
}

func Index(c echo.Context) error {
	Greeting := "Hello, World!"
	User := "Kazuki Isogai"
	return c.Render(http.StatusOK, "index.html", map[string]interface{}{
		"greeting": Greeting,
		"user":     User,
	})
}

func Hello(c echo.Context) error {
	Name := c.Param("name")
	return c.Render(http.StatusOK, "hello.html", map[string]interface{}{
		"name": Name,
	})
}

func FormGet(c echo.Context) error {
	return c.Render(http.StatusOK, "form_get.html", map[string]interface{}{
	})
}

func FormPost(c echo.Context) error {
	Name := c.FormValue("onamae")
	return c.Render(http.StatusOK, "form_post.html", map[string]interface{}{
		"name": Name,
	})
}