const { task, src, dest, series, parallel, watch } = require("gulp"),
    print = require("gulp-print").default,
    vinylPaths = require("vinyl-paths"),
    rename = require("gulp-rename"),
    sourcemaps = require("gulp-sourcemaps"),
    postcss = require("gulp-postcss"),
    postcssImport = require("postcss-import"),
    extend = require("postcss-extend"),
    precss = require("precss"),
    cssnano = require("cssnano"),
    autoprefixer = require("autoprefixer"),
    del = require("del"),
    terser = require("gulp-terser"),
    sync = require("browser-sync"),
    imagemin = require('gulp-imagemin');

task("stylesheets:delete", () => {
    return src("public/stylesheets/*")
        .pipe(print())
        .pipe(vinylPaths(del))
});

task("stylesheets:build", () => {
    return src("src/stylesheets/**/*.pcss")
        .pipe(print())
        .pipe(sourcemaps.init())
        .pipe(postcss([
            postcssImport, extend,
            precss, autoprefixer, cssnano
        ]))
        .pipe(rename({ extname: ".css" }))
        .pipe(sourcemaps.write("."))
        .pipe(dest("public/stylesheets"));
});

task("stylesheets", series("stylesheets:delete", "stylesheets:build"));

task("javascripts:delete", () => {
    return src("public/javascripts/*")
        .pipe(print())
        .pipe(vinylPaths(del));
});

task("javascripts:build", () => {
    return src("src/javascripts/**/*.js")
        .pipe(print())
        .pipe(sourcemaps.init())
        .pipe(terser())
        .pipe(sourcemaps.write("."))
        .pipe(dest("public/javascripts"));
});

task("javascripts", series("javascripts:delete", "javascripts:build"));

task("images:delete", () => {
    return src("public/images/*")
        .pipe(print())
        .pipe(vinylPaths(del));
});

task("images:build", () => {
    return src("src/images/*")
        .pipe(print())
        .pipe(imagemin())
        .pipe(dest('public/images'));
});

task("images", series("images:delete", "images:build"));

task("build", parallel("stylesheets", "javascripts", "images"));

task("reload", (done) => {{
    sync.reload();
    done();
}})

task("watch", () => {
    let { browserSync } = require("./config").external;
    sync.init(browserSync);
    watch("src/javascripts/**/*.js", series("javascripts", "reload"));
    watch("src/stylesheets/**/*.pcss", series("stylesheets", "reload"));
    watch("views/*.pug", series("reload"));
});