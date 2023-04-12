package infra

import (
    "github.com/VictorNapoles/challenge-bravo/infra/cache"
    "github.com/VictorNapoles/challenge-bravo/infra/database"
    env "github.com/VictorNapoles/challenge-bravo/infra/environment"
    "github.com/VictorNapoles/challenge-bravo/infra/http_client"
    "github.com/joho/godotenv"
    "os"
    "time"
)

const (
    ENV_FILE_PATH         = ".env"
    ENV_KEY               = "ENV"
    LOCAL_ENV_VALUE       = "LOCAL"
    DATABASE_NAME_ENV_VAR = "DATABASE_NAME"
)

var (
    environment          env.Environment
    httpClient           http_client.HttpClient
    databaseConnection   database.MongoDatabaseConnection
    redisCacheConnection cache.RedisCacheConnection
)

func loadEnvironment() {

    processEnv := os.Getenv(ENV_KEY)

    if processEnv == LOCAL_ENV_VALUE {
        godotenv.Load(ENV_FILE_PATH)
    }
    environment = env.NewEnvironment()

}

func loadHttpClient() {
    httpClient = http_client.NewHttpClient(10 * time.Second)
}

func loadDatabaseConnection() {
    var err error
    databaseConnection, err = database.NewDatabaseConnection(environment)

    if err != nil {
        panic(err.Error())
    }
}

func loadRedisCacheConnection() {
    var err error
    redisCacheConnection, err = cache.NewCacheConnection(environment)

    if err != nil {
        panic(err.Error())
    }
}

func LoadInfra() {
    loadEnvironment()
    loadHttpClient()
    loadDatabaseConnection()
    loadRedisCacheConnection()
}

func GetEnvironment() env.Environment {
    if environment == nil {
        loadEnvironment()
    }
    return environment
}

func GetDatabaseName() (string, error) {
    database, err := GetEnvironment().Get(DATABASE_NAME_ENV_VAR)
    if err != nil {
        return "", err
    }

    return database, nil
}
func GetHttpClient() http_client.HttpClient {
    if httpClient == nil {
        loadHttpClient()
    }
    return httpClient
}

func GetMongoDatabaseConnection() database.MongoDatabaseConnection {
    if databaseConnection == nil {
        loadDatabaseConnection()
    }
    return databaseConnection
}

func GetRedisCacheConnection() cache.RedisCacheConnection {
    if redisCacheConnection == nil {
        loadRedisCacheConnection()
    }
    return redisCacheConnection
}