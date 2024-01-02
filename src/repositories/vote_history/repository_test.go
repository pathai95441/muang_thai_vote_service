package vote_history_test

import (
	"context"
	"database/sql"
	"regexp"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/pathai95441/muang_thai_vote_service/src/consts"
	"github.com/pathai95441/muang_thai_vote_service/src/repositories/db_models_gen"
	"github.com/pathai95441/muang_thai_vote_service/src/repositories/vote_history"
)

var _ = Describe("Repository", func() {
	var (
		ctx    context.Context
		db     *sql.DB
		mockDB sqlmock.Sqlmock
		mockTx *sql.Tx
		repo   vote_history.IRepository

		historyID   = uuid.New().String()
		userID      = uuid.New().String()
		candidateID = uuid.New().String()
		CreatedBy   = consts.ServiceName
	)
	BeforeEach(func() {
		ctx = context.Background()

		var err error
		db, mockDB, err = sqlmock.New()
		Expect(err).ShouldNot(HaveOccurred())
		repo = vote_history.NewRepository(db)
	})

	Context("GetByUserID", func() {
		var (
			exec = regexp.QuoteMeta("SELECT `vote_history`.* FROM `vote_history` WHERE (`vote_history`.`user_id` = ?) AND (`vote_history`.`deleted_at` is null) LIMIT 1;")
		)

		BeforeEach(func() {

			var err error
			Expect(err).ShouldNot(HaveOccurred())
		})

		It("success", func() {

			rows := sqlmock.NewRows([]string{
				db_models_gen.VoteHistoryColumns.ID,
				db_models_gen.VoteHistoryColumns.CandidateID,
				db_models_gen.VoteHistoryColumns.UserID,
			}).
				AddRow(
					historyID,
					candidateID,
					userID,
				)

			mockDB.ExpectQuery(exec).WithArgs(userID).WillReturnRows(rows)

			history, err := repo.GetByUserID(ctx, userID)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(history.ID).Should(Equal(historyID))
			Expect(history.CandidateID).Should(Equal(candidateID))
			Expect(history.UserID).Should(Equal(userID))

		})

		It("error", func() {
			mockDB.ExpectQuery(exec).WithArgs(userID).WillReturnError(sql.ErrConnDone)
			history, err := repo.GetByUserID(ctx, userID)
			Expect(err).Should(HaveOccurred())
			Expect(history).Should(BeNil())
		})
	})

	Context("GetUserVote", func() {
		var (
			exec = regexp.QuoteMeta("SELECT `vote_history`.* FROM `vote_history` WHERE (`vote_history`.`user_id` = ?) AND (`vote_history`.`candidate_id` = ?) AND (`vote_history`.`deleted_at` is null) LIMIT 1;")
		)

		BeforeEach(func() {

			var err error

			Expect(err).ShouldNot(HaveOccurred())
		})

		It("success", func() {

			rows := sqlmock.NewRows([]string{
				db_models_gen.VoteHistoryColumns.ID,
				db_models_gen.VoteHistoryColumns.CandidateID,
				db_models_gen.VoteHistoryColumns.UserID,
			}).
				AddRow(
					historyID,
					candidateID,
					userID,
				)

			mockDB.ExpectQuery(exec).WithArgs(userID, candidateID).WillReturnRows(rows)

			history, err := repo.GetUserVote(ctx, userID, candidateID)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(history.ID).Should(Equal(historyID))
			Expect(history.CandidateID).Should(Equal(candidateID))
			Expect(history.UserID).Should(Equal(userID))

		})

		It("error", func() {
			mockDB.ExpectQuery(exec).WithArgs(userID, candidateID).WillReturnError(sql.ErrConnDone)
			history, err := repo.GetUserVote(ctx, userID, candidateID)
			Expect(err).Should(HaveOccurred())
			Expect(history).Should(BeNil())
		})
	})

	Context("Insert", func() {
		var (
			exec = regexp.QuoteMeta("INSERT INTO `vote_history` (`id`,`candidate_id`,`user_id`,`created_at`,`created_by`,`updated_at`,`updated_by`,`deleted_at`,`deleted_by`) VALUES (?,?,?,?,?,?,?,?,?)")
		)

		BeforeEach(func() {
			mockDB.ExpectBegin()

			var err error
			mockTx, err = db.BeginTx(ctx, nil)
			Expect(err).ShouldNot(HaveOccurred())
		})

		It("success", func() {
			mockDB.ExpectExec(exec).WithArgs(
				historyID, candidateID, userID, sqlmock.AnyArg(), CreatedBy, sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(),
			).WillReturnResult(sqlmock.NewResult(1, 1))

			err := repo.Insert(ctx, mockTx, vote_history.VoteHistory{
				ID:          historyID,
				CandidateID: candidateID,
				UserID:      userID,
			})
			Expect(err).ShouldNot(HaveOccurred())
		})

		It("error", func() {
			mockDB.ExpectExec(exec).WithArgs(
				historyID, candidateID, userID,
			).WillReturnError(sql.ErrConnDone)

			err := repo.Insert(ctx, mockTx, vote_history.VoteHistory{
				ID:          historyID,
				CandidateID: candidateID,
				UserID:      userID,
			})
			Expect(err).Should(HaveOccurred())
		})
	})

	Context("Delete", func() {
		var (
			exec = regexp.QuoteMeta("UPDATE `vote_history` SET `updated_by`=?,`deleted_at`=?,`deleted_by`=? WHERE `id`=?")
		)
		BeforeEach(func() {
			mockDB.ExpectBegin()

			var err error
			mockTx, err = db.BeginTx(ctx, nil)
			Expect(err).ShouldNot(HaveOccurred())
		})
		It("should update success", func() {
			mockDB.ExpectExec(exec).WithArgs(
				CreatedBy, sqlmock.AnyArg(), CreatedBy, historyID,
			).WillReturnResult(sqlmock.NewResult(1, 1))

			err := repo.DeletedVoted(ctx, mockTx, historyID)

			Expect(err).ShouldNot(HaveOccurred())
		})
		It("should return err when db is error", func() {
			mockDB.ExpectExec(exec).WithArgs(
				CreatedBy, sqlmock.AnyArg(), CreatedBy, historyID,
			).WillReturnError(sql.ErrConnDone)

			err := repo.DeletedVoted(ctx, mockTx, historyID)
			Expect(err).Should(HaveOccurred())
		})
	})
})
