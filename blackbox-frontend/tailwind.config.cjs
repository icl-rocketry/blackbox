module.exports = {
  plugins: [],
  theme: {
    extend: {
      transitionProperty: {
        'width': 'width'
      },
    },
  },
  purge: ["./index.html", './src/**/*.{svelte,js,ts}'], // for unused CSS
  variants: {
    extend: {},
  },
  darkMode: 'class', // or 'media' or 'class'
}