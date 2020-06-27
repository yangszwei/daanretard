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
    sync = require("browser-sync");

task("css:delete", () => {
    return src("public/stylesheets/*")
        .pipe(print())
        .pipe(vinylPaths(del))
});

task("css:build", () => {
    return src("src/stylesheets/*.pcss")
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

task("css", series("css:delete", "css:build"));

task("js:delete", () => {
    return src("public/javascripts/*")
        .pipe(print())
        .pipe(vinylPaths(del));
});

task("js:build", () => {
    return src("src/javascripts/*.js")
        .pipe(print())
        .pipe(sourcemaps.init())
        .pipe(terser())
        .pipe(sourcemaps.write("."))
        .pipe(dest("public/javascripts"));
});

task("js", series("js:delete", "js:build"));

task("build", parallel("css", "js"));

task("reload", (done) => {{
    sync.reload();
    done();
}})

task("watch", () => {
    let { port } = require("./config").server;
    sync.init({ proxy: `http://localhost:${port}` });
    watch("src/javascripts/*.js", series("js", "reload"));
    watch("src/stylesheets/*.pcss", series("css", "reload"));
    watch("views/*.pug", series("reload"));
});