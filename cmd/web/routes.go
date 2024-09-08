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
	mux.Put("/users/password/:id", standardMiddleware.ThenFunc(app.userHandler.ChangePassword)) // update user password

	// PASSWORD RECOVERY
	mux.Post("/password/recovery", dynamicMiddleware.ThenFunc(app.userHandler.SendRecoveryHandler))
	mux.Get("/password/recovery/mail", dynamicMiddleware.ThenFunc(app.userHandler.PasswordRecoveryHandler))

	// EVENTS
	mux.Post("/events", dynamicMiddleware.ThenFunc(app.eventHandler.CreateEvent))      // create event
	mux.Get("/events", standardMiddleware.ThenFunc(app.eventHandler.GetAllEvents))     // get all events
	mux.Get("/events/:id", standardMiddleware.ThenFunc(app.eventHandler.GetEventByID)) // get event by id
	mux.Put("/events/:id", standardMiddleware.ThenFunc(app.eventHandler.UpdateEvent))  // update event by id
	mux.Del("/events/:id", standardMiddleware.ThenFunc(app.eventHandler.DeleteEvent))  // delete event by id

	// BOOKS
	mux.Post("/books", dynamicMiddleware.ThenFunc(app.bookHandler.CreateBook))         // create book
	mux.Get("/books", standardMiddleware.ThenFunc(app.bookHandler.GetAllBooks))        // get all books
	mux.Get("/books/:id", standardMiddleware.ThenFunc(app.bookHandler.GetBookByID))    // get book by id
	mux.Put("/books/:id", standardMiddleware.ThenFunc(app.bookHandler.UpdateBook))     // update book by id
	mux.Del("/books/:id", standardMiddleware.ThenFunc(app.bookHandler.DeleteBook))     // delete book by id
	mux.Get("/books/new", standardMiddleware.ThenFunc(app.bookHandler.GetAllNewBooks)) // get all new books

	// SALES
	mux.Post("/sales", dynamicMiddleware.ThenFunc(app.saleHandler.CreateSale))      // create sale
	mux.Get("/sales", standardMiddleware.ThenFunc(app.saleHandler.GetAllSales))     // get all sales
	mux.Get("/sales/:id", standardMiddleware.ThenFunc(app.saleHandler.GetSaleByID)) // get sale by id
	mux.Put("/sales/:id", standardMiddleware.ThenFunc(app.saleHandler.UpdateSale))  // update sale by id
	mux.Del("/sales/:id", standardMiddleware.ThenFunc(app.saleHandler.DeleteSale))  // delete sale by id

	// ACHIEVEMENTS
	mux.Post("/achievements", dynamicMiddleware.ThenFunc(app.achievementHandler.CreateAchievement))      // create achievement
	mux.Get("/achievements", standardMiddleware.ThenFunc(app.achievementHandler.GetAllAchievements))     // get all achievements
	mux.Get("/achievements/:id", standardMiddleware.ThenFunc(app.achievementHandler.GetAchievementByID)) // get achievement by id
	mux.Put("/achievements/:id", standardMiddleware.ThenFunc(app.achievementHandler.UpdateAchievement))  // update achievement by id
	mux.Del("/achievements/:id", standardMiddleware.ThenFunc(app.achievementHandler.DeleteAchievement))  // delete achievement by id

	// USER_BOOK
	mux.Post("/user_books", dynamicMiddleware.ThenFunc(app.userBookHandler.CreateUserBook))                     // create user_book
	mux.Get("/user_books", standardMiddleware.ThenFunc(app.userBookHandler.GetAllUserBooks))                    // get all user_books
	mux.Get("/user_books/:id", standardMiddleware.ThenFunc(app.userBookHandler.GetUserBookByID))                // get user_book by id
	mux.Put("/user_books/:id", standardMiddleware.ThenFunc(app.userBookHandler.UpdateUserBook))                 // update user_book by id
	mux.Del("/user_books/:id", standardMiddleware.ThenFunc(app.userBookHandler.DeleteUserBook))                 // delete user_book by id
	mux.Get("/user_books/expired/:id", standardMiddleware.ThenFunc(app.userBookHandler.GetAllUserExpiredBooks)) // get all expired books
	mux.Get("/user_books/now/:id", standardMiddleware.ThenFunc(app.userBookHandler.GetAllUserNowBooks))         // get all now books
	mux.Get("/user_books/return/:id", standardMiddleware.ThenFunc(app.userBookHandler.GetAllUserBooks))         // get all returned books

	// USER_EVENT
	mux.Post("/user_events", dynamicMiddleware.ThenFunc(app.userEventHandler.CreateUserEvent))      // create user_event
	mux.Get("/user_events", standardMiddleware.ThenFunc(app.userEventHandler.GetAllUserEvents))     // get all user_events
	mux.Get("/user_events/:id", standardMiddleware.ThenFunc(app.userEventHandler.GetUserEventByID)) // get user_event by id
	mux.Put("/user_events/:id", standardMiddleware.ThenFunc(app.userEventHandler.UpdateUserEvent))  // update user_event by id
	mux.Del("/user_events/:id", standardMiddleware.ThenFunc(app.userEventHandler.DeleteUserEvent))  // delete user_event by id

	//USER_ACHIEVEMENTS
	mux.Post("/user_achievements", dynamicMiddleware.ThenFunc(app.userAchievementHandler.CreateUserAchievement))      // create user_achievement
	mux.Get("/user_achievements", standardMiddleware.ThenFunc(app.userAchievementHandler.GetAllUserAchievements))     // get all user_achievements
	mux.Get("/user_achievements/:id", standardMiddleware.ThenFunc(app.userAchievementHandler.GetUserAchievementByID)) // get user_achievement by id
	mux.Put("/user_achievements/:id", standardMiddleware.ThenFunc(app.userAchievementHandler.UpdateUserAchievement))  // update user_achievement by id
	mux.Del("/user_achievements/:id", standardMiddleware.ThenFunc(app.userAchievementHandler.DeleteUserAchievement))  // delete user_achievement by id

	// AI
	mux.Post("/photo-detect", dynamicMiddleware.ThenFunc(app.bookHandler.PhotoDetect)) // detect photo and return book title

	// NOTIFY
	mux.Post("/notify", dynamicMiddleware.ThenFunc(app.fcmHandler.NotifyChange))             //send notification
	mux.Post("/notify/token/create", dynamicMiddleware.ThenFunc(app.fcmHandler.CreateToken)) //create notification
	mux.Del("/notify/token/:id", dynamicMiddleware.ThenFunc(app.fcmHandler.DeleteToken))     // delete notification

	// FAVORITES
	mux.Post("/favorites", dynamicMiddleware.ThenFunc(app.favoriteHandler.CreateFavorite))          // create favorite
	mux.Get("/favorites/user/id", standardMiddleware.ThenFunc(app.favoriteHandler.GetAllFavorites)) // get all favorites by user id
	mux.Get("/favorites/:id", standardMiddleware.ThenFunc(app.favoriteHandler.GetFavoriteByID))     // get favorite by id
	mux.Put("/favorites/:id", standardMiddleware.ThenFunc(app.favoriteHandler.UpdateFavorite))      // update favorite by id
	mux.Del("/favorites/:id", standardMiddleware.ThenFunc(app.favoriteHandler.DeleteFavorite))      // delete favorite by id

	// REVIEWS
	mux.Post("/reviews", dynamicMiddleware.ThenFunc(app.reviewHandler.CreateReview))           // create review
	mux.Get("/reviews/book/:id", standardMiddleware.ThenFunc(app.reviewHandler.GetAllReviews)) // get all reviews by book id
	mux.Get("/reviews/:id", standardMiddleware.ThenFunc(app.reviewHandler.GetReviewByID))      // get review by id
	mux.Put("/reviews/:id", standardMiddleware.ThenFunc(app.reviewHandler.UpdateReview))       // update review by id
	mux.Del("/reviews/:id", standardMiddleware.ThenFunc(app.reviewHandler.DeleteReview))       // delete review by id

	return standardMiddleware.Then(mux)
}
