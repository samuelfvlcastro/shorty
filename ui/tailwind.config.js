/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "../pkg/views/**/*.{html,ts,qtpl}",
    "../pkg/views/*.{html,ts,qtpl}",
    "../pkg/views/**/**/*.{html,ts,qtpl}"
],
  theme: {
    extend: {},
  },
  plugins: [require("daisyui")],
  daisyui: {
    themes: ["emerald", "dark"],
  },
}
