package models

import (


)
// ต้องสร้าง struct ที่มีชื่อสอดคล้องกับ table ใน db
type TmTest02Tsurvobs struct {
	Id int `json:"id"`
	Avg float32 `json:"avg"`
	Max float32 `json:"max"`
	Min float32 `json:"min"`
	Std float32 `json:"std"`
	Q1 float32 `json:"q1"`
	Q2 float32 `json:"q2"`
	Q3 float32 `json:"q3"`
	LostState float32 `json:"lost_state"`
	AnomalyState float32 `json:"anomaly_state"`
	UTC string `json:"utc"`
	EpochTen uint32 `json:"epoch_ten"`
}
