// Package schema provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.13.4 DO NOT EDIT.
package schema

import (
	"time"

	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
)

// Comment defines model for Comment.
type Comment struct {
	Color       string    `json:"color"`
	CreatedAt   time.Time `json:"createdAt"`
	Id          CommentId `json:"id"`
	IsAnonymous bool      `json:"isAnonymous"`
	MeetingId   MeetingId `json:"meetingId"`
	Text        string    `json:"text"`
	Username    Username  `json:"username"`
}

// CommentId defines model for CommentId.
type CommentId = openapi_types.UUID

// CreateMeetingRequest defines model for CreateMeetingRequest.
type CreateMeetingRequest struct {
	Description string `json:"description"`
	VideoId     string `json:"videoId"`
}

// CreateToken defines model for CreateToken.
type CreateToken struct {
	Description string    `json:"description"`
	ExpireAt    time.Time `json:"expireAt"`
	MeetingId   MeetingId `json:"meetingId"`
	Username    Username  `json:"username"`
}

// Meeting defines model for Meeting.
type Meeting struct {
	Description string    `json:"description"`
	EndedAt     time.Time `json:"endedAt"`
	Id          MeetingId `json:"id"`
	StartedAt   time.Time `json:"startedAt"`
	Thumbnail   string    `json:"thumbnail"`
	Title       string    `json:"title"`
	VideoId     string    `json:"videoId"`
}

// MeetingId defines model for MeetingId.
type MeetingId = openapi_types.UUID

// MeetingsWithTotal defines model for MeetingsWithTotal.
type MeetingsWithTotal struct {
	Meetings []Meeting `json:"meetings"`
	Total    int       `json:"total"`
}

// Reaction defines model for Reaction.
type Reaction struct {
	CreatedAt time.Time  `json:"createdAt"`
	Id        ReactionId `json:"id"`
	MeetingId MeetingId  `json:"meetingId"`
	StampId   string     `json:"stampId"`
	Username  Username   `json:"username"`
}

// ReactionId defines model for ReactionId.
type ReactionId = openapi_types.UUID

// Token defines model for Token.
type Token struct {
	CreatedAt   time.Time `json:"createdAt"`
	CreatorId   string    `json:"creatorId"`
	Description string    `json:"description"`
	ExpireAt    time.Time `json:"expireAt"`
	MeetingId   MeetingId `json:"meetingId"`
	Token       string    `json:"token"`
	Username    Username  `json:"username"`
}

// Tokens defines model for Tokens.
type Tokens struct {
	Meeting Meeting `json:"meeting"`
	Token   Token   `json:"token"`
}

// Username defines model for Username.
type Username = string

// LimitInQuery defines model for limitInQuery.
type LimitInQuery = int

// MeetingIdInPath defines model for meetingIdInPath.
type MeetingIdInPath = MeetingId

// OffsetInQuery defines model for offsetInQuery.
type OffsetInQuery = int

// SessionsInCookie defines model for sessionsInCookie.
type SessionsInCookie = string

// TokenInPath defines model for tokenInPath.
type TokenInPath = string

// GetMeetingsParams defines parameters for GetMeetings.
type GetMeetingsParams struct {
	// Limit 取得する件数
	Limit *LimitInQuery `form:"limit,omitempty" json:"limit,omitempty" query:"limit"`

	// Offset 取得を開始する件数（0-index）
	Offset *OffsetInQuery `form:"offset,omitempty" json:"offset,omitempty" query:"offset"`
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

// UpdateTokenJSONRequestBody defines body for UpdateToken for application/json ContentType.
type UpdateTokenJSONRequestBody = Token
