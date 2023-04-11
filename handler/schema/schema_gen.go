// Package schema provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package schema

import (
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
)

// Comment defines model for Comment.
type Comment struct {
	Color       string             `json:"color"`
	CreatedAt   openapi_types.Date `json:"created_at"`
	Id          string             `json:"id"`
	IsAnonymous bool               `json:"is_anonymous"`
	MeetingId   string             `json:"meeting_id"`
	Text        string             `json:"text"`
	UserId      string             `json:"user_id"`
}

// CreateMeetingRequest defines model for CreateMeetingRequest.
type CreateMeetingRequest struct {
	Description string `json:"description"`
	VideoId     string `json:"video_id"`
}

// CreateToken defines model for CreateToken.
type CreateToken struct {
	Description string             `json:"description"`
	ExpireAt    openapi_types.Date `json:"expire_at"`
	MeetingId   float32            `json:"meeting_id"`
	UserId      string             `json:"user_id"`
}

// Meeting defines model for Meeting.
type Meeting struct {
	Description string             `json:"description"`
	EndedAt     openapi_types.Date `json:"ended_at"`
	Id          openapi_types.UUID `json:"id"`
	StartedAt   openapi_types.Date `json:"started_at"`
	Thumbnail   string             `json:"thumbnail"`
	Title       string             `json:"title"`
	VideoId     string             `json:"video_id"`
}

// MeetingsWithTotal defines model for MeetingsWithTotal.
type MeetingsWithTotal struct {
	Meetings []Meeting `json:"meetings"`
	Total    int       `json:"total"`
}

// Reaction defines model for Reaction.
type Reaction struct {
	CreatedAt openapi_types.Date `json:"created_at"`
	Id        string             `json:"id"`
	MeetingId string             `json:"meeting_id"`
	StampId   string             `json:"stamp_id"`
	UserId    string             `json:"user_id"`
}

// Token defines model for Token.
type Token struct {
	CreatedAt   openapi_types.Date `json:"created_at"`
	CreatorId   string             `json:"creator_id"`
	Description string             `json:"description"`
	ExpireAt    openapi_types.Date `json:"expire_at"`
	MeetingId   string             `json:"meeting_id"`
	Token       string             `json:"token"`
	UserId      string             `json:"user_id"`
}

// Tokens defines model for Tokens.
type Tokens struct {
	Meeting Meeting `json:"meeting"`
	Token   Token   `json:"token"`
}

// LimitInQuery defines model for limitInQuery.
type LimitInQuery = int

// MeetingIdInPath defines model for meetingIdInPath.
type MeetingIdInPath = openapi_types.UUID

// OffsetInQuery defines model for offsetInQuery.
type OffsetInQuery = int

// SessionsInCookie defines model for sessionsInCookie.
type SessionsInCookie = string

// TokenInPath defines model for tokenInPath.
type TokenInPath = string

// GetAllMeetingsParams defines parameters for GetAllMeetings.
type GetAllMeetingsParams struct {
	// Limit 取得する件数
	Limit *LimitInQuery `form:"limit,omitempty" json:"limit,omitempty"`

	// Offset 取得を開始する件数（0-index）
	Offset *OffsetInQuery `form:"offset,omitempty" json:"offset,omitempty"`
}

// CallbackParams defines parameters for Callback.
type CallbackParams struct {
	// Code OAuth2.0のcode
	Code     string            `form:"code" json:"code"`
	Sessions *SessionsInCookie `form:"sessions,omitempty" json:"sessions,omitempty"`
}

// CreateMeetingJSONRequestBody defines body for CreateMeeting for application/json ContentType.
type CreateMeetingJSONRequestBody = CreateMeetingRequest

// UpdateMeetingJSONRequestBody defines body for UpdateMeeting for application/json ContentType.
type UpdateMeetingJSONRequestBody = CreateMeetingRequest

// CreateTokenJSONRequestBody defines body for CreateToken for application/json ContentType.
type CreateTokenJSONRequestBody = CreateToken

// PatchTokenJSONRequestBody defines body for PatchToken for application/json ContentType.
type PatchTokenJSONRequestBody = Token
