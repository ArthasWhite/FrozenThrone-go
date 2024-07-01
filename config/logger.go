package config

type Logger struct {
	Level      string `yaml:"level"`       // Level 最低日志等级，DEBUG<INFO<WARN<ERROR<FATAL 例如：info-->收集info等级以上的日志
	FileName   string `yaml:"file_name"`   // FileName 日志文件位置
	MaxSize    int    `yaml:"max_size"`    // MaxSize 进行切割之前，日志文件的最大大小(MB为单位)，默认为100MB
	MaxAge     int    `yaml:"max_age"`     // MaxAge 是根据文件名中编码的时间戳保留旧日志文件的最大天数。
	MaxBackups int    `yaml:"max_backups"` // MaxBackups 是要保留的旧日志文件的最大数量。默认是保留所有旧的日志文件（尽管 MaxAge 可能仍会导致它们被删除。）
}
