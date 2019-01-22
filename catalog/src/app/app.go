package app

import (
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gustavohenrique/grpc-golang-nodejs/catalog/src/interfaces"
	"github.com/gustavohenrique/grpc-golang-nodejs/catalog/src/app/product"
	"github.com/gustavohenrique/grpc-golang-nodejs/catalog/src/app/customer"
	"github.com/gustavohenrique/grpc-golang-nodejs/catalog/src/database"
)

type Application struct {
	Router *mux.Router
}

func New() *Application {
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "It is working.")
	})

	app := &Application{}
	app.Router = router
	return app
}

func (app *Application) GetServerWithCORS() http.Handler {
	headersOk := handlers.AllowedHeaders([]string{"X-USER-ID", "Content-Type"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "OPTIONS"})
	app.Router.Use(mux.CORSMethodMiddleware(app.Router))
	return handlers.CORS(originsOk, headersOk, methodsOk)(app.Router)
}

func (app *Application) Initialize(params map[string]string) {
	databaseURL := params["databaseURL"]
	db := database.NewDB(databaseURL)
	db.Connect()

	repositories := map[string]interfaces.Repository{
		"product": product.NewRepository(db),
		"customer": customer.NewRepository(db),
	}
	service := product.NewService(repositories)
	handler := product.NewHandler(service)

	app.Router.HandleFunc("/products", handler.FetchAll).Methods("GET")
}
