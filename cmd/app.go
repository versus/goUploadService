package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"mime/multipart"
	"path/filepath"

	"github.com/go-macaron/binding"
	"github.com/go-macaron/toolbox"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/versus/gouploadservice/jwt"
	"gopkg.in/macaron.v1"
)

const (
	defaultPort              = 8080
	defaultTokenValidateDays = 1
	maxFileSize              = 104857600 // max file size for upload is 100Mb
)

var (
	secretKey         string
	tokenValidateDays int
	port              int
	opsProcessed      = promauto.NewCounter(prometheus.CounterOpts{
		Name: "upload_processed_ops_total",
		Help: "The total number of processed request",
	})
)

type uploadForm struct {
	File *multipart.FileHeader `form:"file"`
}

//Run is the function for start web service of application
func Run() {
	m := macaron.Classic()
	m.Use(toolbox.Toolboxer(m, toolbox.Options{
		HealthCheckURL: "/health", // URL for health check request
		HealthCheckers: []toolbox.HealthChecker{
			new(connectChecker),
		}, // Health checkers
		DisableDebug: true, // Turns off all debug functionality when true
	}))
	m.Get("/", func(ctx *macaron.Context) string {
		return "Use `curl -F 'file=@nameFile.ext' http://localhost:8080/upload` and see file into /tmp \n"
	})

	m.Post("/upload", binding.MultipartForm(uploadForm{}), func(uf uploadForm) string {
		if uf.File.Size > maxFileSize {
			log.Println("Upload error: File is too large")
			return "Error: file size is too large"
		}
		file, err := uf.File.Open()
		if err != nil {
			log.Println("Upload error:", err)
			return "Error: Upload error"
		}

		log.Println("Size file is ", uf.File.Size)

		buf := new(bytes.Buffer)
		buf.ReadFrom(file)
		fileName := "/tmp/" + filepath.Base(uf.File.Filename)
		err = ioutil.WriteFile(fileName, buf.Bytes(), 0644)
		if err != nil {
			log.Println("Upload error:", err)
			return "Error: Upload error"
		}
		opsProcessed.Inc()
		return "File " + filepath.Base(uf.File.Filename) + " has been successfully uploaded"
	})

	m.Get("/token", func(ctx *macaron.Context) string {
		//TODO: add security code for IP and/or AUTH for this endpoint
		log.Println("Warning: generated token for ip: " + ctx.RemoteAddr())
		return jwt.GenerateToken(secretKey, tokenValidateDays)
	})

	m.Get("/metrics", bearerValidate, promhttp.Handler())

	addr := fmt.Sprintf("0.0.0.0:%d", port)
	log.Printf("Server is running on %s ... \n", addr)
	http.ListenAndServe(addr, m)
	m.Run()
}
