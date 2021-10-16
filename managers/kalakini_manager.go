package managers

type KalakiniManager struct {
	nameTitle     string
	dayBirth      string
	dayOfKalikini map[string]string
}

func NewKalakiniManager(nameTitle string, dayBirth string) *KalakiniManager {
	return &KalakiniManager{nameTitle: nameTitle, dayBirth: dayBirth,
		dayOfKalikini: map[string]string{"sunday": "ศษสหฬฮ", "monday": "ะ้ื์แ๊ั็าิึี่ำุูเใไโอ", "tuesday": "กขคฆง", "wednesday1": "จฉชซฌญ", "wednesday2": "บปผฝพฟภม", "thursday": "ดตถทธน", "friday": "ยรลว", "saturday": "ฎฏฐฑฒณ"}}
}

func (manager *KalakiniManager) Kalakini() (kalakinis []string) {

	nameTitleRune := manager.nameTitle

	for day, kalakiniList := range manager.dayOfKalikini {

		if manager.dayBirth == day {

			for _, charx := range string(nameTitleRune) {

				for _, kalakiniCharx := range kalakiniList {

					if charx == kalakiniCharx {
						//logrus.Println("KALAKINI", string(charx))
						kalakinis = append(kalakinis, string(charx))
						break
					}
				}
			}
			break
		}
	}

	return

}
