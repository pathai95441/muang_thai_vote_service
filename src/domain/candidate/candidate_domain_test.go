package candidate_domain_test

import (
	"context"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/pathai95441/muang_thai_vote_service/src/consts"
	candidate_domain "github.com/pathai95441/muang_thai_vote_service/src/domain/candidate"
	"github.com/pathai95441/muang_thai_vote_service/src/repositories/candidate"
	mock_candidate "github.com/pathai95441/muang_thai_vote_service/src/repositories/candidate/mock"
)

var _ = Describe("Repository", func() {
	var (
		ctx                  context.Context
		mockCandidateRepo    *mock_candidate.MockIRepository
		candidateDomain      candidate_domain.ICandidateDomain
		CandidateName        = "CandidateName-mock"
		CandidateDescription = "CandidateDescription-mock"
		VoteScore            = int64(0)
	)
	BeforeEach(func() {
		ctx = context.Background()
		mockCandidateRepo = mock_candidate.NewMockIRepository(mockCtrl)
		candidateDomain = candidate_domain.NewCandidateDomain(mockCandidateRepo)
	})

	Context("Success", func() {
		It("should return nil when create candidate success", func() {
			mockCandidateRepo.EXPECT().Insert(ctx, gomock.Any(), consts.ServiceName).DoAndReturn(func(ctx context.Context, candidate candidate.Candidate, createBy string) error {
				Expect(candidate.CandidateDescription).Should(Equal(CandidateDescription))
				Expect(candidate.CandidateName).Should(Equal(CandidateName))
				Expect(candidate.ID).ShouldNot(BeNil())
				Expect(candidate.VoteScore).Should(Equal(VoteScore))
				return nil
			})

			err := candidateDomain.AddNewCandidate(ctx, CandidateName, CandidateDescription, consts.ServiceName)
			Expect(err).ShouldNot(HaveOccurred())
		})

		It("should return data when GetAllCandidate success", func() {
			var list []candidate.Candidate
			list = append(list, candidate.Candidate{
				ID:                   uuid.NewString(),
				CandidateName:        CandidateName,
				CandidateDescription: CandidateDescription,
				VoteScore:            VoteScore,
			})
			mockCandidateRepo.EXPECT().GetAll(ctx, nil, nil).Return(&list, nil)

			listData, err := candidateDomain.GetAllCandidate(ctx, nil, nil)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(len(*listData)).Should(Equal(1))
		})

		// It("should return nil when update candidateInfo success", func() {
		// 	mockCandidateRepo.EXPECT().Update(ctx, nil, consts.ServiceName).DoAndReturn(func(ctx context.Context, candidate candidate.Candidate, createBy string) error {
		// 		Expect(candidate.CandidateDescription).Should(Equal(CandidateDescription))
		// 		Expect(candidate.CandidateName).Should(Equal(CandidateName))
		// 		Expect(candidate.ID).ShouldNot(BeNil())
		// 		Expect(candidate.VoteScore).Should(Equal(VoteScore))
		// 		return nil
		// 	})

		// 	err := candidateDomain.AddNewCandidate(ctx, CandidateName, CandidateDescription, consts.ServiceName)
		// 	Expect(err).ShouldNot(HaveOccurred())
		// })
	})
})
