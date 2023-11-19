import { ReactElement } from "react";
import LevelsMenu from "./components/LevelsMenu";
import ListMenu from "./components/ListMenu";

export type NavItem =
  | { text: string; menu?: ReactElement; mobileOnly?: boolean } & (
      | {
          isLink: false;
        }
      | { isLink: true; href: string }
    );

export type LevelList = { name: string; range: [number, number] }[];

const LEVELS: LevelList = [
  {
    name: "快 pleasant",
    range: [1, 10],
  },
  {
    name: "苦 painful",
    range: [11, 20],
  },
  {
    name: "死 death",
    range: [21, 30],
  },
  {
    name: "地獄 hell",
    range: [31, 40],
  },
  {
    name: "天国 paradise",
    range: [41, 50],
  },
  {
    name: "現実 reality",
    range: [51, 60],
  },
];

export const menuItems: NavItem[] = [
  {
    isLink: false,
    text: "level",
    menu: <LevelsMenu levels={LEVELS} />,
  },
  {
    isLink: false,
    text: "radical",
    menu: <ListMenu levels={LEVELS} kind={"radical"} />,
  },
  {
    isLink: false,
    text: "kanji",
    menu: <ListMenu levels={LEVELS} kind={"kanji"} />,
  },
  {
    isLink: false,
    text: "vocabulary",
    menu: <ListMenu levels={LEVELS} kind={"vocabulary"} />,
  },
];
