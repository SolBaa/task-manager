package routes

import (
	"database/sql"
	"net/http"

	"github.com/SolBaa/task-manager/internal/auth"
	"github.com/SolBaa/task-manager/internal/middleware"
	"github.com/SolBaa/task-manager/internal/project"
	"github.com/SolBaa/task-manager/internal/recipe"
	"github.com/go-chi/chi/v5"
)

// RegisterRoutes registers all application routes
func SetupRouter(r *chi.Mux, db *sql.DB) chi.Router {
	// Initialize repositories, services, and handlers
	authRepo := auth.NewRepository(db)
	authService := auth.NewService(authRepo)
	authHandler := auth.NewHandler(authService)

	pr := project.NewRepository(db)
	ps := project.NewService(pr)
	ph := project.NewHandler(ps)

	rr := recipe.NewRepository(db)
	rs := recipe.NewService(rr)
	rh := recipe.NewHandler(rs)

	// taskRepo := task.NewRepository(db)
	// taskService := task.NewService(taskRepo)
	// taskHandler := task.NewHandler(taskService)

	//Health check
	r.Group(func(r chi.Router) {
		r.Use(middleware.JwtMiddleware)
		r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("OK"))
		})
	})

	// Initialize routes
	// Public routes
	r.Post("/register", authHandler.Register)
	r.Post("/login", authHandler.Login)

	// Protected routes
	r.Route("/projects", func(r chi.Router) {
		r.Use(middleware.JwtMiddleware)
		r.Post("/", ph.CreateProject)
		r.Get("/", ph.GetAll)
		// r.Put("/{id}", projectHandler.Update)
		// r.Delete("/{id}", projectHandler.Delete)

		//     r.Route("/{projectId}/tasks", func(r chi.Router) {
		//         r.Post("/", taskHandler.Create)
		//         r.Get("/", taskHandler.GetAll)
		//         r.Put("/{taskId}", taskHandler.Update)
		//         r.Delete("/{taskId}", taskHandler.Delete)
		//     })

	})

	r.Route("/recipes", func(r chi.Router) {
		r.Use(middleware.JwtMiddleware)
		r.Post("/", rh.CreateRecipe)
		r.Get("/", rh.GetAll)
		r.Get("/{id}", rh.GetByID)

		// r.Put("/{id}", recipeHandler.Update)
		// r.Delete("/{id}", recipeHandler.Delete)
	})

	return r

}
