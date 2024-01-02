// Package api_gen provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.11.0 DO NOT EDIT.
package api_gen

// CandidatesResultData defines model for CandidatesResultData.
type CandidatesResultData struct {
	CandidateDescription string  `json:"candidateDescription"`
	CandidateID          string  `json:"candidateID"`
	CandidateName        string  `json:"candidateName"`
	VoteScore            float32 `json:"voteScore"`
}

// ErrorResultData defines model for ErrorResultData.
type ErrorResultData struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// SignInResultData defines model for SignInResultData.
type SignInResultData struct {
	Token string `json:"token"`
}

// CandidatesResult defines model for CandidatesResult.
type CandidatesResult struct {
	Data *[]CandidatesResultData `json:"data,omitempty"`
}

// ErrorResult defines model for ErrorResult.
type ErrorResult struct {
	Error *ErrorResultData `json:"error,omitempty"`
}

// SignInResult defines model for SignInResult.
type SignInResult struct {
	Data *SignInResultData `json:"data,omitempty"`
}

// CreateNewCandidate defines model for CreateNewCandidate.
type CreateNewCandidate struct {
	CandidateDescription string `json:"candidateDescription"`
	CandidateName        string `json:"candidateName"`
}

// CreateNewUser defines model for CreateNewUser.
type CreateNewUser struct {
	Email    string  `json:"email"`
	Password string  `json:"password"`
	RoleID   float32 `json:"roleID"`
	UserName string  `json:"userName"`
}

// SignIn defines model for SignIn.
type SignIn struct {
	Password string `json:"password"`
	UserName string `json:"userName"`
}

// UpdateCandidate defines model for UpdateCandidate.
type UpdateCandidate struct {
	CandidateDescription *string `json:"candidateDescription,omitempty"`
	CandidateID          string  `json:"candidateID"`
	CandidateName        *string `json:"candidateName,omitempty"`
}

// VoteCandidate defines model for VoteCandidate.
type VoteCandidate struct {
	CandidateID string `json:"candidateID"`
	UnVote      bool   `json:"unVote"`
}

// CreateNewCandidateJSONRequestBody defines body for CreateNewCandidate for application/json ContentType.
type CreateNewCandidateJSONRequestBody CreateNewCandidate

// UpdateCandidateInfoJSONRequestBody defines body for UpdateCandidateInfo for application/json ContentType.
type UpdateCandidateInfoJSONRequestBody UpdateCandidate

// SignInJSONRequestBody defines body for SignIn for application/json ContentType.
type SignInJSONRequestBody SignIn

// CreateNewUserJSONRequestBody defines body for CreateNewUser for application/json ContentType.
type CreateNewUserJSONRequestBody CreateNewUser

// VoteCandidateJSONRequestBody defines body for VoteCandidate for application/json ContentType.
type VoteCandidateJSONRequestBody VoteCandidate
