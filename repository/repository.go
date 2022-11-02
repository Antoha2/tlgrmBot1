package repository

import "gorm.io/gorm"

type Repository interface {
}

type repositoryImplDB struct {
	rep *gorm.DB
	Repository
}

func NewRepository(dbx *gorm.DB) *repositoryImplDB {
	return &repositoryImplDB{
		rep: dbx,
	}
}

/* func NewDB(dbx *gorm.DB) *repositoryImplDB {
	return &repositoryImplDB{
		rep: dbx,
	}
} */
