package application

import (
	auth "github.com/pathai95441/muang_thai_vote_service/src/authorization"
	candidate_domain "github.com/pathai95441/muang_thai_vote_service/src/domain/candidate"
	user_domain "github.com/pathai95441/muang_thai_vote_service/src/domain/user"
	"github.com/pathai95441/muang_thai_vote_service/src/repositories/candidate"
	"github.com/pathai95441/muang_thai_vote_service/src/repositories/user"
	"github.com/pathai95441/muang_thai_vote_service/src/services/commands"
	"github.com/pathai95441/muang_thai_vote_service/src/services/queries"
)

type Commands struct {
	AddNewCandidate     commands.IAddNewCandidateHandler
	UpdateCandidateInfo commands.IUpdateCandidateInfoHandler
	CreateNewUser       commands.ICreateNewUserHandler
}

type Queries struct {
	GetAllCandidate queries.IGetAllCandidateQuery
}

type Domain struct {
	CandidateDomain candidate_domain.ICandidateDomain
	UserDomain      user_domain.IUserDomain
}

type Repositories struct {
	CandidateRepo candidate.IRepository
	UserRepo      user.IRepository
}

type Application struct {
	Commands      Commands
	Queries       Queries
	Authorization auth.IAuthHandler
}
