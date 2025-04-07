package main

import (
	"log"
	"net/http"
	"time"

	"github.com/rs/cors"
	"github.com/uvmain/uvsonic/handlers"
	"github.com/uvmain/uvsonic/logic"
)

func enableCdnCaching(w http.ResponseWriter) {
	expiryDate := time.Now().AddDate(1, 0, 0)
	w.Header().Set("Cache-Control", "public, max-age=31536000, immutable")
	w.Header().Set("Expires", expiryDate.String())
}

func StartServer() {
	router := http.NewServeMux()

	// frontend
	distDir := http.Dir("../dist")
	fileServer := http.FileServer(distDir)
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// serve static files
		if _, err := distDir.Open(r.URL.Path); err == nil {
			enableCdnCaching(w)
			fileServer.ServeHTTP(w, r)
			return
		}
		// serve index.html for non-static files
		http.ServeFile(w, r, "../dist/index.html")
	})

	//auth
	// router.HandleFunc("POST /api/login", auth.LoginHandler)
	// router.HandleFunc("GET /api/logout", auth.LogoutHandler)
	// router.HandleFunc("GET /api/check-session", auth.CheckSessionHandler)

	// public routes
	router.HandleFunc("/api/albums", handlers.HandleAlbums)
	router.HandleFunc("/api/songs", handlers.HandleSongs)

	// authenticated routes
	// router.Handle("DELETE /api/slugs/{slug}", auth.AuthMiddleware(http.HandlerFunc(handleDeleteImageBySlug)))

	handler := cors.AllowAll().Handler(router)

	var serverAddress string
	if logic.IsLocalDevEnv() {
		serverAddress = "localhost:8080"
		log.Println("Application running at https://uvsonic.localhost")
	} else {
		serverAddress = ":8080"
		log.Println("Application running at http://localhost:8080")
	}

	http.ListenAndServe(serverAddress, handler)
}
