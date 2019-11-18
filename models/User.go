package models

type User struct {
	Success bool   `json:"success"`
	Token   string `json:"token"`
	User    struct {
		Username             string `json:"username"`
		Email                string `json:"email"`
		IsPremium            bool   `json:"isPremium"`
		PremiumUntil         int    `json:"premiumUntil"`
		Lang                 string `json:"lang"`
		PreferedDomain       string `json:"preferedDomain"`
		LimitedHostersQuotas struct {
			Filefactory int `json:"filefactory"`
			Gigapeta    int `json:"gigapeta"`
			Isra        int `json:"isra"`
			Rapidu      int `json:"rapidu"`
			Brupload    int `json:"brupload"`
			Userscloud  int `json:"userscloud"`
			Wipfiles    int `json:"wipfiles"`
			Anzfile     int `json:"anzfile"`
		} `json:"limitedHostersQuotas"`
	} `json:"user"`
}
