/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    './app/**/*.{js,ts,jsx,tsx}',
    './page/**/*.{js,ts,jsx,tsx}',
    './ui/**/*.{js,ts,jsx,tsx}',
  ],
  theme: {
    extend: {
      colors: {
        card: {
          radical: {
            DEFAULT: '#1eae53',
            dark: '#167E3C',
            light: '#24D164',
          },
          kanji: {
            DEFAULT: '#FF2400',
            dark: '#C71C00',
            light: '#FF4729',
          },
          vocabulary: {
            DEFAULT: '#5A4FCF',
            dark: '#3D31B5',
            light: '#786FD8',
          },
        },
      },
      fontFamily: {
        grotesk: ['"Rubik"', 'serif'],
        sans: ['"Inter"', '"Noto Sans"', 'sans-serif'],
      },
    },
  },
  plugins: [],
};
