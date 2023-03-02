package files

import (
	"io"
	"os"
	"time"
)

type (
	CreateFileRequest struct {
		FileInfo
		Handle io.Reader
	}

	UpdateFileRequest struct {
		Handle os.File
	}

	ObjectWrapperResponse struct {
		io.ReadCloser
		Size        int64
		ContentType string
	}

	UploadURL struct {
		URL       string    `json:"url"`
		Resource  string    `json:"resource"`
		ExpiresAt time.Time `json:"expires_at"`
	}

	FileInfo struct {
		Size        int64
		Name        string
		Kind        string
		Namespace   string
		ContentType string
	}
)
