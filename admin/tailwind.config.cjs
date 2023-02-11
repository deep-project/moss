/** @type {import('tailwindcss').Config} */
module.exports = {
  // important: true,
  content: [
    './index.html',
    './src/*.{vue,js,ts}',
    './src/**/*.{vue,js,ts}',
  ],
  theme: {
    extend: {},
  },
  plugins: [],
  corePlugins: {
    preflight: false
  },
}
