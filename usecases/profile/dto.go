package profile_usecase

import (
	"time"

	"github.com/kbiits/dealls-take-home-test/domain/entity"
	"github.com/samber/mo"
)

type CompleteProfileSpec struct {
	DisplayName   string                   `json:"display_name" validate:"required"`
	Bio           mo.Option[string]        `json:"bio"`
	Gender        mo.Option[entity.Gender] `json:"gender" validate:"oneof=male female"`
	Dob           mo.Option[time.Time]     `json:"date_of_birth"`
	ProfilePicURL mo.Option[string]        `json:"profile_pic_url"`
	DistrictID    mo.Option[string]        `json:"district_id"`
}

type ProfileResult struct {
	ID            string                   `json:"id"`
	UserID        string                   `json:"user_id"`
	DisplayName   string                   `json:"display_name"`
	Bio           mo.Option[string]        `json:"bio"`
	Gender        mo.Option[entity.Gender] `json:"gender"`
	ProfilePicURL mo.Option[string]        `json:"profile_pic_url"`
	Status        entity.ProfileStatus     `json:"status"`
}
