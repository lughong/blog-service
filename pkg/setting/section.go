package setting

import "time"

type ServerSettingS struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type AppSettingS struct {
	DefaultPageSize      int
	MaxPageSize          int
	LogSavePath          string
	LogFileName          string
	LogFileExt           string
	UploadSavePath       string
	UploadserverUrl      string
	UploadImageMaxSize   int
	UploadImageAllowExts []string
}

type DatabaseSettingS struct {
	DBType       string
	Host         string
	Port         int
	Username     string
	Password     string
	DBName       string
	TablePrefix  string
	Charset      string
	ParseTime    bool
	Loc          string
	MaxIdleConns int
	MaxOpenConns int
}

func (s *Setting) ReadSection(key string, data interface{}) error {
	if err := s.vp.UnmarshalKey(key, data); err != nil {
		return err
	}

	return nil
}
