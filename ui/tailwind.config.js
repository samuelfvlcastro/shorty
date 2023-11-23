/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "../pkg/views/**/*.{html,ts}",
    "../pkg/views/*.{html,ts}",
    "../pkg/views/**/**/*.{html,ts}"
],
  theme: {
    extend: {},
  },
  plugins: [require("daisyui")],
  daisyui: {
    themes: ["emerald", "dark"],
  },
}
