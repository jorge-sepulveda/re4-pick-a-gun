package main

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/jorge-sepulveda/re4-pick-a-gun/core"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	r := mux.NewRouter()
	r.Use(LoggingMiddleware(logger))
	r.HandleFunc("/start", StartHandler).Methods("GET")
	r.HandleFunc("/roll", rollHandler).Methods("POST")
	r.HandleFunc("/load", loadHandler).Methods("POST")
	logger.Info("starting server", "port", 8080)
	http.ListenAndServe(":8080", r)

}

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func LoggingMiddleware(logger *slog.Logger) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			rw := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}
			next.ServeHTTP(rw, r)

			duration := time.Since(start).Milliseconds()

			logger.Info("api_request", slog.Group("request",
				slog.String("method", r.Method),
				slog.String("path", r.URL.Path),
				slog.Int("status", rw.statusCode),
				slog.Int64("duration_ms", duration),
				slog.String("client_ip", r.RemoteAddr),
				slog.Time("timestamp", time.Now()),
			))

		})
	}
}

func StartHandler(w http.ResponseWriter, r *http.Request) {
	sd := core.SaveData{}
	sd.StartGame("L", core.Handguns, core.Shotguns, core.Rifles, core.Subs, core.Magnums)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sd)
}

func rollHandler(w http.ResponseWriter, r *http.Request) {
	sd := core.SaveData{}
	err := json.NewDecoder(r.Body).Decode(&sd)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	if sd.CurrentChapter == sd.FinalChapter {
		http.Error(w, "All out of chapters, stranger!", http.StatusOK)
		return

	}
	sd.RollGun()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sd)
}

func loadHandler(w http.ResponseWriter, r *http.Request) {
	sd := core.SaveData{}
	requestBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Invalid Payload Data, unable to load", http.StatusBadRequest)
	}
	defer r.Body.Close()
	sd.LoadString(requestBytes)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sd)
}
