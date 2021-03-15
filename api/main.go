// f START OMIT
package main

import (
	"context"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/francismarcus/cygnisyd/pkg/beer"
	"github.com/francismarcus/cygnisyd/pkg/routes"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// f END OMIT

// main START OMIT
func main() {
	db, err := initDB()

	if err != nil {
		panic("db err")
	}

	c := db.Collection("beers")
	rep := beer.NewRepository(c)
	service := beer.NewService(rep)
	resource := routes.BeerResource{
		Service: service,
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	r.Mount("/api", resource.Routes())
	http.ListenAndServe(":9000", r)
}

// main END OMIT

// db START OMIT
func initDB() (*mongo.Database, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	uri := "mongodb:/x:y@localhost:1230/<dbname>?retryWrites=true&w=majority"
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	db := client.Database("cygni")
	return db, nil
}

// db END OMIT
