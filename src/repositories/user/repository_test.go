package user_test

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
	"github.com/pathai95441/muang_thai_vote_service/src/repositories/user"
)

var _ = Describe("Repository", func() {
	var (
		ctx    context.Context
		db     *sql.DB
		mockDB sqlmock.Sqlmock
		repo   user.IRepository

		userID = uuid.New().String()

		UserName        = "mock-UserName"
		Password        = "mock-Password"
		Email           = "mock-Email"
		VoteCandidateID = "mock-VoteCandidateID"
		RoleID          = 1
		CreatedBy       = consts.ServiceName
	)
	BeforeEach(func() {
		ctx = context.Background()

		var err error
		db, mockDB, err = sqlmock.New()
		Expect(err).ShouldNot(HaveOccurred())
		repo = user.NewRepository(db)
	})

	Context("Get", func() {
		var (
			exec = regexp.QuoteMeta("SELECT `users`.* FROM `users` WHERE (`users`.`id` = ?) AND (`users`.`deleted_at` is null) LIMIT 1;")
		)

		BeforeEach(func() {

			var err error
			Expect(err).ShouldNot(HaveOccurred())
		})

		It("success", func() {

			rows := sqlmock.NewRows([]string{
				db_models_gen.UserColumns.ID,
				db_models_gen.UserColumns.UserName,
				db_models_gen.UserColumns.Password,
				db_models_gen.UserColumns.Email,
				db_models_gen.UserColumns.RoleID,
				db_models_gen.UserColumns.VoteCandidateID,
			}).
				AddRow(
					userID,
					UserName,
					Password,
					Email,
					RoleID,
					VoteCandidateID,
				)

			mockDB.ExpectQuery(exec).WithArgs(userID).WillReturnRows(rows)

			userInfo, err := repo.Get(ctx, userID)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(userInfo.ID).Should(Equal(userID))
			Expect(userInfo.Password).Should(Equal(Password))
			Expect(userInfo.UserName).Should(Equal(UserName))
			Expect(userInfo.Email).Should(Equal(Email))
			Expect(userInfo.RoleID).Should(Equal(RoleID))
			Expect(userInfo.VoteCandidateID).Should(Equal(&VoteCandidateID))

		})

		It("error", func() {

			mockDB.ExpectQuery(exec).WithArgs(userID).WillReturnError(sql.ErrConnDone)
			userInfo, err := repo.Get(ctx, userID)
			Expect(err).Should(HaveOccurred())
			Expect(userInfo).Should(BeNil())
		})
	})

	Context("GetByUserName", func() {
		var (
			exec = regexp.QuoteMeta("SELECT `users`.* FROM `users` WHERE (`users`.`user_name` = ?) AND (`users`.`deleted_at` is null) LIMIT 1;")
		)

		BeforeEach(func() {

			var err error
			Expect(err).ShouldNot(HaveOccurred())
		})

		It("success", func() {

			rows := sqlmock.NewRows([]string{
				db_models_gen.UserColumns.ID,
				db_models_gen.UserColumns.UserName,
				db_models_gen.UserColumns.Password,
				db_models_gen.UserColumns.Email,
				db_models_gen.UserColumns.RoleID,
				db_models_gen.UserColumns.VoteCandidateID,
			}).
				AddRow(
					userID,
					UserName,
					Password,
					Email,
					RoleID,
					VoteCandidateID,
				)

			mockDB.ExpectQuery(exec).WithArgs(UserName).WillReturnRows(rows)

			userInfo, err := repo.GetByUserName(ctx, UserName)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(userInfo.ID).Should(Equal(userID))
			Expect(userInfo.Password).Should(Equal(Password))
			Expect(userInfo.UserName).Should(Equal(UserName))
			Expect(userInfo.Email).Should(Equal(Email))
			Expect(userInfo.RoleID).Should(Equal(RoleID))
			Expect(userInfo.VoteCandidateID).Should(Equal(&VoteCandidateID))

		})

		It("error", func() {

			mockDB.ExpectQuery(exec).WithArgs(userID).WillReturnError(sql.ErrConnDone)
			userInfo, err := repo.Get(ctx, userID)
			Expect(err).Should(HaveOccurred())
			Expect(userInfo).Should(BeNil())
		})
	})

	Context("Insert", func() {
		var (
			exec = regexp.QuoteMeta("NSERT INTO `users` (`id`,`user_name`,`password`,`email`,`role_id`,`created_at`,`created_by`,`updated_at`,`updated_by`,`deleted_at`,`deleted_by`,`vote_candidate_id`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?)")
		)

		BeforeEach(func() {
			var err error
			Expect(err).ShouldNot(HaveOccurred())
		})

		It("success", func() {
			mockDB.ExpectExec(exec).WithArgs(
				sqlmock.AnyArg(), UserName, Password, Email, RoleID, sqlmock.AnyArg(), CreatedBy, sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), nil,
			).WillReturnResult(sqlmock.NewResult(1, 1))

			err := repo.Insert(ctx, user.UserInfo{
				ID:       userID,
				UserName: UserName,
				Password: Password,
				Email:    Email,
				RoleID:   RoleID,
			})
			Expect(err).ShouldNot(HaveOccurred())
		})

		It("error", func() {
			mockDB.ExpectExec(exec).WithArgs(
				sqlmock.AnyArg(), UserName, Password, Email, RoleID, sqlmock.AnyArg(), CreatedBy, sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), nil,
			).WillReturnError(sql.ErrConnDone)

			err := repo.Insert(ctx, user.UserInfo{
				ID:       userID,
				UserName: UserName,
				Password: Password,
				Email:    Email,
				RoleID:   RoleID,
			})
			Expect(err).Should(HaveOccurred())
		})
	})
})
