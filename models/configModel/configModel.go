package configModel

type System struct {
	AppName        string
	HttpAddr       string
	HttpPort       string
	SessionExpires string
}

type Database struct {
	DbType       string
	Host         string
	Port         int
	DbName       string
	Username     string
	Password     string
	Prefix       string
	Charset      string
	MaxIdleConns int
	MaxOpenConns int
	SslMode      string
	Path         string
}

type Log struct {
	FilePath string
	FileName string
	Level    string
	Mode     string
}
