const { app } = require("../config");
const Router = require("@koa/router");
const router = new Router();

router.get("/", async (ctx) => {
    await ctx.render("index", {
        year: new Date().getFullYear().toString(),
        version: app.version
    });
});

router.use("/user", require("./user").routes());

module.exports = router;