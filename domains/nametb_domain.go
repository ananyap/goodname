package domains

type RealName struct {
	ThaiName           string   `json:"thainame"`
	ReangThai          string   `json:"reangthai"`
	LeksatThai         int      `json:"leksat_thai"`
	PairPoint          int      `json:"pair_point"`
	Shadow             int      `json:"shadow"`
	MatchedLekSatPlusD int      `json:"match_sat_plus_d"`
	MatchedShaNumPlusD int      `json:"match_sha_plus_d"`
	SurNameKalakini    []string `json:"surname_kalakini"`
}
