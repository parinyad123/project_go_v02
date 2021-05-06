package models

import (


)

type tm_data struct {
	Id int `json:"id"`
	Avg float32 
	Max float32
	Min float32
	Std float32
	Q1 float32
	Q2 float32
	Q3 float32
	LostState float32
	AnomalyState float32
	UTC string
	EpochTen uint32
}
