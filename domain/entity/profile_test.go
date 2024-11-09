package entity

import (
	"testing"
	"time"

	"github.com/samber/mo"
)

func TestProfile_ShouldStatusVerified(t *testing.T) {
	type fields struct {
		ID            string
		UserID        string
		DisplayName   string
		Bio           mo.Option[string]
		DateOfBirth   mo.Option[time.Time]
		DistrictID    mo.Option[string]
		Gender        mo.Option[Gender]
		ProfilePicURL mo.Option[string]
		Status        ProfileStatus
		CreatedAt     time.Time
		UpdatedAt     time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "should return false if status is deactivated",
			fields: fields{
				Status: ProfileStatusDeactivated,
			},
			want: false,
		},
		{
			name: "should return false if bio is absent",
			fields: fields{
				Status: ProfileStatusVerified,
				Bio:    mo.None[string](),
			},
			want: false,
		},
		{
			name: "should return false if date of birth is absent",
			fields: fields{
				Status:      ProfileStatusVerified,
				Bio:         mo.Some("bio value"),
				DateOfBirth: mo.None[time.Time](),
			},
			want: false,
		},
		{
			name: "should return false if district id is absent",
			fields: fields{
				Status:      ProfileStatusVerified,
				Bio:         mo.Some("bio value"),
				DateOfBirth: mo.Some(time.Now()),
				DistrictID:  mo.None[string](),
			},
			want: false,
		},
		{
			name: "should return false if gender is absent",
			fields: fields{
				Status:      ProfileStatusVerified,
				Bio:         mo.Some("bio value"),
				DateOfBirth: mo.Some(time.Now()),
				DistrictID:  mo.Some("district id"),
				Gender:      mo.None[Gender](),
			},
			want: false,
		},
		{
			name: "should return false if profile pic is absent",
			fields: fields{
				Status:        ProfileStatusVerified,
				Bio:           mo.Some("bio value"),
				DateOfBirth:   mo.Some(time.Now()),
				DistrictID:    mo.Some("district id"),
				Gender:        mo.Some(GenderFemale),
				ProfilePicURL: mo.None[string](),
			},
			want: false,
		},
		{
			name: "should return true if all fields is filled",
			fields: fields{
				Status:        ProfileStatusVerified,
				Bio:           mo.Some("bio value"),
				DateOfBirth:   mo.Some(time.Now()),
				DistrictID:    mo.Some("district id"),
				Gender:        mo.Some(GenderFemale),
				ProfilePicURL: mo.Some("profile pic url"),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Profile{
				ID:            tt.fields.ID,
				UserID:        tt.fields.UserID,
				DisplayName:   tt.fields.DisplayName,
				Bio:           tt.fields.Bio,
				DateOfBirth:   tt.fields.DateOfBirth,
				DistrictID:    tt.fields.DistrictID,
				Gender:        tt.fields.Gender,
				ProfilePicURL: tt.fields.ProfilePicURL,
				Status:        tt.fields.Status,
				CreatedAt:     tt.fields.CreatedAt,
				UpdatedAt:     tt.fields.UpdatedAt,
			}
			if got := p.ShouldStatusVerified(); got != tt.want {
				t.Errorf("Profile.ShouldStatusVerified() = %v, want %v", got, tt.want)
			}
		})
	}
}
