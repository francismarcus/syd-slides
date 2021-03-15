package beer

// main START OMIT
import (
	"context"

	"github.com/francismarcus/cygnisyd/pkg/entities"
)

type Service interface {
	CreateBeer(ctx context.Context, beer *entities.Beer) (*entities.Beer, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) CreateBeer(ctx context.Context, beer *entities.Beer) (*entities.Beer, error) {
	return s.repository.CreateBeer(ctx, beer)
}

// main END OMIT
