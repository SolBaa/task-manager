package routes

import (
	"database/sql"
	"net/http"

	"github.com/SolBaa/task-manager/internal/auth"
	"github.com/go-chi/chi/v5"
)

// RegisterRoutes registers all application routes
func SetupRouter(r *chi.Mux, db *sql.DB) chi.Router {
	// Initialize repositories, services, and handlers
	authRepo := auth.NewRepository(db)
	authService := auth.NewService(authRepo)
	authHandler := auth.NewHandler(authService)

	// projectRepo := project.NewRepository(db)
	// projectService := project.NewService(projectRepo)
	// projectHandler := project.NewHandler(projectService)

	// taskRepo := task.NewRepository(db)
	// taskService := task.NewService(taskRepo)
	// taskHandler := task.NewHandler(taskService)

	//Health check
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	// Initialize routes
	// Public routes
	r.Post("/register", authHandler.Register)
	r.Post("/login", authHandler.Login)

	// Protected routes
	// router.Route("/projects", func(r chi.Router) {
	//     r.Use(middleware.AuthMiddleware(authService))
	//     r.Post("/", projectHandler.Create)
	//     r.Get("/", projectHandler.GetAll)
	//     r.Put("/{id}", projectHandler.Update)
	//     r.Delete("/{id}", projectHandler.Delete)

	//     r.Route("/{projectId}/tasks", func(r chi.Router) {
	//         r.Post("/", taskHandler.Create)
	//         r.Get("/", taskHandler.GetAll)
	//         r.Put("/{taskId}", taskHandler.Update)
	//         r.Delete("/{taskId}", taskHandler.Delete)
	//     })
	// })

	return r

}
