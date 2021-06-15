package models

// type CSVstruct struct {
// 	UTC          string `csv:"utc"`
// 	Epoch_ten    uint32 `csv:"epoch"`
// 	Avg          float32 `csv:"avg"`
// 	Max          float32 `csv:"max"`
// 	Min          float32 `csv:"min"`
// 	Std          float32 `csv:"std"`
// 	Q1           float32 `csv:"Q1"`
// 	Q2           float32 `csv:"Q2"`
// 	Q3           float32 `csv:"Q3"`
// 	AnomalyState float32 `csv:"anomalyState"`
// }

type CSVstruct struct {
	UTC          string `csv:"utc"`
	Epoch_ten    string `csv:"epoch"`
	Avg          string `csv:"avg"`
	Max          string `csv:"max"`
	Min          string `csv:"min"`
	Std          string `csv:"std"`
	Q1           string `csv:"Q1"`
	Q2           string `csv:"Q2"`
	Q3           string `csv:"Q3"`
	AnomalyState string `csv:"anomalyState"`
}