package drawer

import (
	log "github.com/sirupsen/logrus"
	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	Log struct {
		Level  string
		Colors bool
	}
	Image struct {
		Resolution int // amount of pixels in both dimensions (x and y)
		Offset     struct {
			X float32
			Y float32
		}
	}
	Algorithm struct {
		Iterations  int
		ScaleFactor float32
		Parallel    bool
	}
}

var GlobalConfig Config

var configEnvPrefix string

func fullEnvVarName(name string) string {
	if configEnvPrefix != "" {
		name = configEnvPrefix + "_" + name
	}
	return name
}

func defineConfigValue(key string, defaultValue interface{}, envVarName string) {
	viper.SetDefault(key, defaultValue)
	viper.BindEnv(key, fullEnvVarName(envVarName))
}

func DefineConfig() {
	viper.SetTypeByDefaultValue(true)

	defineConfigValue("Log.Level", "debug", "LOG_LEVEL")
	defineConfigValue("Log.Colors", false, "LOG_COLORS")

	viper.AutomaticEnv()
}

func DefineCommandLineConfig() {
	flag.IntP("width", "w", 5000, "width of an image")
	viper.BindPFlag("Image.Resolution", flag.Lookup("width"))

	flag.IntP("iterations", "i", 200,
		"iterations to check if P is escaping to infinity")
	viper.BindPFlag("Algorithm.Iterations", flag.Lookup("iterations"))

	flag.Float32P("scale", "s", 0,
		"scale factor to use for projecting pixels onto complex value plane")
	viper.BindPFlag("Algorithm.ScaleFactor", flag.Lookup("scale"))

	flag.BoolP("parallel", "p", false, "use parallel computations")
	viper.BindPFlag("Algorithm.Parallel", flag.Lookup("parallel"))

	flag.Float32P("offset.x", "x", 0,
		"offset horizontal center point of M on image with")
	viper.BindPFlag("Image.Offset.X", flag.Lookup("offset.x"))

	flag.Float32P("offset.y", "y", 0,
		"offset vertical center point of M on image with")
	viper.BindPFlag("Image.Offset.Y", flag.Lookup("offset.y"))

}

func LoadConfig() {
	DefineConfig()
	err := viper.Unmarshal(&GlobalConfig)
	if err != nil {
		log.WithError(err).Fatal("failed to load configuration")
	}
}

func InitLogger() {
	log.SetOutput(os.Stdout)
	level, err := log.ParseLevel(GlobalConfig.Log.Level)
	if err != nil {
		log.WithError(err).
			WithField("input", GlobalConfig.Log.Level).
			Panic("cannot determine log level")
	}
	log.SetLevel(level)
	log.SetFormatter(&log.TextFormatter{
		ForceColors:            GlobalConfig.Log.Colors,
		DisableLevelTruncation: true,
	})
}

func Init() {
	DefineConfig()
	DefineCommandLineConfig()
	flag.Parse()
	LoadConfig()
	InitLogger()
}
