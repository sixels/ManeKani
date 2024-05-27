function withOpacity(variable) {
  return ({ opacityValue }) => {
    if (opacityValue === undefined) {
      return `rgb(var(${variable}))`;
    }
    return `rgb(var(${variable}) / ${opacityValue})`;
  };
}

/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./base.html", "./settings/**/*.html", "./components/**/*.html"],
  theme: {
    extend: {
      colors: {
        text: {
          50: withOpacity("--text-50"),
          100: withOpacity("--text-100"),
          200: withOpacity("--text-200"),
          300: withOpacity("--text-300"),
          400: withOpacity("--text-400"),
          500: withOpacity("--text-500"),
          600: withOpacity("--text-600"),
          700: withOpacity("--text-700"),
          800: withOpacity("--text-800"),
          900: withOpacity("--text-900"),
          950: withOpacity("--text-950"),
          DEFAULT: withOpacity("--text"),
        },
        background: {
          50: withOpacity("--background-50"),
          100: withOpacity("--background-100"),
          200: withOpacity("--background-200"),
          300: withOpacity("--background-300"),
          400: withOpacity("--background-400"),
          500: withOpacity("--background-500"),
          600: withOpacity("--background-600"),
          700: withOpacity("--background-700"),
          800: withOpacity("--background-800"),
          900: withOpacity("--background-900"),
          950: withOpacity("--background-950"),
          DEFAULT: withOpacity("--background"),
        },
        primary: {
          50: withOpacity("--primary-50"),
          100: withOpacity("--primary-100"),
          200: withOpacity("--primary-200"),
          300: withOpacity("--primary-300"),
          400: withOpacity("--primary-400"),
          500: withOpacity("--primary-500"),
          600: withOpacity("--primary-600"),
          700: withOpacity("--primary-700"),
          800: withOpacity("--primary-800"),
          900: withOpacity("--primary-900"),
          950: withOpacity("--primary-950"),
          DEFAULT: withOpacity("--primary"),
        },
        secondary: {
          50: withOpacity("--secondary-50"),
          100: withOpacity("--secondary-100"),
          200: withOpacity("--secondary-200"),
          300: withOpacity("--secondary-300"),
          400: withOpacity("--secondary-400"),
          500: withOpacity("--secondary-500"),
          600: withOpacity("--secondary-600"),
          700: withOpacity("--secondary-700"),
          800: withOpacity("--secondary-800"),
          900: withOpacity("--secondary-900"),
          950: withOpacity("--secondary-950"),
          DEFAULT: withOpacity("--secondary"),
        },
        accent: {
          50: withOpacity("--accent-50"),
          100: withOpacity("--accent-100"),
          200: withOpacity("--accent-200"),
          300: withOpacity("--accent-300"),
          400: withOpacity("--accent-400"),
          500: withOpacity("--accent-500"),
          600: withOpacity("--accent-600"),
          700: withOpacity("--accent-700"),
          800: withOpacity("--accent-800"),
          900: withOpacity("--accent-900"),
          950: withOpacity("--accent-950"),
          DEFAULT: withOpacity("--accent"),
        },
      },
      borderOpacity: ({ theme }) => theme("opacity"),
    },
  },
  plugins: [require("daisyui")],
  daisyui: {
    themes: [
      {
        light: {
          ...require("daisyui/src/theming/themes")["light"],
          text: "#100f10",
          background: "#f3f3f3",
          primary: "#441e48",
          secondary: "#ddd4df",
          accent: "#cd65e2",
        },
        dark: {
          ...require("daisyui/src/theming/themes")["dark"],
          text: "#f0eff0",
          background: "#100f10",
          primary: "#dfb9df",
          secondary: "#4f3b54",
          accent: "#851d9a",
        },
      },
    ],
  },
};
