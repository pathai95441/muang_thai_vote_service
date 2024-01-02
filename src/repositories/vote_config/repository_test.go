package vote_config_test

import (
	"context"
	"database/sql"
	"regexp"

	"github.com/DATA-DOG/go-sqlmock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/pathai95441/muang_thai_vote_service/src/repositories/db_models_gen"
	"github.com/pathai95441/muang_thai_vote_service/src/repositories/vote_config"
)

var _ = Describe("Repository", func() {
	var (
		ctx    context.Context
		db     *sql.DB
		mockDB sqlmock.Sqlmock
		repo   vote_config.IRepository

		ID      = 1
		isClose = false
	)
	BeforeEach(func() {
		ctx = context.Background()

		var err error
		db, mockDB, err = sqlmock.New()
		Expect(err).ShouldNot(HaveOccurred())
		repo = vote_config.NewRepository(db)
	})

	Context("Get", func() {
		var (
			exec = regexp.QuoteMeta("SELECT `vote_config`.* FROM `vote_config` WHERE (`vote_config`.`deleted_at` is null) LIMIT 1;")
		)

		BeforeEach(func() {

			var err error
			Expect(err).ShouldNot(HaveOccurred())
		})

		It("success", func() {

			rows := sqlmock.NewRows([]string{
				db_models_gen.VoteConfigColumns.ID,
				db_models_gen.VoteConfigColumns.VoteClose,
			}).
				AddRow(
					ID,
					0,
				)

			mockDB.ExpectQuery(exec).WithArgs().WillReturnRows(rows)

			isClosed, err := repo.Get(ctx)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(isClosed).Should(Equal(isClose))

		})

		It("error", func() {
			mockDB.ExpectQuery(exec).WithArgs().WillReturnError(sql.ErrConnDone)
			userInfo, err := repo.Get(ctx)
			Expect(err).Should(HaveOccurred())
			Expect(userInfo).Should(Equal(false))
		})
	})
})
