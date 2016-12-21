package main

import (
	"html/template"
	"io"
	"net/http"

	"log"

	"github.com/kazuhei/goweb-tutorial/application/settings"
	"github.com/kazuhei/goweb-tutorial/domain/repository"
	"github.com/labstack/echo"
)

func main() {
	t := &Template{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}
	e := echo.New()
	e.Renderer = t
	settings := settings.NewApplicationSettings()
	e.GET("/", handler(settings, t))
	e.Logger.Fatal(e.Start(":1323"))
}

func handler(a *settings.ApplicationSettings, t *Template) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		log.Println(id)
		repository := getUserRepository(a)
		users := repository.GetUsers()
		var str string
		for _, user := range users {
			str = str + user.ToString() + " "
		}
		return c.Render(http.StatusOK, "hello", "World")
	}
}

func getUserRepository(s *settings.ApplicationSettings) repository.UserRepository {
	if s.IsDebug {
		return repository.NewUserDbRepository(s.DB)
	}
	return repository.NewUserCachedDbRepository(s.DB, s.Cache)
}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	log.Println("呼ばれてるよ")
	log.Println(name)
	log.Println(data)
	return t.templates.ExecuteTemplate(w, name, data)
}
