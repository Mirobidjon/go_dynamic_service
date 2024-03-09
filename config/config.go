package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

const (
	// DebugMode indicates service mode is debug.
	DebugMode = "debug"
	// TestMode indicates service mode is test.
	TestMode = "test"
	// ReleaseMode indicates service mode is release.
	ReleaseMode = "release"

	LocalMode = "local"
)

type Config struct {
	ServiceName string
	Environment string // debug, test, release
	Version     string

	HTTPPort   string
	HTTPScheme string

	MongoHost1    string
	MongoHost2    string
	MongoPort     int
	MongoUser     string
	MongoPassword string
	MongoDatabase string

	PostgresHost     string
	PostgresPort     int
	PostgresUser     string
	PostgresPassword string
	PostgresDatabase string

	PostgresMaxConnections int32

	AuthServiceHost string
	AuthGRPCPort    string

	DynamicServiceHost string
	DynamicGRPCPort    string

	MinioEndpoint    string
	MinioHost        string
	MinioAccessKeyID string
	MinioSecretKey   string
	MinioImageHost   string
	RPCPort          string

	SecretKey string

	PasscodePool   string
	PasscodeLength int

	DefaultOffset string
	DefaultLimit  string

	TGBotToken string
	TGChatId   int64

	// CacheTTL is time to live for cache
	CacheTTL int64

	Username  string
	Password  string
	SignInKey string
}

// Load ...
func Load() Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	config := Config{}

	config.ServiceName = cast.ToString(getOrReturnDefaultValue("SERVICE_NAME", "ss_go_dynamic_service"))
	config.Environment = cast.ToString(getOrReturnDefaultValue("ENVIRONMENT", DebugMode))
	config.Version = cast.ToString(getOrReturnDefaultValue("VERSION", "1.0"))

	config.HTTPPort = cast.ToString(getOrReturnDefaultValue("HTTP_PORT", ":8080"))
	config.HTTPScheme = cast.ToString(getOrReturnDefaultValue("HTTP_SCHEME", "http"))

	config.RPCPort = cast.ToString(getOrReturnDefaultValue("RPC_PORT", ":5004"))

	config.MongoHost1 = cast.ToString(getOrReturnDefaultValue("MONGO_HOST", "172.26.10.13"))
	config.MongoHost2 = cast.ToString(getOrReturnDefaultValue("MONGO_HOST2", "172.26.10.13"))
	config.MongoPort = cast.ToInt(getOrReturnDefaultValue("MONGO_PORT", 27017))
	config.MongoUser = cast.ToString(getOrReturnDefaultValue("MONGO_USER", "ss_go_dynamic_service"))
	config.MongoPassword = cast.ToString(getOrReturnDefaultValue("MONGO_PASSWORD", ""))
	config.MongoDatabase = cast.ToString(getOrReturnDefaultValue("MONGO_DATABASE", "ss_go_dynamic_service"))

	config.PostgresHost = cast.ToString(getOrReturnDefaultValue("POSTGRES_HOST", ""))
	config.PostgresPort = cast.ToInt(getOrReturnDefaultValue("POSTGRES_PORT", 5432))
	config.PostgresUser = cast.ToString(getOrReturnDefaultValue("POSTGRES_USER", ""))
	config.PostgresPassword = cast.ToString(getOrReturnDefaultValue("POSTGRES_PASSWORD", ""))
	config.PostgresDatabase = cast.ToString(getOrReturnDefaultValue("POSTGRES_DATABASE", ""))

	config.PostgresMaxConnections = cast.ToInt32(getOrReturnDefaultValue("POSTGRES_MAX_CONNECTIONS", 30))

	config.AuthServiceHost = cast.ToString(getOrReturnDefaultValue("AUTH_SERVICE_HOST", "grpc-auth.soliqservis.local"))
	config.AuthGRPCPort = cast.ToString(getOrReturnDefaultValue("AUTH_GRPC_PORT", ":8080"))

	config.DynamicServiceHost = cast.ToString(getOrReturnDefaultValue("DYNAMIC_SERVICE_HOST", "localhost"))
	config.DynamicGRPCPort = cast.ToString(getOrReturnDefaultValue("DYNAMIC_GRPC_PORT", ":9103"))

	config.DefaultOffset = cast.ToString(getOrReturnDefaultValue("DEFAULT_OFFSET", "0"))
	config.DefaultLimit = cast.ToString(getOrReturnDefaultValue("DEFAULT_LIMIT", "10"))

	config.TGBotToken = cast.ToString(getOrReturnDefaultValue("TG_BOT_TOKEN", ""))
	config.TGChatId = cast.ToInt64(getOrReturnDefaultValue("TG_CHAT_ID", "-1001514410398"))

	config.CacheTTL = cast.ToInt64(getOrReturnDefaultValue("CACHE_TTL", 10))

	config.Username = cast.ToString(getOrReturnDefaultValue("USERNAME", ""))
	config.Password = cast.ToString(getOrReturnDefaultValue("PASSWORD", ""))
	config.SignInKey = cast.ToString(getOrReturnDefaultValue("SIGN_IN_KEY", ""))

	return config
}

func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}

	return defaultValue
}
