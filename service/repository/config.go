package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/i"
)

type Config interface {
	List(req *pb.ListQuery) ([]*pb.Config, error)
	Total(req *pb.ListQuery) (int64, error)
	Create(config *pb.Config) error
	Delete(config *pb.Config) (bool, error)
	Update(config *pb.Config) error
	Get(config *pb.Config) error
	Exist(config *pb.Config) bool
}

type ConfigRepository struct {
	DB *gorm.DB
}
