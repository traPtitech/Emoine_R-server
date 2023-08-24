package pbconv

import (
	"github.com/samber/lo"
	"github.com/traPtitech/Emoine_R/model/dbschema"
	emoine_rv1 "github.com/traPtitech/Emoine_R/pkg/pbgen/emoine_r/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func FromDBMeeting(m dbschema.Meeting) *emoine_rv1.Meeting {
	return &emoine_rv1.Meeting{
		Id:          m.ID.String(),
		VideoId:     m.VideoID,
		Title:       m.Title,
		Thumbnail:   m.Thumbnail,
		Description: lo.Ternary(m.Description.Valid, m.Description.String, ""),
		StartedAt:   timestamppb.New(m.StartedAt),
		EndedAt:     lo.Ternary(m.EndedAt.Valid, timestamppb.New(m.EndedAt.Time), nil),
	}
}
