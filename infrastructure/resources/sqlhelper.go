package resources

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"github.com/nelsonlpco/transactions/domain/entity"
	"github.com/nelsonlpco/transactions/domain/valueobjects"
	"github.com/nelsonlpco/transactions/infrastructure/inframodel"
	"github.com/sirupsen/logrus"
)

func CreateDatabase(db *sql.DB, schemaPath string) {
	logrus.WithField("schemaPath", schemaPath).Trace()

	data, err := os.ReadFile(schemaPath)
	if err != nil {
		log.Panicf("error on open file: %v", err)
	}

	_, err = db.Exec(string(data))
	if err != nil {
		log.Panicf("error on exec sql: %v", err)
	}
}

func SeedOperationsType(db *sql.DB) {
	tx, err := db.Begin()
	if err != nil {
		logrus.Panic(err)
	}

	debit1, _ := new(inframodel.OperationTypeModel).FromEntity(entity.NewOperationType(uuid.New(), "COMPRA A VISTA", valueobjects.Debit))
	debit2, _ := new(inframodel.OperationTypeModel).FromEntity(entity.NewOperationType(uuid.New(), "COMPRA PARCELADA", valueobjects.Debit))
	debit3, _ := new(inframodel.OperationTypeModel).FromEntity(entity.NewOperationType(uuid.New(), "SAQUE", valueobjects.Debit))
	credit, _ := new(inframodel.OperationTypeModel).FromEntity(entity.NewOperationType(uuid.New(), "PAGAMENTO", valueobjects.Credit))

	operationTypes := []*inframodel.OperationTypeModel{
		debit1,
		debit2,
		debit3,
		credit,
	}

	query := "INSERT INTO operation_type(id, description, operation) VALUES(?,?,?)"
	stmt, err := tx.Prepare(query)
	if err != nil {
		logrus.Panic(err)
	}
	defer stmt.Close()

	for _, operation := range operationTypes {
		_, err := stmt.Exec(operation.Id, operation.Description, operation.Operation)
		if err != nil {
			logrus.Panic(err)
		}

		id := uuid.New()
		id.UnmarshalBinary(operation.Id)
		logrus.Trace(operation.Description, " - ", id.String())
	}
	tx.Commit()
}
