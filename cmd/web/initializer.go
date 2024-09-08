package main

import (
	"BS_Hackathon/internal/handlers"
	"BS_Hackathon/internal/repositories"
	"BS_Hackathon/internal/services"
	"context"
	_ "context"
	"database/sql"
	_ "firebase.google.com/go"
	firebase "firebase.google.com/go"
	"fmt"
	"google.golang.org/api/option"
	_ "google.golang.org/api/option"
	"log"
	"net/http"
)

type application struct {
	errorLog               *log.Logger
	infoLog                *log.Logger
	userHandler            *handlers.UserHandler
	eventHandler           *handlers.EventHandler
	bookHandler            *handlers.BookHandler
	saleHandler            *handlers.SaleHandler
	achievementHandler     *handlers.AchievementHandler
	userBookHandler        *handlers.UserBookHandler
	userEventHandler       *handlers.UserEventHandler
	userAchievementHandler *handlers.UserAchievementHandler
	fcmHandler             *handlers.FCMHandler
	favoriteHandler        *handlers.FavoriteHandler
	reviewHandler          *handlers.ReviewHandler
}

func initializeApp(db *sql.DB, errorLog, infoLog *log.Logger) *application {

	ctx := context.Background()
	sa := option.WithCredentialsFile("/root/go/src/BS_Hackathon/cmd/web/serviceAccountKey.json")

	firebaseApp, err := firebase.NewApp(ctx, &firebase.Config{ProjectID: "librarygogol"}, sa)
	if err != nil {
		errorLog.Fatalf("Ошибка в нахождении приложения: %v\n", err)
	}

	fcmClient, err := firebaseApp.Messaging(ctx)
	if err != nil {
		errorLog.Fatalf("Ошибка при неверном ID устройства: %v\n", err)
	}

	fcmHandler := handlers.NewFCMHandler(fcmClient, db)

	userRepo := &repositories.UserRepository{Db: db}
	userService := &services.UserService{Repo: userRepo}
	userHandler := &handlers.UserHandler{Service: userService}

	eventRepo := &repositories.EventRepository{Db: db}
	eventService := &services.EventService{Repo: eventRepo}
	eventHandler := &handlers.EventHandler{Service: eventService}

	bookRepo := &repositories.BookRepository{Db: db}
	bookService := &services.BookService{Repo: bookRepo}
	bookHandler := &handlers.BookHandler{Service: bookService}

	saleRepo := &repositories.SaleRepository{Db: db}
	saleService := &services.SaleService{Repo: saleRepo}
	saleHandler := &handlers.SaleHandler{Service: saleService}

	achievementRepo := &repositories.AchievementRepository{Db: db}
	achievementService := &services.AchievementService{Repo: achievementRepo}
	achievementHandler := &handlers.AchievementHandler{Service: achievementService}

	userBookRepo := &repositories.UserBookRepository{Db: db}
	userBookService := &services.UserBookService{Repo: userBookRepo}
	userBookHandler := &handlers.UserBookHandler{Service: userBookService}

	userEventRepo := &repositories.UserEventRepository{Db: db}
	userEventService := &services.UserEventService{Repo: userEventRepo}
	userEventHandler := &handlers.UserEventHandler{Service: userEventService}

	userAchievementRepo := &repositories.UserAchievementRepository{Db: db}
	userAchievementService := &services.UserAchievementService{Repo: userAchievementRepo}
	userAchievementHandler := &handlers.UserAchievementHandler{Service: userAchievementService}

	favoriteRepo := &repositories.FavoriteRepository{Db: db}
	favoriteService := &services.FavoriteService{Repo: favoriteRepo}
	favoriteHandler := &handlers.FavoriteHandler{Service: favoriteService}

	reviewRepo := &repositories.ReviewRepository{Db: db}
	reviewService := &services.ReviewService{Repo: reviewRepo}
	reviewHandler := &handlers.ReviewHandler{Service: reviewService}

	return &application{
		errorLog:               errorLog,
		infoLog:                infoLog,
		userHandler:            userHandler,
		eventHandler:           eventHandler,
		bookHandler:            bookHandler,
		saleHandler:            saleHandler,
		achievementHandler:     achievementHandler,
		userBookHandler:        userBookHandler,
		userEventHandler:       userEventHandler,
		userAchievementHandler: userAchievementHandler,
		fcmHandler:             fcmHandler,
		favoriteHandler:        favoriteHandler,
		reviewHandler:          reviewHandler,
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Printf("%v", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Printf("%v", err)
		panic("failed to connect to database")
		return nil, err
	}
	db.SetMaxIdleConns(35)
	if err = db.Ping(); err != nil {
		log.Printf("%v", err)
		panic("failed to ping the database")
		return nil, err
	}
	fmt.Println("successfully connected")

	return db, nil
}

func addSecurityHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cross-Origin-Opener-Policy", "same-origin")
		w.Header().Set("Cross-Origin-Embedder-Policy", "require-corp")
		w.Header().Set("Cross-Origin-Resource-Policy", "same-origin")
		next.ServeHTTP(w, r)
	})
}
