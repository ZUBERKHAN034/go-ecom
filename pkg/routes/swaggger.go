package routes

import (
	"log"
	"net/http"

	"github.com/ZUBERKHAN034/go-ecom/pkg/config"
	_ "github.com/ZUBERKHAN034/go-ecom/pkg/docs"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func SwaggerRoutes(router *mux.Router) {
	var baseUrl string
	if config.Env.BaseURL != "" {
		baseUrl = config.Env.BaseURL
	} else {
		log.Fatal("BASE_URL is not set in .env file")
	}

	swaggerURL := baseUrl + "/swagger/doc.json"
	log.Println("Swagger UI is available at ", baseUrl + "/swagger/index.html") // http://localhost:8080/swagger/index.html

	router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL(swaggerURL),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)
}
