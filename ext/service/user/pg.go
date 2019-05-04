package user

import (
	"context"

	"github.com/jinzhu/gorm"

	"github.com/luanngominh/goa-example/model"
)

type pgService struct {
	db *gorm.DB
}

//NewPGService return pgservice
func NewPGService(db *gorm.DB) Service {
	return &pgService {
		db: db,
	}
}

func (s *pgService) Create(ctx context.Context, u *model.User) (*model.User, error) {
	if err := s.db.Create(u).Error; err != nil {
		return nil, err
	}

	return u, nil
}

func (s *pgService) Update(ctx context.Context, u *model.User) (*model.User, error) {
	if err := s.db.Update(u).Error; err != nil {
		return nil, err
	}

	return u, nil
}

func (s *pgService) Get(ctx context.Context, q *UserQuery) (*model.User, error) {
	db := s.db

	if q.Email != "" {
		db = db.Where("email = ?", q.Email)
	}

	if !model.UUIDIsZero(&q.ID) {
		db = db.Where("id = ?", q.ID)
	}

	u := model.User{}
	return &u, db.First(&u).Error
}