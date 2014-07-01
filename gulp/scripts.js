var gulp       = require('gulp')
var concat     = require('gulp-concat')
var ngAnnotate = require('gulp-ng-annotate')

gulp.task('scripts', function () {
  gulp.src(['js/**/module.js', 'js/**/*.js'])
    .pipe(ngAnnotate())
    .pipe(concat('app.js'))
    .pipe(gulp.dest('assets'))
})

gulp.task('watch:scripts', ['scripts'], function () {
  gulp.watch('js/**/*.js', ['scripts'])
})
