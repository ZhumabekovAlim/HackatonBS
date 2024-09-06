package main

import (
	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
	"net/http"
)

func (app *application) routes() http.Handler {
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders, makeResponseJSON)

	dynamicMiddleware := alice.New()

	mux := pat.New()

	// USERS
	mux.Post("/users/signup", dynamicMiddleware.ThenFunc(app.userHandler.SignUp))               // sign up
	mux.Post("/users/login", dynamicMiddleware.ThenFunc(app.userHandler.LogIn))                 // login
	mux.Get("/users", standardMiddleware.ThenFunc(app.userHandler.GetAllUsers))                 // get all users
	mux.Get("/users/details/:id", standardMiddleware.ThenFunc(app.userHandler.GetUserByID))     // get one user info http://localhost:4000/clients/details/1
	mux.Del("/users/:id", standardMiddleware.ThenFunc(app.userHandler.DeleteUserByID))          // delete user by id
	mux.Put("/users/:id", standardMiddleware.ThenFunc(app.userHandler.UpdateUser))              // update user by id
	mux.Put("/users/password/:id", standardMiddleware.ThenFunc(app.userHandler.ChangePassword)) // update user balance

	return standardMiddleware.Then(mux)
}
