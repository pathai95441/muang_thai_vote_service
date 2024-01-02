package candidate_test

import (
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/pathai95441/muang_thai_vote_service/src/repositories/candidate"
)

var _ = Describe("Candidate", func() {
	It("NewCandidate", func() {
		result := candidate.NewCandidate("Pathai", "Software Engineer")
		Expect(result.ID).ShouldNot(BeNil())
		Expect(result.CandidateName).Should(Equal("Pathai"))
		Expect(result.CandidateDescription).Should(Equal("Software Engineer"))
		Expect(result.VoteScore).Should(Equal(int64(0)))
	})

	It("UnmarshallCandidateFromDB", func() {
		ID := uuid.New().String()
		result := candidate.UnmarshallCandidateFromDB(ID, "Pathai", "Software Engineer", int64(1))
		Expect(result.ID).Should(Equal(ID))
		Expect(result.CandidateName).Should(Equal("Pathai"))
		Expect(result.CandidateDescription).Should(Equal("Software Engineer"))
		Expect(result.VoteScore).Should(Equal(int64(1)))
	})
})
