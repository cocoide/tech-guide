/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./pages/**/*.{js,ts,jsx,tsx,mdx}",
    "./components/**/*.{js,ts,jsx,tsx,mdx}",
    "./app/**/*.{js,ts,jsx,tsx,mdx}",
  ],
  theme: {
    extend: {
      keyframes: {
        scale: {
          "0%": { transform: "scale(.8)" },
          "100%": { transform: "scale(1)" },
        },
        appear: {
          "0%": { opacity: 0 },
          "100%": { opacity: 1 },
        },
      },
      animation: {
        scale: "scale 0.2s ease-in-out",
        appear: "appear 0.2s ease-in-out",
      },
    },
  },
  plugins: [require("@tailwindcss/forms")],
};
