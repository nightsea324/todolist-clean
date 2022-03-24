package po

import "main/service/member/domain/model/aggregate"

// ConvertMemberToPo -
func ConvertMemberToPo(in *aggregate.Member) *MemberPo {
	memberPo := &MemberPo{
		ID:       in.ID,
		Name:     in.Name,
		Password: in.Password,
	}

	return memberPo
}
