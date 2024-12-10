package handler

import (
	"fmt"
	"net/http"
)

// Handler is the required exported function for Vercel
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, `
<!DOCTYPE html>
<html>
<head>
    <title>Ngoc-Tu Chau - Portfolio</title>
    <style>
        body { font-family: Arial, sans-serif; margin: 40px; }
        h1 { color: #333; }
        p { margin: 10px 0; }
    </style>
</head>
<body>
    <h1>Ngoc-Tu Chau</h1>
    <h2>Work Experiences</h2>
    <p>Senior AppSec Engineer at National Australia Bank</p>
    <p>Previous roles: Researcher, Manager, Engineer</p>
    <h2>Education</h2>
    <p>Your academic background here</p>
</body>
</html>
	`)
}
