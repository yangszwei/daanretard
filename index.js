const { server, security } = require("./config");
const Koa = require("koa"),
    http = require("http"),
    bodyParser = require("koa-better-body"),
    assets = require("koa-static"),
    views = require("koa-views"),
    path = require("path"),
    helmet = require("koa-helmet"),
    database = require("./app/utils/database"),
    loadUser = require("./middlewares/load-user"),
    errorPage = require("./middlewares/error-page");
const app = new Koa();

!function json() {
    app.use(async (ctx, next) => {
        ctx.json = (object) => {
            ctx.body = JSON.stringify(object);
        };
        await next();
    });
}();

!function () {
    if (!String.prototype.replaceAll) {
        String.prototype.replaceAll = function(search, replacement) {
            let target = this;
            return target.replace(new RegExp(search, 'g'), replacement);
        };
    }
}();

async function initApp() {
    app.use(helmet());
    app.use(bodyParser({
        formLimit : "15mb",
        jsonLimit:"15mb",
        textLimit:"15mb"
    }));
    app.use(assets("public"));
    app.use(views(path.join(__dirname, "views"), { extension: "pug" }));
    app.use(loadUser());
    app.use(errorPage());
    await database.connect();
    app.use(require("./routes/index").routes());
    app.use(require("./routes/index").allowedMethods());
    app.proxy = server.proxy;
    app.keys = security.keys;
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