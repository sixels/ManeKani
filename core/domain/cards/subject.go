package cards

import (
	"time"

	"github.com/google/uuid"
	"github.com/sixels/manekani/core/domain/cards/filters"
)

type (
	Subject struct {
		ID        uuid.UUID `json:"id"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`

		Kind  string `json:"kind"`
		Level int32  `json:"level"`

		// e.g. "ground", "一", nil, "一", 2
		Name       string         `json:"name"`
		Value      *string        `json:"value"`
		ValueImage *RemoteContent `json:"value_image"`
		Slug       string         `json:"slug"`
		Priority   uint8          `json:"priority"`

		Resources           *map[string][]RemoteContent `json:"resources"`
		StudyData           []StudyData                 `json:"study_data"`
		AdditionalStudyData *map[string]any             `json:"additional_study_data"`

		Dependencies []uuid.UUID `json:"dependencies"`
		Dependents   []uuid.UUID `json:"dependents"`
		Similars     []uuid.UUID `json:"similars"`
		Deck         uuid.UUID   `json:"deck"`

		Owner string `json:"owner"`
	}

	PartialSubject struct {
		ID uuid.UUID `json:"id"`

		Kind  string `json:"kind"`
		Level int32  `json:"level"`

		Name       string         `json:"name"`
		Value      *string        `json:"value"`
		ValueImage *RemoteContent `json:"value_image"`
		Slug       string         `json:"slug"`
		Priority   uint8          `json:"priority"`

		StudyData []StudyData `json:"study_data"`

		Dependencies []uuid.UUID `json:"dependencies"`
		Dependents   []uuid.UUID `json:"dependents"`
		Similars     []uuid.UUID `json:"similars"`
		Deck         uuid.UUID   `json:"deck"`

		Owner string `json:"owner"`
	}

	MinimalSubject struct {
		ID uuid.UUID `json:"id"`

		Kind  string `json:"kind"`
		Level int32  `json:"level"`

		Name       string         `json:"name"`
		Value      *string        `json:"value"`
		ValueImage *RemoteContent `json:"value_image"`
		Slug       string         `json:"slug"`
		Priority   uint8          `json:"priority"`

		Deck uuid.UUID `json:"deck"`
	}

	CreateSubjectRequest struct {
		Kind  string `json:"kind,omitempty" form:"kind" binding:"required"`
		Level int32  `json:"level,omitempty" form:"level" binding:"required"`

		Name       string         `json:"name" form:"name" binding:"required"`
		Value      *string        `json:"value,omitempty" form:"value" binding:"-"`
		ValueImage *RemoteContent `json:"value_image,omitempty" form:"value_image" binding:"-"`
		Slug       string         `json:"slug" form:"slug" binding:"required"`
		Priority   uint8          `json:"priority" form:"priority"`

		StudyData           []StudyData                 `json:"study_data,omitempty" form:"study_data" binding:"-"`
		Resources           *map[string][]RemoteContent `json:"resources,omitempty" form:"resources" binding:"-"`
		AdditionalStudyData *map[string]any             `json:"additional_study_data" form:"additional_study_data" binding:"-"`

		Dependencies []uuid.UUID `json:"dependencies,omitempty" form:"dependencies"`
		Dependents   []uuid.UUID `json:"dependents,omitempty" form:"dependents"`
		Similars     []uuid.UUID `json:"similars,omitempty" form:"similars"`

		Deck uuid.UUID `json:"deck" form:"deck" binding:"required"`
	}

	UpdateSubjectRequest struct {
		Kind  *string `json:"kind,omitempty" form:"kind"`
		Level *int32  `json:"level,omitempty" form:"level"`

		Name       *string        `json:"name,omitempty" form:"name"`
		Value      *string        `json:"value,omitempty" form:"value"`
		ValueImage *RemoteContent `json:"value_image,omitempty" form:"value_image"`
		Slug       *string        `json:"slug,omitempty" form:"slug"`
		Priority   *uint8         `json:"priority,omitempty" form:"priority"`

		StudyData           *[]StudyData                `json:"study_data,omitempty" form:"study_data"`
		Resources           *map[string][]RemoteContent `json:"resources,omitempty" form:"resources"`
		AdditionalStudyData *map[string]any             `json:"additional_study_data" form:"additional_study_data" binding:"-"`

		Dependencies *[]uuid.UUID `json:"dependencies,omitempty" form:"dependencies"`
		Dependents   *[]uuid.UUID `json:"dependents,omitempty" form:"dependents"`
		Similars     *[]uuid.UUID `json:"similars,omitempty" form:"similars"`
	}

	QueryManySubjectsRequest struct {
		filters.FilterPagination
		filters.FilterIDs
		filters.FilterKinds
		filters.FilterLevels
		filters.FilterSlugs
		filters.FilterDecks
		filters.FilterOwners
	}
)

type (
	RemoteContent struct {
		URL         string          `json:"url"`
		ContentType string          `json:"content_type"`
		Metadata    *map[string]any `json:"metadata"`
	}

	StudyData struct {
		Kind     string      `json:"kind"`
		Items    []StudyItem `json:"items"`
		Mnemonic string      `json:"mnemonic,omitempty"`
	}

	StudyItem struct {
		Value         string  `json:"value"`
		IsPrimary     bool    `json:"is_primary"`
		IsValidAnswer bool    `json:"is_valid_answer"`
		IsHidden      bool    `json:"is_hidden"`
		Category      *string `json:"category"`
		Resource      *string `json:"resource,omitempty"`
	}
)

type (
	VocabularySentence struct {
		Sentence string `json:"sentence"`
		Meaning  string `json:"meaning"`
	}

	VocabularyPattern struct {
		Name      string               `json:"name"`
		Sentences []VocabularySentence `json:"sentences"`
	}
)
