package repository

import (
	"context"
	"facade/config"
	"facade/internal/model"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"testing"
)

func TestUserRepository(t *testing.T) {
	cfg := config.GetConfig()
	db, err := gorm.Open(postgres.Open(cfg.DbConnString), &gorm.Config{FullSaveAssociations: true})
	if err != nil {
		panic(err.Error())
	}

	ctx := context.Background()
	userRepo := NewUserRepository(db)
	err = userRepo.Create(ctx, &model.User{})
	assert.NoError(t, err)
}
