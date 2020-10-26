const { task, src, dest, series, parallel, watch } = require('gulp')
const del = require('del')

const CSS_SOURCE = 'src/css/*.pcss'
const CSS_DIST = 'public/css'
const JS_SOURCE = 'src/js/**/*.js'
const JS_DIST = 'public/js'
const TMPL_SOURCE = 'templates/**/*.html'

function buildCSS () {
  const sourcemaps = require('gulp-sourcemaps')
  return src(CSS_SOURCE)
    .pipe(sourcemaps.init())
    .pipe(require('gulp-postcss')([
      require('tailwindcss'),
      require('autoprefixer'),
      require('cssnano')
    ]))
    .pipe(require('gulp-rename')({
      extname: '.css'
    }))
    .pipe(sourcemaps.write('.'))
    .pipe(dest(CSS_DIST))
}

function buildJS (done) {
  return src(JS_SOURCE)
    .pipe(require('webpack-stream')(
      require('./webpack.config')
    ))
    .on('error', done)
    .pipe(dest(JS_DIST))
}

task('clean:css', () => {
  return del(CSS_DIST)
})

task('build:css', buildCSS)

task('css', series('clean:css', 'build:css'))

task('clean:js', () => {
  return del(JS_DIST)
})

task('build:js', buildJS)

task('js', series('clean:js', 'build:js'))

task('watch', () => {
  const browserSync = require('browser-sync').create()
  browserSync.init({
    proxy: `http://localhost:${process.env.PORT || 8000}`
  })
  watch(CSS_SOURCE, function streamCSS () {
    return buildCSS()
      .pipe(browserSync.stream())
  })
  watch(JS_SOURCE, function prepareJS (done) {
    buildJS(done)
    browserSync.notify('JavaScript ready!', 2000)
    done()
  })
  watch(TMPL_SOURCE).on('change', browserSync.reload)
})

task('default', parallel('css', 'js'))
