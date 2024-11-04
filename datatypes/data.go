package datatypes

import "time"

// Root struct for User Activity
type UserActivity struct {
	ID        string    `json:"id"`
	Type      string    `json:"type"`
	Actor     Actor     `json:"actor"`
	Repo      Repo      `json:"repo"`
	Payload   Payload   `json:"payload"`
	Public    bool      `json:"public"`
	CreatedAt time.Time `json:"created_at"`
}

// Actor information
type Actor struct {
	ID           int    `json:"id"`
	Login        string `json:"login"`
	DisplayLogin string `json:"display_login"`
	GravatarID   string `json:"gravatar_id"`
	URL          string `json:"url"`
	AvatarURL    string `json:"avatar_url"`
}

// Repository details
type Repo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

// Payload contains the main activity data
type Payload struct {
	Action       string    `json:"action,omitempty"`
	Issue        *Issue    `json:"issue,omitempty"`
	Comment      *Comment  `json:"comment,omitempty"`
	RepositoryID int       `json:"repository_id,omitempty"`
	PushID       int       `json:"push_id,omitempty"`
	Size         int       `json:"size,omitempty"`
	DistinctSize int       `json:"distinct_size,omitempty"`
	Ref          string    `json:"ref,omitempty"`
	Head         string    `json:"head,omitempty"`
	Before       string    `json:"before,omitempty"`
	Commits      []Commit  `json:"commits,omitempty"`
}

// Issue information
type Issue struct {
	URL         string     `json:"url"`
	HTMLURL     string     `json:"html_url"`
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	User        User       `json:"user"`
	State       string     `json:"state"`
	Locked      bool       `json:"locked"`
	Comments    int        `json:"comments"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	ClosedAt    *time.Time `json:"closed_at,omitempty"`
	Body        string     `json:"body"`
	Reactions   Reactions  `json:"reactions"`
}

// User details
type User struct {
	Login string `json:"login"`
	ID    int    `json:"id"`
	URL   string `json:"url"`
}

type Reactions struct {
	URL       string `json:"url"`
	TotalCount int   `json:"total_count"`
	PlusOne    int   `json:"+1"`
	MinusOne   int   `json:"-1"`
}

// Comment on an issue
type Comment struct {
	URL     string    `json:"url"`
	HTMLURL string    `json:"html_url"`
	ID      int       `json:"id"`
	User    User      `json:"user"`
	Body    string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Reactions Reactions `json:"reactions"`
}

// Commit details for PushEvent
type Commit struct {
	SHA     string `json:"sha"`
	Author  Author `json:"author"`
	Message string `json:"message"`
	URL     string `json:"url"`
}

// Author of the commit
type Author struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}