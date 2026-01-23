package routes

import (
	"kasir-api/controllers"
	"net/http"
)

func CategoryRoutes() {
	http.HandleFunc("/api/categories", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			controllers.GetAllCategories(w, r)
		case "POST":
			controllers.CreateCategory(w, r)
		}
	})

	http.HandleFunc("/api/categories/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			controllers.GetCategoryById(w, r)
		case "PUT":
			controllers.UpdateCategory(w, r)
		case "DELETE":
			controllers.DeleteCategory(w, r)
		}
	})
}