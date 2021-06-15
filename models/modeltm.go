package models

type ParamInput struct {
	Idtm          string `json:"tm_id"`
	EpochTenStart uint32 `json:"Epoch_start"`
	EpochTenEnd   uint32 `json:"Epoch_end"`
}

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

type DataSlice struct {
	Utc_tm []string `json:"tm_utc"`
	Avg_tm []float32 `json:"tm_avg"`
	Std_tm []float32 `json:"tm_std"`
	Min_tm []float32 `json:"tm_min"`
	Max_tm []float32 `json:"tm_max"`
	Q1_tm []float32 `json:"tm_q1"`
	Q2_tm []float32 `json:"tm_q2"`
	Q3_tm []float32 `json:"tm_q3"`
	Utc_ano1 []string `json:"tm_utc_ano1"`
	Ano1 []float32 `json:"tm_ano1"`
	Utc_ano2 []string `json:"tm_utc_ano2"`
	Ano2 []float32 `json:"tm_ano2"`
	Utc_ano3 []string `json:"tm_utc_ano3"`
	Ano3 []float32 `json:"tm_ano3"`
	Ano_bar []VerticalLine `json:"bar_ano"`
}

type LineStruct struct {
	// Color string `json:"color"`
	Width float32 `json:"width"`
}

type VerticalLine struct {
	Tyte string `json:"type"`
	X0 string `json:"x0"`
	Y0 uint8 `json:"y0"`
	Xref string `json:"xref"`
	Yref string `json:"yref"`
	X1 string `json:"x1"`
	Y1 uint8 `json:"y1"`
	Fillcolor string `json:"fillcolor"`
	Opacity float32 `json:"opacity"`
	Layer string `json:"layer"`
	Line LineStruct `json:"line"`
}
