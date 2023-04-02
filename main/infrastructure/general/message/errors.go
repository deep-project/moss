package message

import "errors"

var (
	ErrNotFound       = errors.New("not found")
	ErrRecordNotFound = errors.New("record not found")

	ErrIdRequired      = errors.New("id is required")
	ErrSlugRequired    = errors.New("slug is required")
	ErrNameRequired    = errors.New("name is required")
	ErrTitleRequired   = errors.New("title is required")
	ErrContentRequired = errors.New("content is required")
	ErrUrlRequired     = errors.New("url is required")

	ErrSlugStartSpaceRequired = errors.New("slug cannot start with a space")
	ErrSlugEndSpaceRequired   = errors.New("slug cannot end with a space")

	ErrIdAlreadyExists    = errors.New("id already exists")
	ErrSlugAlreadyExists  = errors.New("slug already exists")
	ErrTitleAlreadyExists = errors.New("title already exists")
)
