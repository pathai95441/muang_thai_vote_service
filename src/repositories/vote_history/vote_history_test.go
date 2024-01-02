package vote_history_test

import (
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/pathai95441/muang_thai_vote_service/src/repositories/vote_history"
)

var _ = Describe("VoteHistory", func() {
	It("NewVoteHistory", func() {
		result := vote_history.NewVoteHistory("mock-candidateID", "mock-UserID")
		Expect(result.ID).ShouldNot(BeNil())
		Expect(result.UserID).Should(Equal("mock-UserID"))
		Expect(result.CandidateID).Should(Equal("mock-candidateID"))
	})

	It("UnmarshallVoteHistoryFromDB", func() {
		ID := uuid.New().String()
		result := vote_history.UnmarshallVoteHistoryFromDB(ID, "mock-candidateID", "mock-UserID")
		Expect(result.ID).Should(Equal(ID))
		Expect(result.UserID).Should(Equal("mock-UserID"))
		Expect(result.CandidateID).Should(Equal("mock-candidateID"))
	})
})
