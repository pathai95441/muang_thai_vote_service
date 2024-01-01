package candidate

import (
	"github.com/google/uuid"
)

type Candidate struct {
	ID                   string
	CandidateName        string
	CandidateDescription string
	VoteScore            int64
}

func NewCandidate(candidateName string, candidateDescription string) Candidate {
	return Candidate{
		ID:                   uuid.NewString(),
		CandidateName:        candidateName,
		CandidateDescription: candidateDescription,
		VoteScore:            0,
	}
}

func UnmarshallCandidateFromDB(
	ID string,
	CandidateName string,
	CandidateDescription string,
	VoteScore int64,
) Candidate {
	return Candidate{
		ID,
		CandidateName,
		CandidateDescription,
		VoteScore,
	}
}
