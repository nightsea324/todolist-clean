package mongo

import (
	"main/service/member/domain/repo"

	"go.mongodb.org/mongo-driver/mongo"
)

//
const (
	memberC = "member"
)

// memberRepo
type memberRepo struct {
	db           *mongo.Client
	databaseName string
}

// New -
func New(m *mongo.Client, databaseName string) repo.MemberRepo {
	return &memberRepo{
		db:           m,
		databaseName: databaseName,
	}
}
