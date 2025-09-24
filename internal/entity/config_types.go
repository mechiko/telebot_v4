package entity

// import "github.com/spf13/viper"

// type IConfiguration interface{}

type ConfigInterface interface {
	Get() interface{}
	Save() error
	SaveAs(fn string) error
	SaveSafe() error
	GetByName(name string) interface{}
	Set(key string, value interface{}, save ...bool) error
	// GetViper() (*viper.Viper, error)
	Unmarshal(*Configuration) error
}

type Configuration struct {
	Hostname     string                `json:"hostname"`
	HostPort     string                `json:"hostport"`
	HtmlPort     string                `json:"htmlport"`
	Debug        bool                  `json:"debug"`
	Telebot      bool                  `json:"telebot"`
	Token        string                `json:"token"`
	TestGroupId  int64                 `json:"testgroupid"`
	AdminGroupId int64                 `json:"admingroupid"`
	GroupId      int64                 `json:"groupid"`
	ChannelId    int64                 `json:"channelid"`
	AdminId      int64                 `json:"adminid"`
	Bot          botConfiguration      `json:"bot"`
	Application  appConfiguration      `json:"application"`
	Layouts      layoutConfiguration   `json:"layouts"`
	Database     databaseConfiguration `json:"database"`
}

type layoutConfiguration struct {
	TimeLayout    string `json:"timelayout"`
	TimeLayoutDay string `json:"timelayoutday"`
}

type databaseConfiguration struct {
	ConnectionUri string `json:"connectionuri"`
	Driver        string `json:"driver"`
	DbName        string `json:"dbname"`
	Timeout       int    `json:"timeout"`
}

type appConfiguration struct {
}

type botConfiguration struct {
	Token string `json:"token"`
}
