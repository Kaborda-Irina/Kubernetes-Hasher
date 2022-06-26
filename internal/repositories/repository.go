package repositories

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/Kaborda-Irina/Kubernetes-Hasher/internal/core/models"
	"github.com/Kaborda-Irina/Kubernetes-Hasher/internal/core/ports"
	"github.com/sirupsen/logrus"
)

type AppRepository struct {
	ports.IHashRepository
	db     *sql.DB
	logger *logrus.Logger
}

func NewAppRepository(db *sql.DB, logger *logrus.Logger) *AppRepository {
	return &AppRepository{
		IHashRepository: NewHashRepository(db, logger),
		db:              db,
		logger:          logger,
	}
}

func (ar AppRepository) CheckIsEmptyDB(kuberData models.KuberData) (bool, error) {
	var count int
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE name_deployment=$1 LIMIT 1;", os.Getenv("TABLE_NAME"))
	row := ar.db.QueryRow(query, kuberData.TargetName)
	err := row.Scan(&count)
	if err != nil {
		ar.logger.Info("err while scan row in database ", err)
		return false, err
	}

	if count < 1 {
		return true, nil
	}
	return false, nil
}
