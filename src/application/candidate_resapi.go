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
		CreateBy:             "test",
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
		UpdateBy:             "test",
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
