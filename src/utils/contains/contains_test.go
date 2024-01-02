package contains_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/pathai95441/muang_thai_vote_service/src/utils/contains"
)

var _ = Describe("contain_util", func() {

	Context("Should process correct", func() {
		It("case contains", func() {
			result1 := contains.ContainsElement([]int{1, 2, 3, 4, 5}, 1)
			Expect(result1).To(Equal(true))

			result2 := contains.ContainsElement([]int{1, 2, 3, 4, 5}, 2)
			Expect(result2).To(Equal(true))

			result3 := contains.ContainsElement([]int{1, 2, 3, 4, 5}, 3)
			Expect(result3).To(Equal(true))
		})

		It("case not contains", func() {
			result1 := contains.ContainsElement([]int{1, 2, 3, 4, 5}, 7)
			Expect(result1).To(Equal(false))

			result2 := contains.ContainsElement([]int{1, 2, 3, 4, 5}, 8)
			Expect(result2).To(Equal(false))

			result3 := contains.ContainsElement([]int{1, 2, 3, 4, 5}, 9)
			Expect(result3).To(Equal(false))
		})
	})
})
