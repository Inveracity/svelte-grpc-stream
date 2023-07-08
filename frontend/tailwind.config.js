module.exports = {
  content: ['./src/routes/**/*.{svelte,js,ts}', './src/components/**/*.{svelte,js,ts}'],
  plugins: [require('daisyui')],
  daisyui: {
    themes: [
      {
        mytheme: {

          "primary": "#21efde",

          "secondary": "#ed84cf",

          "accent": "#81f4b5",

          "neutral": "#181b25",

          "base-100": "#44304b",

          "info": "#9cc3ec",

          "success": "#1c9274",

          "warning": "#ed9c2c",

          "error": "#e87380",
                   },
      },
    ]},
};
