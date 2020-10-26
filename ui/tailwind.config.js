module.exports = {
  future: {
    removeDeprecatedGapUtilities: true,
    purgeLayersByDefault: true
  },
  purge: [
    './templates/**/*.html'
  ],
  theme: {
    extend: {
      fontFamily: {
        noto: ['Noto Sans TC', 'sans-serif']
      },
      colors: {
        facebook: '#3b5998',
        instagram: '#3f729b'
      }
    }
  },
  variants: {},
  plugins: []
}
