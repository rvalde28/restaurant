/** @type {import('tailwindcss').Config} */
export default {
  purge: ['./index.html', './src/**/*.{vue,js,ts,jsx,tsx}'],
  content: [],
  theme: {
    extend: {
      colors: {
        'primary': '#57B8FF',
        'secondary': '#b992c1',
        'third': '#B6DCFE',
        'fourth':'#dec5e3',
        'menuCard': '#444',
        'menubtn': '#57B8FF'

      },
    },
  },
  daisyui: {
    themes: ["light"],
  },
  plugins: [
    require('daisyui'),
  ],
}

