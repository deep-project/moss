/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["../../main/tmp/themes/seed/**/*.{html,htm}"],
  darkMode: 'class',
  theme: {
    extend: {},
    container: {
      center: true,
      padding: {
        DEFAULT: '1rem',
        sm: '1rem',
        md: '1rem',
        lg: '1rem',
        xl: '4rem',
        '2xl': '7rem',
      },
    },
  },
  plugins: [
    require('@tailwindcss/line-clamp'),
  ],
}
