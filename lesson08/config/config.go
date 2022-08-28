package config

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
)

type bUrl string

// func (bu *bUrl) Decode(s string) error {
// 	if checkUrl(s) {
//     return nil
// 	}
// 	err := errors.New("неверный формат URL")
// 	return err
// }

func checkUrl(s string) bool {
	u, err := url.Parse(s)
	return err == nil && u.Scheme != "" && u.Host != ""
}

type Config struct {
	Port        int
	Db          bUrl
	Jaeger      bUrl
	Sentry      bUrl
	KafkaBroker string
	AppId       string
	AppKey      string
}

func Init() *Config {
	return &Config{
		Port:        getConfInt("PORT", 0),
		Db:          getConfUrl("DB", "postgres://db-user:dp-password@petstore-db:5432/petstore?sslmode=disable"),
		Jaeger:      getConfUrl("JAEGER", ""),
		Sentry:      getConfUrl("SENTRY", ""),
		KafkaBroker: getConfStr("KAFKA_BROKER", ""),
		AppId:       getConfStr("APP_ID", ""),
		AppKey:      getConfStr("APP_KEY", ""),
	}
}

func (c Config) ShowConfig() {
	fmt.Println("   PORT: ", c.Port)
	fmt.Println("     DB: ", c.Db)
	fmt.Println(" JAEGER: ", c.Jaeger)
	fmt.Println(" SENTRY: ", c.Sentry)
	fmt.Println("  KAFKA: ", c.KafkaBroker)
	fmt.Println(" APP_ID: ", c.AppId)
	fmt.Println("APP_KEY: ", c.AppKey)
}

func getConfStr(name string, def string) string {
	if v, ok := os.LookupEnv(name); ok {
		return v
	}
	return def
}

func getConfUrl(envName string, defVal string) bUrl {
	urlStr := getConfStr(envName, "")
	if checkUrl(urlStr) {
		return bUrl(urlStr)
	}
	return bUrl(defVal)
}

func getConfInt(envName string, defVal int) int {
	val := getConfStr(envName, "")
	if val, err := strconv.Atoi(val); err == nil {
		return val
	}
	return defVal
}
