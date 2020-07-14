const { development } = require("./config");
const { task, src, dest, series, parallel, watch } = require("gulp"),
    vinylPaths = require("vinyl-paths"),
    sourcemaps = require("gulp-sourcemaps"),
    stylus = require("gulp-stylus"),
    poststylus = require("poststylus"),
    rucksack = require("rucksack-css"),
    cleanCSS = require("gulp-clean-css"),
    del = require("del"),
    named = require("vinyl-named"),
    webpack = require("webpack-stream"),
    TerserPlugin = require("terser-webpack-plugin"),
    sync = require("browser-sync");

task("css:delete", () => {
    return src("public/css/*")
        .pipe(vinylPaths(del))
});

task("css:build", () => {
    return src("src/css/**/*.styl")
        .pipe(sourcemaps.init())
        .pipe(stylus({
            use: [ poststylus([ rucksack({ autoprefixer: true }) ]) ]
        }))
        .pipe(cleanCSS())
        .pipe(sourcemaps.write("."))
        .pipe(dest("public/css"));
});

task("css", series("css:delete", "css:build"));

task("js:delete", () => {
    return src("public/js/*")
        .pipe(vinylPaths(del));
});

task("js:build", (done) => {
    return src("src/js/**/*.js")
        .pipe(named())
        .pipe(webpack({
            ...(development ? {
                mode: "development",
                devtool: "eval-source-map"
            } : {
                mode: "production"
            }),
            module: {
                rules: [
                    {
                        test: /\.js$/,
                        exclude: /node_modules/,
                        use: {
                            loader: "babel-loader",
                            options: {
                                presets: ["@babel/preset-env"],
                                plugins: [
                                    "@babel/plugin-transform-runtime",
                                    "@babel/plugin-proposal-optional-chaining"
                                ]
                            }
                        }
                    }
                ]
            },
            optimization: {
                minimize: true,
                minimizer: [new TerserPlugin()]
            }
        }))
        .on("error", done)
        .pipe(dest("public/js"));
});

task("js", series("js:delete", "js:build"));

task("images:delete", () => {
    return src("public/images/*")
        .pipe(vinylPaths(del));
});

task("images:copy", () => {
    return src("src/images/*")
        .pipe(dest('public/images'));
});

task("images", series("images:delete", "images:copy"));

task("build", parallel("css", "js", "images"));

task("reload", (done) => {{
    sync.reload();
    done();
}})

task("watch", (done) => {
    let { browserSync } = require("./config").external;
    sync.init(browserSync);
    watch("src/js/**/*.js", series("js", "reload"));
    watch("src/css/**/*.styl", series("css", "reload"));
    watch("views/*.pug", series("reload"));
    done();
});