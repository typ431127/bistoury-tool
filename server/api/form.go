package api

type createbody struct {
	Code      string `json:"code" binding:"required"`
	Name      string `json:"Name" binding:"required"`
	GroupCode string `json:"Group_code" binding:"required"`
	Status    int    `json:"Status"`
	Creator   string `json:"Creator"`
	IP        string `json:"ip" binding:"required"`
	Port      int32  `json:"port" binding:"required"`
	Hostname  string `json:"hostname" binding:"required"`
	Logdir    string `json:"logdir"`
}

type deletebody struct {
	Hostname string `json:"hostname"`
	IP       string `json:"IP"`
}
