/** @type {import('tailwindcss').Config} */
export default {
  purge: ['./index.html', './src/**/*.{vue,js,ts,jsx,tsx}'],
  content: [],
  theme: {
    backgroundColor: theme => ({
      'primary': '#57B8FF',
      'secondary': '#ffed4a',
      'menuCard': '#444',
     })
  },
  plugins: [],
}

