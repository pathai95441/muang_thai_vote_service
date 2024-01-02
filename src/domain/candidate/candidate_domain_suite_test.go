package candidate_domain_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestContact(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Candidate Domain Suite")
}

var (
	mockCtrl *gomock.Controller
	
)

var _ = BeforeEach(func() {
	mockCtrl = gomock.NewController(GinkgoT())

})

var _ = AfterEach(func() {
	mockCtrl.Finish()
})
