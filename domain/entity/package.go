package entity

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type PremiumPackage struct {
	ID          string               `db:"id"`
	Name        string               `db:"name"`
	Description string               `db:"description"`
	Price       float64              `db:"price"`
	Validity    int                  `db:"validity"` // in minutes
	CreatedAt   time.Time            `db:"created_at"`
	UpdatedAt   time.Time            `db:"updated_at"`
	Config      PremiumPackageConfig `db:"config"`
}

type PremiumPackageConfig struct {
	QuotaPerDay    int  `json:"quota_per_day"`
	UnlimitedQuota bool `json:"unlimited"`
}

func (p *PremiumPackageConfig) Scan(src any) error {
	switch s := src.(type) {
	case []byte:
		return json.Unmarshal(s, p)
	case string:
		return json.Unmarshal([]byte(s), p)
	default:
		return errors.New("incompatible type for PremiumPackage")
	}
}

func (p PremiumPackageConfig) Value() (driver.Value, error) {
	return json.Marshal(p)
}

var _ sql.Scanner = &PremiumPackageConfig{}
var _ driver.Valuer = PremiumPackageConfig{}
