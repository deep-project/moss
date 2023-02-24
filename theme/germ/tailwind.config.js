/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["../../main/tmp/themes/germ/**/*.{html,htm}"],
  darkMode: 'class',
  theme: {
    extend: {
      gridTemplateColumns: {
        // Simple 16 column grid
        '16': 'repeat(16, minmax(0, 1fr))',
      },
      gridColumn: {
        'span-16': 'span 16 / span 16',
      }
    },
    container: {
      screens: {
        sm: '640px',
        md: '768px',
        lg: '980px',
        xl: '1000px',
        '2xl': '1050px',
      },
      padding: {
        DEFAULT: '.7rem',
        sm: '1rem',
        md: '1rem',
        lg: '1rem',
        xl: '1rem',
        '2xl': '1rem',
      },
      center: true,
    },
  },
  plugins: [
    require('@tailwindcss/line-clamp'),
  ],
}
