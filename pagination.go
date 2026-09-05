package mailafrica

import "time"

// Pagination carries the standard page metadata.
type Pagination struct {
	Page       int   `json:"page"`
	PerPage    int   `json:"per_page"`
	Total      int64 `json:"total"`
	TotalPages int   `json:"total_pages"`
}

// ListOpts configures list endpoint pagination.
type ListOpts struct {
	Page    int
	PerPage int
}

func (o ListOpts) page() int {
	if o.Page <= 0 {
		return 1
	}
	return o.Page
}

func (o ListOpts) perPage() int {
	if o.PerPage <= 0 {
		return 25
	}
	if o.PerPage > 100 {
		return 100
	}
	return o.PerPage
}

func (o *ListOpts) applyDefaults() {
	if o == nil {
		return
	}
	if o.Page <= 0 {
		o.Page = 1
	}
	if o.PerPage <= 0 {
		o.PerPage = 25
	}
	if o.PerPage > 100 {
		o.PerPage = 100
	}
}

// Time is a custom time type that handles common timestamp formats.
type Time struct {
	time.Time
}
