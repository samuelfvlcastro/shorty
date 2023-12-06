/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "../pkg/templates/**/*.go.html",
    "../pkg/templates/*.go.html",
],
  theme: {
    extend: {},
  },
  plugins: [require("daisyui")],
  daisyui: {
    themes: ["emerald", "dark"],
  },
}
