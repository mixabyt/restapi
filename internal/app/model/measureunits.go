package model

type MeasureUnits struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	AllowFractional bool   `json:"allow_fractional"`
}
