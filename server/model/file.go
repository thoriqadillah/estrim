package model

type Type = string

const (
	Image Type = "image"
	Video Type = "video"
	PDF   Type = "pdf"
)

type File struct {
	Id         string `json:"id"`
	Path       string `json:"path"`
	Type       Type   `json:"type"`
	Mime       string `json:"mime"`
	Size       int64  `json:"size"`
	Compressed bool   `json:"compressed"`
}

type CreateFile struct {
	Path string `json:"path"`
	Type Type   `json:"type"`
	Mime string `json:"mime"`
	Size int64  `json:"size"`
}

type UpdateFile struct {
	Path       *string `json:"path"`
	Compressed *bool   `json:"compressed"`
	Size       *int64  `json:"size"`
}
