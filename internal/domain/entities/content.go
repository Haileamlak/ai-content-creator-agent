package entities

import "time"

// Content represents a piece of content that will be created, managed, and posted by the agent.
type Content struct {
    ID              string    `json:"id" firestore:"id"`
    PostID          string    `json:"post_id" firestore:"post_id" binding:"required"` // ID of the post in the platform
    Message         string    `json:"message" firestore:"message" binding:"required"`
    ContentType     string    `json:"content_type" firestore:"content_type" binding:"required"` // E.g., Blog, Tweet, Facebook Post, etc.
    CreatedAt       time.Time `json:"created_at" firestore:"created_at" binding:"required"`
    Platform        string    `json:"platform" firestore:"platform" binding:"required"` // E.g., Twitter, Facebook, LinkedIn
    Status          string    `json:"status" firestore:"status" binding:"required"` // E.g., Draft, Scheduled, Posted
    CompanyID       string    `json:"company_id" firestore:"company_id" binding:"required"` // ID of the Company this content belongs to
    PositiveComment int       `json:"positive_comment" firestore:"positive_comment" binding:"required"` // Number of positive Comment received
    NegativeComment int       `json:"negative_comment" firestore:"negative_comment" binding:"required"` // Number of negative Comment received
    NeutralComment  int       `json:"neutral_comment" firestore:"neutral_comment" binding:"required"` // Number of neutral Comment received
    Likes           int       `json:"likes" firestore:"likes" binding:"required"` // Number of likes received
    Shares          int       `json:"shares" firestore:"shares" binding:"required"` // Number of shares received
}
