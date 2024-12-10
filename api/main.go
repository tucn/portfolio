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
	// Get the absolute path to the templates folder
	basePath, err := os.Getwd()
	if err != nil {
		http.Error(w, "Could not determine base path", http.StatusInternalServerError)
		return
	}

	tmplPath := filepath.Join(basePath, "templates", "portfolio.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Error loading template: "+err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, data)
}
