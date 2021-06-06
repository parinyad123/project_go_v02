package models

type ParamInput struct {
	Idtm          string `json:"tm_id"`
	EpochTenStart uint32 `json:"Epoch_start"`
	EpochTenEnd   uint32 `json:"Epoch_end"`
}

// ต้องสร้าง struct ที่มีชื่อสอดคล้องกับ table ใน db
// โดยถ้า TmTest02Tsurvobs golang จะจับกับ tm_Test02_tsurvobs
// type TmTest02Tsurvobs struct {
// 	Id int `json:"id"`
// 	Avg float32 `json:"avg"`
// 	Max float32 `json:"max"`
// 	Min float32 `json:"min"`
// 	Std float32 `json:"std"`
// 	Q1 float32 `json:"q1"`
// 	Q2 float32 `json:"q2"`
// 	Q3 float32 `json:"q3"`
// 	LostState float32 `json:"lost_state"`
// 	AnomalyState float32 `json:"anomaly_state"`
// 	UTC string `json:"utc"`
// 	EpochTen uint32 `json:"epoch_ten"`
// }

// TmTestTsurvobs := TmTest02Tsurvobs
type tmStruct struct {
	Id           int     `json:"id"`
	Avg          float32 `json:"avg"`
	Max          float32 `json:"max"`
	Min          float32 `json:"min"`
	Std          float32 `json:"std"`
	Q1           float32 `json:"q1"`
	Q2           float32 `json:"q2"`
	Q3           float32 `json:"q3"`
	LostState    float32 `json:"lost_state"`
	AnomalyState float32 `json:"anomaly_state"`
	UTC          string  `json:"utc"`
	EpochTen     uint32  `json:"epoch_ten"`
}

// ต้องสร้าง struct ที่มีชื่อสอดคล้องกับ table ใน db
// โดยถ้า TmTest02Tsurvobs golang จะจับกับ tm_Test02_tsurvobs
type TmTest02Tsurvobs struct {
	tmStruct
}

type Tm0010010001Theos struct {
	tmStruct
}

type TmTestTsurvobs struct {
	tmStruct
}

// type tmStringdata struct {
// 	Id           string `json:"id"`
// 	Avg          string `json:"avg"`
// 	Max          string `json:"max"`
// 	Min          string `json:"min"`
// 	Std          string `json:"std"`
// 	Q1           string `json:"q1"`
// 	Q2           string `json:"q2"`
// 	Q3           string `json:"q3"`
// 	LostState    string `json:"lost_state"`
// 	AnomalyState string `json:"anomaly_state"`
// 	UTC          string `json:"utc"`
// 	EpochTen     string `json:"epoch_ten"`
// }

// type tmFload struct {
// 	Id           uint32  `json:"id"`
// 	Avg          float32 `json:"avg"`
// 	Max          float32 `json:"max"`
// 	Min          float32 `json:"min"`
// 	Std          float32 `json:"std"`
// 	Q1           float32 `json:"q1"`
// 	Q2           float32 `json:"q2"`
// 	Q3           float32 `json:"q3"`
// 	LostState    float32 `json:"lost_state"`
// 	AnomalyState float32 `json:"anomaly_state"`
// 	UTC          string  `json:"utc"`
// 	EpochTen     uint32  `json:"epoch_ten"`
// }
