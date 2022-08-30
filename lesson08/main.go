package main

import (
	"binaryty/config"
	"os"
)

func main(){
	os.Clearenv()
	os.Setenv("PORT", "8080")
	//Добавляем переменную окружения с поломанной ссылкой для проверки
	os.Setenv("DB", "postgres/db-user:dp-password@petstore-db:5432/petstore?sslmode=disable")
	os.Setenv("JAEGER", "http://jaeger:16686")
	os.Setenv("SENTRY", "http://sentry:9000")
	os.Setenv("KAFKA_BROKER", "kafka:9092")
	os.Setenv("APP_ID", "myID")
	os.Setenv("APP_KEY", "4m234m32fe23f24m23412m3421m3")
	
	conf := config.Init()
  conf.ShowConfig()
}
