package domains

type NamesMatchDomain struct {
	MyName       NameDomain `json:"my_name"`
	RealNameList []RealName `json:"real_name_list"`
}
