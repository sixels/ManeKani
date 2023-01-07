package cards

type Level struct {
	Kanji      []*PartialKanjiResponse      `json:"kanji,omitempty"`
	Radical    []*PartialRadicalResponse    `json:"radical,omitempty"`
	Vocabulary []*PartialVocabularyResponse `json:"vocabulary,omitempty"`
}
