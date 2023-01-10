package files

import (
	"io"
	"os"
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

	FileInfo struct {
		Size      int64
		Name      string
		Kind      string
		Namespace string
	}
)
