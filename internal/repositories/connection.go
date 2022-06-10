package repositories

import (
	"database/sql"
	"fmt"
	"github.com/Kaborda-Irina/Kubernetes-Hasher/internal/core/models"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"log"
)

func InitializeDB(connectionDB models.ConnectionDB, logger *logrus.Logger) (*sql.DB, error) {
	DBURL := fmt.Sprintf("host=%v port=%s user=%s dbname=%s sslmode=disable password=%s", connectionDB.DbHost, connectionDB.DbPort, connectionDB.DbUser, connectionDB.DbName, connectionDB.DbPassword)

	db, err := sql.Open(connectionDB.Dbdriver, DBURL)
	if err != nil {
		fmt.Printf("Cannot connect to %s database ", connectionDB.Dbdriver)
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the %s database ", connectionDB.Dbdriver)
	}

	//ticker := time.NewTicker(5 * time.Second)
	//go func() {
	//	if true {
	//		for range ticker.C {
	//			_, err := db.Exec("SELECT * FROM hashfiles;")
	//			if err != nil {
	//				log.Fatalln(err)
	//				os.Exit(1)
	//			}
	//		}
	//	}
	//}()

	return db, nil
}
