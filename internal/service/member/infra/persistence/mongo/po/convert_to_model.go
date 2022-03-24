package po

import "main/service/member/domain/model/aggregate"

// ConvertMemberToModel
func ConvertMemberToModel(in *MemberPo) *aggregate.Member {
	return &aggregate.Member{
		ID:       in.ID,
		Name:     in.Name,
		Password: in.Password,
	}
}
