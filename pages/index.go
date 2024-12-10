package handler

import (
	"encoding/json"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

type PortfolioData struct {
	Name            string
	WorkExperiences []string
	Education       []string
}

func loadPortfolioData() PortfolioData {
	file, _ := os.Open(filepath.Join("data", "portfolio.json"))
	defer file.Close()

	var data PortfolioData
	json.NewDecoder(file).Decode(&data)
	return data
}

func Handler(w http.ResponseWriter, r *http.Request) {
	data := loadPortfolioData()
	// Load template from the templates folder
	tmplPath := filepath.Join("templates", "portfolio.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}
	// Render the template with the data
	tmpl.Execute(w, data)
}
