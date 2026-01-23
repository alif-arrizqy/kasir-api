package routes

import (
	"kasir-api/controllers"
	"net/http"
)

func ProductRoutes() {
	http.HandleFunc("/api/products", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			controllers.GetAllProducts(w, r)
		case "POST":
			controllers.CreateProduct(w, r)
		}
	})

	http.HandleFunc("/api/products/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			controllers.GetProductById(w, r)
		case "PUT":
			controllers.UpdateProduct(w, r)
		case "DELETE":
			controllers.DeleteProduct(w, r)
		}
	})
}