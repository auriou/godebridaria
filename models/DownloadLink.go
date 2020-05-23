package models

type DownloadLink struct {
	Status string `json:"status"`
	Data   struct {
		Link       string `json:"link"`
		Host       string `json:"host"`
		Filename   string `json:"filename"`
		HostDomain string `json:"hostDomain"`
	} `json:"data"`
}
