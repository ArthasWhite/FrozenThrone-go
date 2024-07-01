package core

import (
	"github.com/Arthaslixin/FrozenThrone-go/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

// 负责设置 encoding 的日志格式
func getEncoder() zapcore.Encoder {
	// 获取一个指定的的EncoderConfig，进行自定义
	encodeConfig := zap.NewProductionEncoderConfig()

	// 设置每个日志条目使用的键。如果有任何键为空，则省略该条目的部分。

	// 序列化时间。eg: 2022-09-01T19:11:35.921+0800
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// "time":"2022-09-01T19:11:35.921+0800"
	encodeConfig.TimeKey = "time"
	// 将Level序列化为全大写字符串。例如，将info level序列化为INFO。
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	// 以 package/file:行 的格式 序列化调用程序，从完整路径中删除除最后一个目录外的所有目录。
	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewConsoleEncoder(encodeConfig)
}

// 负责日志写入的位置
func getLogWriter(filename string, maxsize, maxBackup, maxAge int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,  // 文件位置
		MaxSize:    maxsize,   // 进行切割之前,日志文件的最大大小(MB为单位)
		MaxAge:     maxAge,    // 保留旧文件的最大天数
		MaxBackups: maxBackup, // 保留旧文件的最大个数
		Compress:   false,     // 是否压缩/归档旧文件
	}
	// AddSync 将 io.Writer 转换为 WriteSyncer。
	// 它试图变得智能：如果 io.Writer 的具体类型实现了 WriteSyncer，我们将使用现有的 Sync 方法。
	// 如果没有，我们将添加一个无操作同步。

	// syncFile := zapcore.AddSync(lumberJackLogger) // 打印到文件
	// syncConsole := zapcore.AddSync(os.Stderr)     // 打印到控制台
	// return zapcore.NewMultiWriteSyncer(syncFile, syncConsole)
	return zapcore.AddSync(lumberJackLogger)
}

// InitLogger 初始化Logger
func InitLogger(logConf *config.Logger) *zap.Logger {
	// 获取日志写入位置
	writeSyncer := getLogWriter(logConf.FileName, logConf.MaxSize, logConf.MaxBackups, logConf.MaxAge)
	// 获取日志编码格式
	encoder := getEncoder()

	// 获取日志最低等级，即>=该等级，才会被写入。
	var l = new(zapcore.Level)

	// 创建一个将日志写入 WriteSyncer 的核心。
	core := zapcore.NewCore(encoder, writeSyncer, l)
	logger := zap.New(core, zap.AddCaller())

	// 替换zap包中全局的logger实例，后续在其他包中只需使用zap.L()调用即可
	zap.ReplaceGlobals(logger)
	return zap.L()
}

// // 日志颜色
// const (
// 	red    = 31
// 	yellow = 33
// 	blue   = 36
// 	gray   = 37
// )

// type LogFormatter struct{}

// func (t *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
// 	// 根据不同的level 显示颜色
// 	var levelColor int
// 	switch entry.Level {
// 	case logrus.DebugLevel, logrus.TraceLevel:
// 		levelColor = gray
// 	case logrus.WarnLevel:
// 		levelColor = yellow
// 	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
// 		levelColor = red
// 	default:
// 		levelColor = blue
// 	}

// 	var b *bytes.Buffer
// 	if entry.Buffer != nil {
// 		b = entry.Buffer
// 	} else {
// 		b = &bytes.Buffer{}
// 	}
// 	prefix := global.Config.Logger.Prefix
// 	// 自定义日期格式
// 	timestamp := entry.Time.Format("2006-01-02 15:04:05")
// 	if entry.HasCaller() {
// 		// 自定义文件路径
// 		funcVal := entry.Caller.Function
// 		fileVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)
// 		// 自定义输出格式
// 		fmt.Fprintf(b, "%s[%s] \x1b[%dm[%s]\x1b[0m %s %s %s]\n", prefix, timestamp, levelColor, entry.Level, fileVal, funcVal, entry.Message)
// 	} else {
// 		fmt.Fprintf(b, "%s[%s] \x1b[%dm[%s]\x1b[0m %s]\n", prefix, timestamp, levelColor, entry.Level, entry.Message)
// 	}
// 	return b.Bytes(), nil
// }

// func InitLogger() *logrus.Logger {
// 	mLog := logrus.New()                                // 创建一个实例
// 	mLog.SetOutput(os.Stdout)                           // 设置输出类型
// 	mLog.SetReportCaller(global.Config.Logger.ShowLine) // 开启返回函数和行号
// 	mLog.SetFormatter(&LogFormatter{})                  // 设置自己定义的Formatter
// 	level, err := logrus.ParseLevel(global.Config.Logger.Level)
// 	if err != nil {
// 		level = logrus.InfoLevel
// 	}
// 	mLog.SetLevel(level) // 设置最低的Level
// 	InitDefaultLogger()  // 修改 logrus 全局log样式
// 	return mLog
// }

// func InitDefaultLogger() {
// 	// 修改 logrus 全局log样式
// 	logrus.SetOutput(os.Stdout)                           // 设置输出类型
// 	logrus.SetReportCaller(global.Config.Logger.ShowLine) // 开启返回函数和行号
// 	logrus.SetFormatter(&LogFormatter{})                  // 设置自己定义的Formatter
// 	level, err := logrus.ParseLevel(global.Config.Logger.Level)
// 	if err != nil {
// 		level = logrus.InfoLevel
// 	}
// 	logrus.SetLevel(level) // 设置最低的Level
// }
