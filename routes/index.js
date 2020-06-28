const { app } = require("../config");
const Router = require("@koa/router");
const router = new Router();

router.get("/", async (ctx) => {
    await ctx.render("index", {
        year: "2020",
        version: app.version
    });
});

module.exports = router;