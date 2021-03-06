package routers

import (
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"github.com/wiseman-ska/tech-assessment/user-manager-api/commons"
	"github.com/wiseman-ska/tech-assessment/user-manager-api/controllers"
)

func SetupUserRoutes(authRouter *mux.Router) *mux.Router {
	userRouter := mux.NewRouter()
	authRouter.HandleFunc("/users/login", controllers.UserLoginHandler).Methods("POST")
	authRouter.HandleFunc("/users/create", controllers.UserRegisterHandler).Methods("POST")
	userRouter.HandleFunc("/api/v1/users/update/{id}", controllers.UserUpdateHandler).Methods("PUT")
	userRouter.HandleFunc("/api/v1/users/delete/{id}", controllers.UserDeleteHandler).Methods("DELETE")
	userRouter.HandleFunc("/api/v1/users/all", controllers.GetUsersHandler).Methods("GET")
	userRouter.HandleFunc("/api/v1/users/{id}", controllers.GetUserByIdHandler).Methods("GET")

	authRouter.PathPrefix("/api/v1").Handler(negroni.New(
		negroni.HandlerFunc(commons.Authorize), negroni.Wrap(userRouter),
	))
	return authRouter
}
