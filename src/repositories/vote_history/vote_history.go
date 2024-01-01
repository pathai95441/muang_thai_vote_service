package vote_history

import "github.com/google/uuid"

type VoteHistory struct {
	ID          string
	CandidateID string
	UserID      string
}

func NewVoteHistory(
	CandidateID string,
	UserID string,
) VoteHistory {
	return VoteHistory{
		ID:          uuid.NewString(),
		CandidateID: CandidateID,
		UserID:      UserID,
	}
}

func UnmarshallVoteHistoryFromDB(
	ID string,
	CandidateID string,
	UserID string,
) VoteHistory {
	return VoteHistory{
		ID:          ID,
		CandidateID: CandidateID,
		UserID:      UserID,
	}
}
