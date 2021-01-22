package config

import (
	"fmt"
	"os"

	"github.com/geeksmy/go-libs/jwt"
	"github.com/geeksmy/go-libs/redis"
	"github.com/jinzhu/configor"
	"go.uber.org/zap"
)

type Config struct {
	Debug    bool           `yaml:"debug,omitempty" default:"false" `
	Logger   LoggerConfig   `yaml:"logger,omitempty"`
	Database DatabaseConfig `yaml:"database,omitempty"`
	Jwt      JwtConfig      `yaml:"jwt,omitempty"`
	Server   ServerConfig   `yaml:"server,omitempty"`
	Metrics  MetricsConfig  `yaml:"metrics,omitempty"`
	Pprof    PprofConfig    `yaml:"pprof,omitempty"`
}

type DatabaseConfig struct {
	// 仅支持 mysql
	DSN          string `yaml:"dsn"`
	LogMode      bool   `yaml:"log_mode" default:"false"`
	MaxIdleConns int    `yaml:"max_idle_conns" default:"10"`
	MaxOpenConns int    `yaml:"max_open_conns" default:"100"`
	// format: https://golang.org/pkg/time/#ParseDuration
	ConnMaxLifetime string `yaml:"conn_max_lifetime" default:"1h"`
}

type MetricsConfig struct {
	Enabled bool   `yaml:"enabled" default:"false"`
	Addr    string `yaml:"addr" default:":31999"`
}

type PprofConfig struct {
	Enabled bool   `yaml:"enabled" default:"false"`
	Addr    string `yaml:"addr" default:"127.0.0.1:32999"`
}

type ServerConfig struct {
	Host     string `yaml:"host,omitempty" default:"0.0.0.0"`
	HTTPPort string `yaml:"http_port,omitempty" default:"8080"`
	GrpcPort string `yaml:"grpc_port,omitempty" default:"8082"`
}

type JwtConfig struct {
	Secret   string `json:"secret,omitempty"`
	ExpireIn int    `json:"expire_in,omitempty" default:"86400"`
}

type LoggerConfig struct {
	Level string `yaml:"level,omitempty" default:"debug"`
	// json or text
	Format string `yaml:"format,omitempty" default:"json"`
	// file
	Output string `yaml:"output,omitempty" default:""`
}

func initJwt(secret string) {
	jwt.C = jwt.Conf{
		Secret: secret,
	}
	err := jwt.Init()

	if err != nil {
		panic(fmt.Errorf("初始化 jwt 失败 %w", err))
	}
}

func InitLogger(debug bool, level, output string) {
	var conf zap.Config
	if debug {
		conf = zap.NewDevelopmentConfig()
	} else {
		conf = zap.NewProductionConfig()
	}

	var zapLevel = zap.NewAtomicLevel()
	if err := zapLevel.UnmarshalText([]byte(level)); err != nil {
		zap.L().Panic("设置日志记录级别失败",
			zap.Strings("only", []string{"debug", "info", "warn", "error", "dpanic", "panic", "fatal"}),
			zap.Error(err),
		)
	}

	conf.Level = zapLevel
	conf.Encoding = "console"

	if output != "" {
		conf.OutputPaths = []string{output}
		conf.ErrorOutputPaths = []string{output}
	}

	logger, _ := conf.Build()

	zap.RedirectStdLog(logger)
	zap.ReplaceGlobals(logger)
}

func InitRedis(conf redis.Conf) {
	redis.C = conf
}

var C = &Config{}

func Init(cfgFile string) {
	_ = os.Setenv("SCAN", "-")

	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)

	if cfgFile != "" {
		if err := configor.New(&configor.Config{AutoReload: true}).Load(C, cfgFile); err != nil {
			zap.L().Panic("init config fail", zap.Error(err))
		}
	} else {
		if err := configor.New(&configor.Config{AutoReload: true}).Load(C); err != nil {
			zap.L().Panic("init config fail", zap.Error(err))
		}
	}

	InitLogger(C.Debug, C.Logger.Level, C.Logger.Output)
	initJwt(C.Jwt.Secret)

	zap.L().Debug("[+]: 加载配置文件")
}

func init() {
	C = &Config{}
}
