type RetroFrameProps = {
  backColor: string;
  children?: React.ReactNode;
};

export default function RetroFrame({ backColor, children }: RetroFrameProps) {
  return (
    <div className="frame relative z-10 h-full w-full">
      <div
        className={`frame-front h-full w-full after:absolute after:left-[6px] after:top-[6px] after:-z-10 after:h-full after:w-full after:rounded-md after:ring-[3px] after:ring-black after:content-[''] ${backColor}`}
      >
        <div className="content-wrapper h-full w-full  overflow-hidden rounded-md ring-[3px] ring-black">
          {children}
        </div>
      </div>
    </div>
  );
}
