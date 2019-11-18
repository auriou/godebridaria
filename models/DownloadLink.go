package models

type DownloadLink struct {
	Success bool `json:"success"`
	Infos   struct {
		Link      string        `json:"link"`
		Host      string        `json:"host"`
		Filename  string        `json:"filename"`
		Streaming []interface{} `json:"streaming"`
		Paws      bool          `json:"paws"`
		Filesize  int           `json:"filesize"`
		Streams   []struct {
			Quality  int    `json:"quality"`
			Ext      string `json:"ext"`
			Filesize int    `json:"filesize"`
			Name     string `json:"name"`
			Link     string `json:"link"`
			ID       string `json:"id"`
		} `json:"streams"`
		ID         string `json:"id"`
		HostDomain string `json:"hostDomain"`
	} `json:"infos"`
}
