package mongo

import (
	"context"
	"errors"
	"main/service/member/domain/model/aggregate"
	"main/service/member/infra/persistence/mongo/po"
	"time"
)

// Create - 建立會員
func (m *memberRepo) Create(ctx context.Context, member *aggregate.Member) error {
	coll := m.db.Database(m.databaseName).Collection(memberC)

	memberPo := po.ConvertMemberToPo(member)
	memberPo.CreatedAt = time.Now().In(time.Local)

	if _, err := coll.InsertOne(ctx, memberPo); err != nil {
		errors.New("create member failed: %s", err.Error())
		return err
	}

	return nil
}
