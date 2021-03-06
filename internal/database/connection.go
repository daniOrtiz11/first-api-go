package database

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"github.com/daniOrtiz11/table-booking/internal/utils"
	//postgresql driver
	_ "github.com/lib/pq"
)

func getConnection() (*sql.DB, error) {
	host := utils.GetEnv("DB_HOST", "")
	port, errPort := strconv.Atoi(utils.GetEnv("DB_PORT", ""))
	if errPort != nil {
		log.Println(errPort)
		return nil, errPort
	}
	user := utils.GetEnv("DB_USER", "")
	password := utils.GetEnv("DB_PASS", "")
	dbname := utils.GetEnv("DB_NAME", "")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return db, nil

}
