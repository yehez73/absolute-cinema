/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./src/**/*.{html,ts}",
  ],
  theme: {
    extend: {
      fontFamily: {
        'popins': ['Poppins', 'sans-serif'],
        'mate-sc': ['Mate SC', 'serif'],
      },
    },
  },
  plugins: [],
}