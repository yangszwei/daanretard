const { app } = require("../config");
const Router = require("@koa/router");
const App = require("../app/app");
const router = new Router();

router.get("/", async (ctx) => {
    await ctx.render("index", {
        user: ctx.user,
        version: app.version
    });
});

router.get("/rules", async (ctx) => {
    // TODO: load & cache rules on app start to ease database pressure
    await ctx.render("rules", {
        rules: (await App.rules).content
    });
});

router.use("/post", require("./post").routes());
router.use("/user", require("./user").routes());

module.exports = router;