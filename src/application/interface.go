package application

import (
	candidate_domain "github.com/pathai95441/muang_thai_vote_service/src/domain/candidate"
	"github.com/pathai95441/muang_thai_vote_service/src/repositories/candidate"
	"github.com/pathai95441/muang_thai_vote_service/src/services/commands"
	"github.com/pathai95441/muang_thai_vote_service/src/services/queries"
)

type Commands struct {
	AddNewCandidate     commands.IAddNewCandidateHandler
	UpdateCandidateInfo commands.IUpdateCandidateInfoHandler
}

type Queries struct {
	GetAllCandidate queries.IGetAllCandidateQuery
}

type Domain struct {
	CandidateDomain candidate_domain.CandidateDomain
}

type Repositories struct {
	CandidateRepo candidate.IRepository
}

type Application struct {
	Commands Commands
	Queries  Queries
}
