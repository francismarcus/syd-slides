// f START OMIT
package routes

import (
	"context"
	"errors"
	"net/http"

	"github.com/francismarcus/cygnisyd/pkg/beer"
	"github.com/francismarcus/cygnisyd/pkg/entities"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// f END OMIT
// s START OMIT
// BeerResource ...
type BeerResource struct {
	Service beer.Service
}

// Routes is the beer router
func (rs *BeerResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Post("/", rs.Create)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("beers"))
	})

	return r
}

// s END OMIT
// main START OMIT

// CreatePostRequest ...
type CreatePostRequest struct {
	*entities.Beer
}

// Bind binds the create post request
func (pr *CreatePostRequest) Bind(r *http.Request) error {
	if pr.Beer == nil {
		return errors.New("missing required entity")
	}

	return nil
}

// main END OMIT
// create START OMIT

// Create creates a new beer
func (rs *BeerResource) Create(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	req := &CreatePostRequest{}

	err := render.Bind(r, req)

	if err != nil {
		render.JSON(w, r, err)
		return
	}

	res, err := rs.Service.CreateBeer(ctx, req.Beer)
	if err != nil {
		render.JSON(w, r, err)
	}

	render.JSON(w, r, res)
}

// create END OMIT
