package main

import (
	"context"
	"encoding/json"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
	"github.com/qr-decoder/helper"
	"github.com/qr-decoder/models"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	r.Post("/decode", DecodeHandler)

	server := &http.Server{
		Addr:    ":8082",
		Handler: r,
	}

	serverCtx, serverStopCtx := context.WithCancel(context.Background())

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig

		shutdownCtx, cancel := context.WithTimeout(serverCtx, 10*time.Second)
		defer cancel()

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				log.Println("Graceful shutdown timed out. Forcing exit...")
			}
		}()

		if err := server.Shutdown(shutdownCtx); err != nil {
			log.Println("Failed to gracefully shutdown:", err)
		}
		serverStopCtx()
	}()

	log.Println("Server started on port 8082")
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Println("Failed to start server:", err)
	}

	<-serverCtx.Done()
}

func DecodeHandler(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		log.Println("Error decoding image:", err)
		http.Error(w, "Error decoding image", http.StatusInternalServerError)
		return
	}

	bmp, err := gozxing.NewBinaryBitmapFromImage(img)
	if err != nil {
		log.Println("Error creating binary bitmap:", err)
		http.Error(w, "Error creating binary bitmap", http.StatusInternalServerError)
		return
	}

	qrReader := qrcode.NewQRCodeReader()
	result, err := qrReader.Decode(bmp, nil)
	if err != nil {
		log.Println("Error decoding QR code:", err)
		http.Error(w, "Error decoding QR code", http.StatusInternalServerError)
		return
	}

	qrText := result.GetText()
	parsedData := helper.ParseQRCode(qrText)

	resp := models.Response{
		Success: true,
		Message: "QR code decoded successfully",
		Data:    parsedData,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Println("Failed to encode response:", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
