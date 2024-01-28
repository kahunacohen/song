/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    './templates/*.templ'
  ],
  theme: {
    extend: {
      borderColor: {
        'alert-success': '#c3e6cb',
      },
      fontFamily: {
        sans: ['Nunito', 'sans-serif'],
      },
      textColor: {
        'alert-success': '#155724',
      },
    },

  },
  plugins: [],
}

