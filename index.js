const { server, security } = require("./config");
const Koa = require("koa"),
    http = require("http"),
    bodyParser = require("koa-bodyparser"),
    assets = require("koa-static"),
    views = require("koa-views"),
    path = require("path"),
    helmet = require("koa-helmet")
    database = require("./app/database"),
    json = require("./app/json");
const app = new Koa(),
    router = require("./routes/index");

async function initApp() {
    app.use(helmet());
    app.use(bodyParser());
    app.use(assets("public"));
    app.use(views(path.join(__dirname, "views"), { extension: "pug" }));
    app.use(json());
    app.use(router.routes());
    app.use(router.allowedMethods());
    app.proxy = server.proxy;
    app.keys = security.keys;
    await database.connect();
}

function startApp() {
    let httpServer = http.createServer(app.callback());
    httpServer.listen(server.port);
    console.log(`Server listening on http://localhost:${server.port}`);
    // console.debug(`Server listening on routes:`, router.stack.map(i=>i.path));
}

(async () => {
    await initApp();
    startApp();
})();