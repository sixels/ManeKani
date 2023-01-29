import { extendTheme } from "@chakra-ui/react";

export const cardColors: { [key: string]: string } = {
  radical: "blue.500",
  kanji: "pink.500",
  vocabulary: "purple.500",
};

export const theme = extendTheme({
  fonts: {
    heading: `'Inter', 'Roboto', 'Helvetica', sans-serif`,
    body: `'Inter', 'Roboto', 'Helvetica', sans-serif`,
  },
  theme: "light",
  colors: {
    orange: {
      "50": "#FFECE5",
      "100": "#FFC9B8",
      "200": "#FFA68A",
      "300": "#FF835C",
      "400": "#FF612E",
      "500": "#FF3E00",
      "600": "#CC3100",
      "700": "#992500",
      "800": "#661900",
      "900": "#330C00",
    },
  },
});
