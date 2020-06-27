const { server } = require("./config");
const Koa = require("koa"),
    http = require("http"),
    bodyParser = require("koa-bodyparser"),
    assets = require("koa-static"),
    views = require("koa-views"),
    path = require("path"),
    helmet = require("koa-helmet");
const app = new Koa(),
    router = require("./routes/index");

function initApp() {
    app.use(helmet());
    app.use(bodyParser());
    app.use(assets("public"));
    app.use(views(path.join(__dirname, "views"), { extension: "pug" }));
    app.use(router.routes());
    app.use(router.allowedMethods());
    app.proxy = server.proxy;
}

function startApp() {
    let httpServer = http.createServer(app.callback());
    httpServer.listen(server.port);
    console.log(`Server listening on http://localhost:${server.port}`);
}

(async () => {
    initApp();
    startApp();
})();