package model

type BistouryUserApp struct {
	ID        int32
	App_code  string
	User_code string
}
type BistouryApp struct {
	Id         int32
	Code       string
	Name       string
	Group_code string
	Status     int
	Creator    string
}

type BistouryServer struct {
	Server_id              string
	Ip                     string
	Port                   int32
	Host                   string
	Log_dir                string
	Room                   string
	App_code               string
	Auto_jstack_enable     int
	Auto_jmap_histo_enable int
}

func (BistouryApp) TableName() string {
	return "bistoury_app"
}

func (BistouryServer) TableName() string {
	return "bistoury_server"
}

func (BistouryUserApp) TableName() string {
	return "bistoury_user_app"
}
