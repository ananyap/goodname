package domains

type NameDomain struct {
	NameTitle     string   `json:"name_title"`
	DayBirth      string   `json:"day_birth"`
	LekReang      string   `json:"lek_reang"`
	LekSat        string   `json:"lek_sat"`
	Shadow        string   `json:"shadow"`
	Kalakini      []string `json:"kalakini"`
	NumberMiracle Number   `json:"number"`
}
