package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/kavit-munshi/bookings/pkg/config"
	"github.com/kavit-munshi/bookings/pkg/models"
	"github.com/kavit-munshi/bookings/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{})
}
func (m *Repository) Greet(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = string(time.Now().Local().Format("02-Jan-2006 15:04:05"))
	stringMap["name"] = r.URL.Query().Get(":name")
	stringMap["remote_ip"] = m.App.Session.GetString(r.Context(), "remote_ip")
	fmt.Println("The name is ", stringMap["name"])
	fmt.Println("The IP is ", stringMap["remote_ip"])
	render.RenderTemplate(w, "greet.page.tmpl", &models.TemplateData{StringMap: stringMap})
}
