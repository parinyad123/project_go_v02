package models

type CSVstruct struct {
	UTC          string
	epoch        uint32
	Avg          float32
	Max          float32
	Min          float32
	Std          float32
	Q1           float32
	Q2           float32
	Q3           float32
	AnomalyState float32
}
