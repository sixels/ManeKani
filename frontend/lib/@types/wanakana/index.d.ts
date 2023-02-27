declare module "wanakana" {
  /**  Automatically bind IME (toKana) functionality to a form textarea or input. */
  declare function bind(elem: HTMLInputElement, ...options: Options[]);
  /**  Unbind IME from element. */
  declare function unbind(elem: HTMLInputElement);
  /**  Returns true if string contains only Hiragana and/or Katakana. */
  declare function isKana(str: string);
  /**  Returns false if string contains mixed characters, otherwise true if Hiragana. */
  declare function isHiragana(str: string);
  /**  Returns false if string contains mixed characters, otherwise true if Katakana. */
  declare function isKatakana(str: string);
  /**  Convert Romaji to Kana. outputs Hiragana by default; uppercase text outputs Katakana. */
  declare function toKana(str: string, ...options: Options[]);
  /**  Convert Katakana or Romaji to Hiragana. */
  declare function toHiragana(str: string, ...options: Options[]);
  /**  Convert Hiragana or Romaji to Katakana. */
  declare function toKatakana(str: string, ...options: Options[]);
  /**  Convert Kana to Romaji. */
  declare function toRomaji(str: string, ...options: Options[]);
  /**  Remove trailing kana. */
  declare function stripOkurigana(str: string, ...options: Options[]);
  /**  Split text into language/kana tokens. */
  declare function tokenize(str: string, ...options: Options[]);

  declare interface Options {
    IMEMode: bool;
    useObsoleteKana: bool;
    passRomaji: bool;
    upcaseKatakana: bool;
    convertLongVowelMark: bool;
    customKanaMapping: { [key: string]: string };
    customRomajiMapping: { [key: string]: string };
  }
}
