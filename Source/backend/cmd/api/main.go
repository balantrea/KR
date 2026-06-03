package main

import (
	"log"
	"net/http"
	"sports-backend/Source/backend/internal/config"
	"sports-backend/Source/backend/internal/database"
	"sports-backend/Source/backend/internal/handler"
	"sports-backend/Source/backend/internal/middleware"
	"sports-backend/Source/backend/internal/repository"
	"sports-backend/Source/backend/internal/service"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	cfg := config.LoadConfig()

	db, err := database.NewPostgresDB(cfg.DBConn)
	if err != nil {
		log.Fatalf("DB connection error: %v", err)
	}
	defer db.Close()

	userRepo := repository.NewUserRepository(db)
	sportRepo := repository.NewSportRepository(db)
	teamRepo := repository.NewTeamRepository(db)
	playerRepo := repository.NewPlayerRepository(db)
	matchRepo := repository.NewMatchRepository(db)

	authServ := service.NewAuthService(userRepo)
	sportServ := service.NewSportService(sportRepo)
	teamServ := service.NewTeamService(teamRepo)
	playerServ := service.NewPlayerService(playerRepo)
	matchServ := service.NewMatchService(matchRepo)

	authHand := handler.NewAuthHandler(authServ)
	sportHand := handler.NewSportHandler(sportServ)
	teamHand := handler.NewTeamHandler(teamServ)
	playerHand := handler.NewPlayerHandler(playerServ)
	matchHand := handler.NewMatchHandler(matchServ)

	mux := http.NewServeMux()

	authHand.RegisterRoutes(mux)
	sportHand.RegisterRoutes(mux)
	teamHand.RegisterRoutes(mux)
	playerHand.RegisterRoutes(mux)
	matchHand.RegisterRoutes(mux)

	protectedRouter := middleware.AuthMiddleware(mux)
	finalRouter := corsMiddleware(protectedRouter)

	log.Printf("Server starting on port %s...", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, finalRouter); err != nil {
		log.Fatal(err)
	}
}
