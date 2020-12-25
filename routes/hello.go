package routes

import (
	"github.com/go-chi/chi"

	"../controllers"
//	"../middleware"
)

func Routes() *chi.Mux{
	router:= chi.NewRouter()
	router.Get("/hello", controllers.Hello)
	//router.With(...).Post(...)
	return router
}
