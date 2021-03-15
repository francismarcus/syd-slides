// f START OMIT
package beer

import (
	"context"

	"github.com/francismarcus/cygnisyd/pkg/entities"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Reader ...
type Reader interface{}

// Writer ...
type Writer interface {
	CreateBeer(ctx context.Context, e *entities.Beer) (*entities.Beer, error)
}

// Repository holds the set of method signatures
type Repository interface {
	Reader
	Writer
}

// f END OMIT
// main START OMIT
type repository struct {
	c *mongo.Collection
}

// NewRepository returns the Repository
func NewRepository(c *mongo.Collection) Repository {
	return &repository{
		c: c,
	}
}

func (r *repository) CreateBeer(ctx context.Context, e *entities.Beer) (*entities.Beer, error) {
	e.ID = primitive.NewObjectID()
	_, err := r.c.InsertOne(ctx, e)
	if err != nil {
		return nil, err
	}

	return e, nil
}

// main END OMIT
