package repository

import (
	"log"

	"github.com/kazuhei/goweb-tutorial/domain/entity"
)

type UserRepository interface {
	GetUsers() []entity.User
	// getUser(id int) User
	// addUser(id int, username string) User
	// deleteUser(id int) error
}

type UserDbRepository struct {
	db string
}

func NewUserDbRepository(db string) *UserDbRepository {
	return &UserDbRepository{
		db: db,
	}
}

func (r *UserDbRepository) GetUsers() []entity.User {
	log.Println(r.db)
	return []entity.User{
		entity.User{ID: 1, Username: "taro"},
		entity.User{ID: 2, Username: "jiro"},
		entity.User{ID: 3, Username: "saburo"},
		entity.User{ID: 4, Username: "shiro"},
	}
}

type UserCachedDbRepository struct {
	db    string
	cache string
}

func NewUserCachedDbRepository(db, cache string) *UserCachedDbRepository {
	return &UserCachedDbRepository{
		db:    db,
		cache: cache,
	}
}

func (r *UserCachedDbRepository) GetUsers() []entity.User {
	log.Println(r.cache)
	log.Println(r.db)
	return []entity.User{
		entity.User{ID: 1, Username: "a"},
	}
}
