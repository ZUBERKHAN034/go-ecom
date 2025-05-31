package routes

import (
	"log"
	"net/http"

	"github.com/ZUBERKHAN034/go-ecom/cmd/config"
	_ "github.com/ZUBERKHAN034/go-ecom/cmd/docs"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func SwaggerRoutes(router *mux.Router) {
	baseUrl := config.Env.BaseURL
	swaggerURL := baseUrl + "/swagger/doc.json"
	log.Println("Swagger UI is available at ", baseUrl+"/swagger/index.html")

	router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL(swaggerURL),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)
}
