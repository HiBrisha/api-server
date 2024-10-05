package tests

import (
	"fmt"
	monggoDb "module/internal/database"
	api "module/internal/services/api"
	env "module/pkg/utils/environment"
	logger "module/pkg/utils/logs"
	"testing"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
)

var log logger.SysLog = logger.SysLog{
	StrDir: "../docs/logs/",
}

func TestLogger(t *testing.T) {

	fmt.Println("Start*****")
	log.Init()
	log.Write("Info", "Testing log")
	log.Close()
	fmt.Println("*****End")
}

type User struct {
	UserName string `bson:"name"`
	PassWord string `bson:"password"`
}

type dbInfo struct {
	DbName     string
	Collection string
}

func TestMonggoDB(t *testing.T) {
	log.Init()

	// Initialize Variable Environment
	err := godotenv.Load("../config/.env")
	if err != nil {
		log.Write("ERR", "Lỗi load file .env: %v", err)
	}
	log.Write("Info", "Load environment success")
	connStr := env.Get("monggoDB", &log)

	serverDb := monggoDb.Server{
		ConnStr: connStr,
		Log:     &log,
	}

	serverDb.Init()

	dbInfo := dbInfo{
		DbName:     "lib-management",
		Collection: "user",
	}

	MyUser := User{
		UserName: "admin",
		PassWord: "Pass@work1",
	}

	serverDb.Insert(dbInfo.DbName, dbInfo.Collection, MyUser)

	//Delete document
	filter := bson.M{"name": "admin"}
	serverDb.Delete(dbInfo.DbName, dbInfo.Collection, filter)

	defer serverDb.Close()

}

func TestAPIServer(t *testing.T) {
	log.Init()
	//Test API Server
	apiServer := &api.APIServer{
		Port: "3000",
		Log:  &log,
	}
	err := apiServer.Init()
	if err != nil {
		log.Write("ERR", "Unable to initialize server: %v\n", err)
	}
}
