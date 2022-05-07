package db

import (
	"time"

	"github.com/gjlanc65/gintest/models"
)

var Segments = []models.Segment{
	{ID: 1, LocalDate: time.Date(2022, 9, 10, 00, 00, 00, 00, time.FixedZone("UTC+10", 10*3600)), Name: "GJL", Details: "Leave Sydney", Who: "DML&GJL"},
}
