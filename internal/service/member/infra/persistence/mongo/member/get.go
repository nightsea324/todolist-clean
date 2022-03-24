package mongo

import (
	"context"
	"errors"
	"main/service/member/domain/model/aggregate"
	"main/service/member/infra/persistence/mongo/po"

	"go.mongodb.org/mongo-driver/bson"
)

// GetByName - 透過名稱取得會員
func (m *memberRepo) GetByName(ctx context.Context, name string) (*aggregate.Member, error) {
	coll := m.db.Database(m.databaseName).Collection(memberC)

	filter := bson.M{
		"name": name,
	}

	result := new(po.MemberPo)
	if err := coll.FindOne(ctx, filter).Decode(&result); err != nil {
		errors.New("get member failed: %s", err.Error())
	}

	return po.ConvertMemberToModel(result), nil
}
