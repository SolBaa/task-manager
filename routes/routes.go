package routes

import "github.com/go-chi/chi/v5"

func NewRouter(authHandler *auth.Handler, projectHandler *project.Handler, taskHandler *task.Handler, authService *auth.Service) *chi.Mux {
	r := chi.NewRouter()

	// Public routes
	r.Post("/register", authHandler.Register)
	r.Post("/login", authHandler.Login)

	// Protected routes
	r.Route("/projects", func(r chi.Router) {
		r.Use(middleware.AuthMiddleware(authService))
		r.Post("/", projectHandler.Create)
		r.Get("/", projectHandler.GetAll)
		r.Put("/{id}", projectHandler.Update)
		r.Delete("/{id}", projectHandler.Delete)

		r.Route("/{projectId}/tasks", func(r chi.Router) {
			r.Post("/", taskHandler.Create)
			r.Get("/", taskHandler.GetAll)
			r.Put("/{taskId}", taskHandler.Update)
			r.Delete("/{taskId}", taskHandler.Delete)
		})
	})

	return r
}
