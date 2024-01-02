package candidate_test

import (
	"context"
	"database/sql"
	"regexp"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/pathai95441/muang_thai_vote_service/src/consts"
	"github.com/pathai95441/muang_thai_vote_service/src/repositories/candidate"
	"github.com/pathai95441/muang_thai_vote_service/src/repositories/db_models_gen"
)

var _ = Describe("Repository", func() {
	var (
		ctx    context.Context
		db     *sql.DB
		mockDB sqlmock.Sqlmock
		mockTx *sql.Tx
		repo   candidate.IRepository

		candidateID = uuid.New().String()

		CandidateName        = "mock-CandidateName"
		CandidateDescription = "mock-CandidateDescription"
		VoteScore            = int64(0)
		VoteScoreInt         = int(0)
		CreatedBy            = consts.ServiceName
	)
	BeforeEach(func() {
		ctx = context.Background()

		var err error
		db, mockDB, err = sqlmock.New()
		Expect(err).ShouldNot(HaveOccurred())
		repo = candidate.NewRepository(db)
	})

	Context("Get", func() {
		var (
			exec = regexp.QuoteMeta("SELECT `candidate`.* FROM `candidate` WHERE (`candidate`.`id` = ?) AND (`candidate`.`deleted_at` is null) LIMIT 1;")
		)

		BeforeEach(func() {

			var err error
			Expect(err).ShouldNot(HaveOccurred())
		})

		It("success", func() {

			rows := sqlmock.NewRows([]string{
				db_models_gen.CandidateColumns.ID,
				db_models_gen.CandidateColumns.CandidateName,
				db_models_gen.CandidateColumns.CandidateDescription,
				db_models_gen.CandidateColumns.VoteScore,
			}).
				AddRow(
					candidateID,
					CandidateName,
					CandidateDescription,
					VoteScore,
				)

			mockDB.ExpectQuery(exec).WithArgs(candidateID).WillReturnRows(rows)

			candidateInfo, err := repo.Get(ctx, candidateID)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(candidateInfo.ID).Should(Equal(candidateID))
			Expect(candidateInfo.CandidateName).Should(Equal(CandidateName))
			Expect(candidateInfo.CandidateDescription).Should(Equal(CandidateDescription))
			Expect(candidateInfo.VoteScore).Should(Equal(VoteScore))

		})

		It("error", func() {
			mockDB.ExpectQuery(exec).WithArgs(candidateID).WillReturnError(sql.ErrConnDone)
			userInfo, err := repo.GetAll(ctx, nil, nil)
			Expect(err).Should(HaveOccurred())
			Expect(userInfo).Should(BeNil())
		})
	})

	Context("GetAll", func() {
		var (
			exec = regexp.QuoteMeta("SELECT `candidate`.* FROM `candidate` WHERE (`candidate`.`deleted_at` is null);")
		)

		BeforeEach(func() {

			var err error
			Expect(err).ShouldNot(HaveOccurred())
		})

		It("success", func() {

			rows := sqlmock.NewRows([]string{
				db_models_gen.CandidateColumns.ID,
				db_models_gen.CandidateColumns.CandidateName,
				db_models_gen.CandidateColumns.CandidateDescription,
				db_models_gen.CandidateColumns.VoteScore,
			}).
				AddRow(
					candidateID,
					CandidateName,
					CandidateDescription,
					VoteScore,
				)

			mockDB.ExpectQuery(exec).WithArgs().WillReturnRows(rows)

			candidateInfo, err := repo.GetAll(ctx, nil, nil)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(len(*candidateInfo)).Should(Equal(1))

		})

		It("error", func() {
			mockDB.ExpectQuery(exec).WithArgs().WillReturnError(sql.ErrConnDone)
			candidateInfo, err := repo.GetAll(ctx, nil, nil)
			Expect(err).Should(HaveOccurred())
			Expect(candidateInfo).Should(BeNil())
		})
	})

	Context("Insert", func() {
		var (
			exec = regexp.QuoteMeta("INSERT INTO `candidate` (`id`,`candidate_name`,`candidate_description`,`vote_score`,`created_at`,`created_by`,`updated_at`,`updated_by`,`deleted_at`,`deleted_by`) VALUES (?,?,?,?,?,?,?,?,?,?)")
		)

		BeforeEach(func() {
			var err error
			Expect(err).ShouldNot(HaveOccurred())
		})

		It("success", func() {
			mockDB.ExpectExec(exec).WithArgs(
				candidateID, CandidateName, CandidateDescription, VoteScore, sqlmock.AnyArg(), CreatedBy, sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(),
			).WillReturnResult(sqlmock.NewResult(1, 1))

			err := repo.Insert(ctx, candidate.Candidate{
				ID:                   candidateID,
				CandidateName:        CandidateName,
				CandidateDescription: CandidateDescription,
				VoteScore:            VoteScore,
			}, CreatedBy)
			Expect(err).ShouldNot(HaveOccurred())
		})

		It("error", func() {
			mockDB.ExpectExec(exec).WithArgs(
				candidateID, CandidateName, CandidateDescription, VoteScore, sqlmock.AnyArg(), CreatedBy, sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(),
			).WillReturnError(sql.ErrConnDone)

			err := repo.Insert(ctx, candidate.Candidate{
				ID:                   candidateID,
				CandidateName:        CandidateName,
				CandidateDescription: CandidateDescription,
				VoteScore:            VoteScore,
			}, CreatedBy)
			Expect(err).Should(HaveOccurred())
		})
	})

	Context("Update", func() {
		var (
			exec = regexp.QuoteMeta("UPDATE `candidate` SET `updated_by`=?,`candidate_name`=?,`candidate_description`=?,`vote_score`=? WHERE `id`=?")
		)
		BeforeEach(func() {
			mockDB.ExpectBegin()

			var err error
			mockTx, err = db.BeginTx(ctx, nil)
			Expect(err).ShouldNot(HaveOccurred())
		})
		It("should update success", func() {
			mockDB.ExpectExec(exec).WithArgs(
				CreatedBy, CandidateName, CandidateDescription, VoteScoreInt, candidateID,
			).WillReturnResult(sqlmock.NewResult(1, 1))

			err := repo.Update(ctx, mockTx, candidateID, &CandidateName, &CandidateDescription, &VoteScoreInt, CreatedBy)

			Expect(err).ShouldNot(HaveOccurred())
		})
		It("should return err when db is error", func() {
			mockDB.ExpectExec(exec).WithArgs(
				CreatedBy, CandidateName, CandidateDescription, VoteScoreInt, candidateID,
			).WillReturnError(sql.ErrConnDone)

			err := repo.Update(ctx, mockTx, candidateID, &CandidateName, &CandidateDescription, &VoteScoreInt, CreatedBy)
			Expect(err).Should(HaveOccurred())
		})
	})

	Context("Delete", func() {
		var (
			exec = regexp.QuoteMeta("UPDATE `candidate` SET `deleted_at`=?,`deleted_by`=? WHERE `id`=?")
		)
		BeforeEach(func() {
			mockDB.ExpectBegin()

			var err error
			mockTx, err = db.BeginTx(ctx, nil)
			Expect(err).ShouldNot(HaveOccurred())
		})
		It("should update success", func() {
			mockDB.ExpectExec(exec).WithArgs(
				sqlmock.AnyArg(), CreatedBy, candidateID,
			).WillReturnResult(sqlmock.NewResult(1, 1))

			err := repo.Delete(ctx, mockTx, candidateID, CreatedBy)

			Expect(err).ShouldNot(HaveOccurred())
		})
		It("should return err when db is error", func() {
			mockDB.ExpectExec(exec).WithArgs(
				sqlmock.AnyArg(), CreatedBy, candidateID,
			).WillReturnError(sql.ErrConnDone)

			err := repo.Delete(ctx, mockTx, candidateID, CreatedBy)
			Expect(err).Should(HaveOccurred())
		})
	})
})
