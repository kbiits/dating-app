package entity

import (
	"time"

	"github.com/samber/mo"
)

type Gender string

var (
	GenderMale   Gender = "male"
	GenderFemale Gender = "female"
)

type ProfileStatus string

var (
	ProfileStatusUnverified  ProfileStatus = "unverified"
	ProfileStatusVerified    ProfileStatus = "verified"
	ProfileStatusDeactivated ProfileStatus = "deactivated"
)

type Profile struct {
	ID          string               `db:"id"`
	UserID      string               `db:"user_id"`
	DisplayName string               `db:"display_name"`
	Bio         mo.Option[string]    `db:"bio"`
	DateOfBirth mo.Option[time.Time] `db:"date_of_birth"`
	DistrictID  mo.Option[string]    `db:"district_id"`
	Gender      mo.Option[Gender]    `db:"gender"`

	ProfilePicURL mo.Option[string] `db:"profile_pic_url"`
	Status        ProfileStatus     `db:"status"`
	CreatedAt     time.Time         `db:"created_at"`
	UpdatedAt     time.Time         `db:"updated_at"`
}

func (p *Profile) ShouldStatusVerified() bool {
	if p.Status == ProfileStatusDeactivated {
		return false
	}

	if p.Bio.IsAbsent() {
		return false
	}

	if p.DateOfBirth.IsAbsent() {
		return false
	}

	if p.DistrictID.IsAbsent() {
		return false
	}

	if p.Gender.IsAbsent() {
		return false
	}

	if p.ProfilePicURL.IsAbsent() {
		return false
	}

	return true
}
