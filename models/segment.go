package models

import (
	"time"
)

type Segment struct {
	ID        int       `json:"id"`
	LocalDate time.Time `form:"local_date" binding:"required" time_format:"2006-01-02"`
	Name      string    `json:"name"`
	Details   string    `json:"details"`
	Who       string    `json:"who"`
}
