package application

import (
	"database/sql"
	"fmt"

	"github.com/pathai95441/muang_thai_vote_service/src/config"
	candidate_domain "github.com/pathai95441/muang_thai_vote_service/src/domain/candidate"
	"github.com/pathai95441/muang_thai_vote_service/src/repositories/candidate"
	"github.com/pathai95441/muang_thai_vote_service/src/services/commands"
	"github.com/pathai95441/muang_thai_vote_service/src/services/queries"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

var serverApp Application
var validate *validator.Validate

func initRepositories(db *sql.DB) Repositories {
	return Repositories{
		CandidateRepo: candidate.NewRepository(db),
	}
}

func initDatabase(dbConfig config.DBConfig) *sql.DB {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=UTC",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Database)
	db, err := sql.Open(dbConfig.Driver, dsn)

	if err != nil {
		println("tang888")
		println(err.Error())
		panic(err)
	}

	return db
}

func initDomain(repositories Repositories) Domain {
	return Domain{
		CandidateDomain: candidate_domain.NewCandidateDomain(repositories.CandidateRepo),
	}
}

func initQueries(domain Domain) Queries {
	return Queries{
		GetAllCandidate: queries.NewGetAllCandidateQuery(domain.CandidateDomain),
	}
}

func initCommands(domain Domain) Commands {
	return Commands{
		AddNewCandidate:     commands.NewAddNewCandidateHandler(domain.CandidateDomain),
		UpdateCandidateInfo: commands.NewUpdateCandidateInfoHandler(domain.CandidateDomain),
	}
}

func initApplication() {
	db := initDatabase(config.CurrentConfig.DBConfig)
	repositories := initRepositories(db)
	domain := initDomain(repositories)
	commandHandlers := initCommands(domain)
	queriesCmd := initQueries(domain)

	validate = validator.New()
	if err := validate.RegisterValidation("required_if", ValidateRequiredIfValue); err != nil {
		fmt.Printf("Register Validator error reason: %s\n", err.Error())
	}

	serverApp = Application{
		Commands: commandHandlers,
		Queries:  queriesCmd,
	}
}
