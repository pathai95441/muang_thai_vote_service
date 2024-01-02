package user_test

import (
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/pathai95441/muang_thai_vote_service/src/repositories/user"
)

var _ = Describe("User", func() {
	It("NewUserInfo", func() {
		result := user.NewUserInfo("Pathai", "Test", "test@gmail.com", 1, nil)
		Expect(result.ID).ShouldNot(BeNil())
		Expect(result.UserName).Should(Equal("Pathai"))
		Expect(result.Password).Should(Equal("Test"))
		Expect(result.Email).Should(Equal("test@gmail.com"))
		Expect(result.RoleID).Should(Equal(1))
		Expect(result.VoteCandidateID).Should(BeNil())
	})

	It("UnmarshallUserInfoFromDB", func() {
		userID := uuid.New().String()
		result := user.UnmarshallUserInfoFromDB(userID, "Pathai", "Test", "test@gmail.com", 1, nil)
		Expect(result.ID).Should(Equal(userID))
		Expect(result.UserName).Should(Equal("Pathai"))
		Expect(result.Password).Should(Equal("Test"))
		Expect(result.Email).Should(Equal("test@gmail.com"))
		Expect(result.RoleID).Should(Equal(1))
		Expect(result.VoteCandidateID).Should(BeNil())
	})
})
