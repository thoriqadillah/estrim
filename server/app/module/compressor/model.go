package compressor

type Type = string

const (
	Image Type = "image"
	Video Type = "video"
	PDF   Type = "pdf"
)

type CreateFile struct {
	UserId string
	Path   string
	Name   string
	Type   Type
	Mime   string
	Size   int64
}

type UpdateFile struct {
	Path       *string `json:"path"`
	Compressed *bool   `json:"compressed"`
	Size       *int64  `json:"size"`
}
