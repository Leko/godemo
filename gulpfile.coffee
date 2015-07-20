'use strict';

gulp = require 'gulp'
sass = require 'gulp-sass'

PATHS =
  SCSS: './assets/scss/**/*.scss'
  SCSS_MAIN: './assets/scss/all.scss'
  SCSS_DIST: './assets/dist/css'

gulp.task 'build:sass', ->
  gulp.src(PATHS.SCSS_MAIN)
    .pipe(sass(outputStyle: 'compressed').on('error', sass.logError))
    .pipe(gulp.dest(PATHS.SCSS_DIST))

gulp.task 'default', ['build:sass']
