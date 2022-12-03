import RetroFrame from './RetroFrame';

export default function Header() {
  return (
    <header className="flex h-20 w-full items-center justify-between bg-gray-100 px-5">
      <div className="left">
        <div className="logo text-3xl" lang="ja">
          <ruby className="raw-kanji">
            真<rt className="furigana text-sm">ま</rt>
          </ruby>
          <ruby className="raw-kanji">
            似<rt className="furigana text-sm">ね</rt>
          </ruby>
          <ruby className="raw-kanji">
            蟹<rt className="furigana text-sm">かに</rt>
          </ruby>
        </div>
      </div>
      <div className="right flex items-center gap-5 divide-x-2 divide-gray-400 font-grotesk after:bg-red-900">
        <div className="shortcuts flex gap-5">
          <span>Radical</span>
          <span>Kanji</span>
          <span>Vocabulary</span>
          <div className="search-icon">S</div>
        </div>
        <div className="login-options flex items-center gap-5 pl-5">
          <span>Login</span>
          <RetroFrame backColor="after:bg-gray-600">
            <span className="bg-gray-100 px-3 py-1">Sign Up</span>
          </RetroFrame>
        </div>
      </div>
    </header>
  );
}
