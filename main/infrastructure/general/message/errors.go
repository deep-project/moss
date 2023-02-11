package message

import "errors"

var (
	ErrRecordNotFound         = errors.New("record not found")
	ErrIdRequired             = errors.New("id is required")
	ErrSlugRequired           = errors.New("slug is required")
	ErrNameRequired           = errors.New("name is required")
	ErrTitleRequired          = errors.New("title is required")
	ErrUrlRequired            = errors.New("url is required")
	ErrSlugStartSpaceRequired = errors.New("slug cannot start with a space")
	ErrSlugEndSpaceRequired   = errors.New("slug cannot end with a space")
	ErrSlugAlreadyExists      = errors.New("slug already exists")
	ErrIdAlreadyExists        = errors.New("id already exists")
	ErrNotFound               = errors.New("not found")
)
