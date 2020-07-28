package setting

import "time"

type ServerSettingS struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type AppSettingS struct {
	DefaultPageSize uint8
	MaxPageSize     uint8
	LogSavePath     string
	LogFileName     string
	LogFileExt      string
}

type DatabaseSettingS struct {
	DBType      string
	Host        string
	Port        uint16
	Username    string
	Password    string
	DBName      string
	Charset     string
	MaxIdleConn uint8
}

func (s *Setting) ReadSection(key string, data interface{}) error {
	if err := s.vp.UnmarshalKey(key, data); err != nil {
		return err
	}

	return nil
}
