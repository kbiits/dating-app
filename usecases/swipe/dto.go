package swipe_usecase

import (
	"time"

	"github.com/kbiits/dealls-take-home-test/domain/entity"
	"github.com/samber/mo"
)

type GetNextProfileResult struct {
	DisplayName   string                   `json:"display_name"`
	Bio           mo.Option[string]        `json:"bio"`
	DateOfBirth   mo.Option[time.Time]     `json:"date_of_birth"`
	Gender        mo.Option[entity.Gender] `json:"gender"`
	ProfilePicURL mo.Option[string]        `json:"profile_pic_url"`
}
