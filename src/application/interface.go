package application

import (
	auth "github.com/pathai95441/muang_thai_vote_service/src/authorization"
	candidate_domain "github.com/pathai95441/muang_thai_vote_service/src/domain/candidate"
	user_domain "github.com/pathai95441/muang_thai_vote_service/src/domain/user"
	vote_domain "github.com/pathai95441/muang_thai_vote_service/src/domain/vote"
	vote_config_domain "github.com/pathai95441/muang_thai_vote_service/src/domain/vote_config"
	"github.com/pathai95441/muang_thai_vote_service/src/repositories/candidate"
	"github.com/pathai95441/muang_thai_vote_service/src/repositories/user"
	"github.com/pathai95441/muang_thai_vote_service/src/repositories/vote_config"
	"github.com/pathai95441/muang_thai_vote_service/src/repositories/vote_history"
	"github.com/pathai95441/muang_thai_vote_service/src/services/commands"
	"github.com/pathai95441/muang_thai_vote_service/src/services/queries"
)

type Commands struct {
	AddNewCandidate     commands.IAddNewCandidateHandler
	UpdateCandidateInfo commands.IUpdateCandidateInfoHandler
	CreateNewUser       commands.ICreateNewUserHandler
	VoteCandidate       commands.IVoteCandidateHandler
	DeleteCandidate     commands.IDeleteCandidateHandler
}

type Queries struct {
	GetAllCandidate queries.IGetAllCandidateQuery
}

type Domain struct {
	CandidateDomain  candidate_domain.ICandidateDomain
	UserDomain       user_domain.IUserDomain
	VoteDomain       vote_domain.IVoteDomain
	VoteConfigDomain vote_config_domain.IVoteConfigDomain
}

type Repositories struct {
	CandidateRepo   candidate.IRepository
	UserRepo        user.IRepository
	VoteHistoryRepo vote_history.IRepository
	VoteConfigRepo  vote_config.IRepository
}

type UserContext struct {
	UserID   string
	UserName string
	Email    string
	RoleID   int
}

type Application struct {
	Commands      Commands
	Queries       Queries
	Authorization auth.IAuthHandler
	UserContext   *UserContext
}
