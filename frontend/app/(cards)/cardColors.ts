export type CardKind = 'radical' | 'kanji' | 'vocabulary';

export type CardBgAccent = `bg-card-${CardKind}`;
export type CardBgAccentDark = `${CardBgAccent}-dark`;
export type CardBgAccentLight = `${CardBgAccent}-light`;

export const CARD_BG: {
  [name in CardKind]: CardBgAccent;
} = Object.freeze({
  radical: 'bg-card-radical',
  kanji: 'bg-card-kanji',
  vocabulary: 'bg-card-vocabulary',
});

export const CARD_BG_DARK: {
  [name in CardKind]: CardBgAccentDark;
} = Object.freeze({
  radical: 'bg-card-radical-dark',
  kanji: 'bg-card-kanji-dark',
  vocabulary: 'bg-card-vocabulary-dark',
});

export const CARD_BG_LIGHT: {
  [name in CardKind]: CardBgAccentLight;
} = Object.freeze({
  radical: 'bg-card-radical-light',
  kanji: 'bg-card-kanji-light',
  vocabulary: 'bg-card-vocabulary-light',
});
