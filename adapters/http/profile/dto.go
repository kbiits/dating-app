package http_profile

import (
	"time"

	"github.com/kbiits/dealls-take-home-test/domain/entity"
)

type CompleteProfileReq struct {
	DisplayName   string         `json:"display_name" validate:"required,min=3"`
	Bio           *string        `json:"bio" validate:"min=10"`
	Gender        *entity.Gender `json:"gender" validate:"oneof=male female"`
	Dob           *time.Time     `json:"date_of_birth"`
	ProfilePicURL *string        `json:"profile_pic_url"`
	DistrictID    *string        `json:"district_id"`
}
