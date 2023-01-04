package files

import ports "sixels.io/manekani/core/ports/files"

type FilesService struct {
	ports.FilesRepository
}

func NewService(repo FilesRepository) *FilesService {
	return &FilesService{
		FilesRepository: repo,
	}
}
