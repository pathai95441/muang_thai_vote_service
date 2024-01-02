package application

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pathai95441/muang_thai_vote_service/src/consts"
	"github.com/pathai95441/muang_thai_vote_service/src/restapis/api_gen"
	"github.com/pathai95441/muang_thai_vote_service/src/services/commands"
)

// Candidate
func CreateNewCandidate(c echo.Context) error {
	b, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, api_gen.ErrorResult{
			Error: &api_gen.ErrorResultData{
				Code:    consts.InternalServerError,
				Message: err.Error(),
			},
		})
	}
	var r api_gen.CreateNewCandidate
	if err := json.Unmarshal(b, &r); err != nil {
		return c.JSON(http.StatusBadRequest, api_gen.ErrorResult{
			Error: &api_gen.ErrorResultData{
				Code:    consts.BadRequest,
				Message: err.Error(),
			},
		})
	}

	payload := commands.AddNewCandidateRequest{
		CandidateName:        r.CandidateName,
		CandidateDescription: r.CandidateDescription,
		CreateBy:             serverApp.UserContext.UserID,
	}

	err = validate.Struct(payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, api_gen.ErrorResult{
			Error: &api_gen.ErrorResultData{
				Code:    consts.BadRequest,
				Message: err.Error(),
			},
		})
	}

	err = serverApp.Commands.AddNewCandidate.Handle(context.Background(), payload)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, api_gen.ErrorResult{
			Error: &api_gen.ErrorResultData{
				Code:    consts.InternalServerError,
				Message: err.Error(),
			},
		})
	}

	return c.String(http.StatusOK, "Create Candidate Success")
}

func GetAllCandidate(c echo.Context) error {
	candidates, err := serverApp.Queries.GetAllCandidate.Handle(context.Background())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, api_gen.ErrorResult{
			Error: &api_gen.ErrorResultData{
				Code:    consts.InternalServerError,
				Message: err.Error(),
			},
		})
	}
	var response []api_gen.CandidatesResultData

	for _, data := range *candidates {
		response = append(response, api_gen.CandidatesResultData{
			CandidateDescription: data.CandidateDescription,
			CandidateID:          data.ID,
			CandidateName:        data.CandidateName,
			VoteScore:            float32(data.VoteScore),
		})
	}

	return c.JSON(http.StatusOK, api_gen.CandidatesResult{
		Data: &response,
	})
}

func UpdateCandidateInfo(c echo.Context) error {
	b, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, api_gen.ErrorResult{
			Error: &api_gen.ErrorResultData{
				Code:    consts.InternalServerError,
				Message: err.Error(),
			},
		})
	}
	var r api_gen.UpdateCandidate
	if err := json.Unmarshal(b, &r); err != nil {
		return c.JSON(http.StatusBadRequest, api_gen.ErrorResult{
			Error: &api_gen.ErrorResultData{
				Code:    consts.BadRequest,
				Message: err.Error(),
			},
		})
	}

	payload := commands.UpdateCandidateInfoRequest{
		CandidateID:          r.CandidateID,
		CandidateName:        r.CandidateName,
		CandidateDescription: r.CandidateDescription,
		UpdateBy:             serverApp.UserContext.UserID,
	}

	err = validate.Struct(payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, api_gen.ErrorResult{
			Error: &api_gen.ErrorResultData{
				Code:    consts.BadRequest,
				Message: err.Error(),
			},
		})
	}

	err = serverApp.Commands.UpdateCandidateInfo.Handle(context.Background(), payload)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, api_gen.ErrorResult{
			Error: &api_gen.ErrorResultData{
				Code:    consts.InternalServerError,
				Message: err.Error(),
			},
		})
	}

	return c.String(http.StatusOK, "Update Candidate info Success")
}

func DeleteCandidateByID(c echo.Context) error {
	candidateID := c.Param("candidateID")

	payload := commands.DeleteCandidateRequest{
		CandidateID: candidateID,
		DeletedBy:   serverApp.UserContext.UserID,
	}

	err := validate.Struct(payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, api_gen.ErrorResult{
			Error: &api_gen.ErrorResultData{
				Code:    consts.BadRequest,
				Message: err.Error(),
			},
		})
	}

	err = serverApp.Commands.DeleteCandidate.Handle(context.Background(), payload)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, api_gen.ErrorResult{
			Error: &api_gen.ErrorResultData{
				Code:    consts.InternalServerError,
				Message: err.Error(),
			},
		})
	}

	return c.String(http.StatusOK, "Deleted Candidate Success")
}

// User
func CreateNewUser(c echo.Context) error {
	b, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, api_gen.ErrorResult{
			Error: &api_gen.ErrorResultData{
				Code:    consts.InternalServerError,
				Message: err.Error(),
			},
		})
	}
	var r api_gen.CreateNewUser
	if err := json.Unmarshal(b, &r); err != nil {
		return c.JSON(http.StatusBadRequest, api_gen.ErrorResult{
			Error: &api_gen.ErrorResultData{
				Code:    consts.BadRequest,
				Message: err.Error(),
			},
		})
	}

	payload := commands.CreateNewUserRequest{
		UserName: r.UserName,
		Password: r.Password,
		Email:    r.Email,
		RoleID:   int(r.RoleID),
	}

	err = validate.Struct(payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, api_gen.ErrorResult{
			Error: &api_gen.ErrorResultData{
				Code:    consts.BadRequest,
				Message: err.Error(),
			},
		})
	}

	err = serverApp.Commands.CreateNewUser.Handle(context.Background(), payload)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, api_gen.ErrorResult{
			Error: &api_gen.ErrorResultData{
				Code:    consts.InternalServerError,
				Message: err.Error(),
			},
		})
	}

	return c.String(http.StatusOK, "Create User Success")
}

func SignIn(c echo.Context) error {
	b, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, api_gen.ErrorResult{
			Error: &api_gen.ErrorResultData{
				Code:    consts.InternalServerError,
				Message: err.Error(),
			},
		})
	}
	var r api_gen.SignIn
	if err := json.Unmarshal(b, &r); err != nil {
		return c.JSON(http.StatusBadRequest, api_gen.ErrorResult{
			Error: &api_gen.ErrorResultData{
				Code:    consts.BadRequest,
				Message: err.Error(),
			},
		})
	}

	token, err := serverApp.Authorization.SignIn(context.Background(), r.UserName, r.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, api_gen.ErrorResult{
			Error: &api_gen.ErrorResultData{
				Code:    consts.InternalServerError,
				Message: err.Error(),
			},
		})
	}

	return c.JSON(http.StatusInternalServerError, api_gen.SignInResult{
		Data: &api_gen.SignInResultData{
			Token: *token,
		},
	})
}

// vote
func VoteCandidate(c echo.Context) error {
	b, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, api_gen.ErrorResult{
			Error: &api_gen.ErrorResultData{
				Code:    consts.InternalServerError,
				Message: err.Error(),
			},
		})
	}
	var r api_gen.VoteCandidate
	if err := json.Unmarshal(b, &r); err != nil {
		return c.JSON(http.StatusBadRequest, api_gen.ErrorResult{
			Error: &api_gen.ErrorResultData{
				Code:    consts.BadRequest,
				Message: err.Error(),
			},
		})
	}
	
	payload := commands.VoteCandidateRequest{
		CandidateID: r.CandidateID,
		UserID:      serverApp.UserContext.UserID,
		UnVote:      r.UnVote,
	}

	err = validate.Struct(payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, api_gen.ErrorResult{
			Error: &api_gen.ErrorResultData{
				Code:    consts.BadRequest,
				Message: err.Error(),
			},
		})
	}

	err = serverApp.Commands.VoteCandidate.Handle(context.Background(), payload)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, api_gen.ErrorResult{
			Error: &api_gen.ErrorResultData{
				Code:    consts.InternalServerError,
				Message: err.Error(),
			},
		})
	}

	return c.String(http.StatusOK, "Vote Success")
}
