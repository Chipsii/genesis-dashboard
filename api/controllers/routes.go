package controllers

import "github.com/KnockOutEZ/rest-api-portfolio/api/middlewares"

func (s *Server) initializeRoutes() {

	// Login Route
	s.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST", "OPTIONS")

	//Users routes
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST", "OPTIONS")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(s.GetUser)).Methods("GET", "OPTIONS")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(s.UpdateUser)).Methods("PUT", "OPTIONS")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(s.DeleteUser)).Methods("DELETE", "OPTIONS")

	//CustomSchemas routes
	s.Router.HandleFunc("/customschemas", middlewares.SetMiddlewareJSON(s.CreateCustomSchema)).Methods("POST", "OPTIONS")
	s.Router.HandleFunc("/mycustomschemas", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareJSON(s.GetMyCustomSchemas))).Methods("GET", "OPTIONS")
	s.Router.HandleFunc("/customschemas/{id}", middlewares.SetMiddlewareJSON(s.GetCustomSchema)).Methods("GET", "OPTIONS")
	//public route start
	s.Router.HandleFunc("/{key}/mycustomschemas", middlewares.SetMiddlewareJSON(s.GoGetAllCustomSchemas)).Methods("GET", "OPTIONS")
	s.Router.HandleFunc("/{key}/mycustomschemas/{id}", middlewares.SetMiddlewareJSON(s.GoGetOneCustomSchemas)).Methods("GET", "OPTIONS")
	//public route end
	s.Router.HandleFunc("/customschemas/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareJSON(s.UpdateCustomSchema))).Methods("PUT", "OPTIONS")
	s.Router.HandleFunc("/customschemas/{id}", middlewares.SetMiddlewareJSON(s.DeleteCustomSchema)).Methods("DELETE", "OPTIONS")

	//routes for admin
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.GetUsers)).Methods("GET", "OPTIONS")
	s.Router.HandleFunc("/customschemas", middlewares.SetMiddlewareJSON(s.GetCustomSchema)).Methods("GET", "OPTIONS")
}