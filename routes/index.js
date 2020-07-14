const { app } = require("../config");
const Router = require("@koa/router");
const config = require("../app/config");
const router = new Router();

router.get("/", async (ctx) => {
    await ctx.render("index", {
        user: ctx.user,
        version: app.version
    });
});

router.get("/rules", async (ctx) => {
    await ctx.render("rules", {
        title: "版規",
        rules: await config.getPostingRules(),
        user: ctx.user
    });
});

router.use("/user", require("./user").routes());
router.use("/post", require("./post").routes());
router.use("/resource", require("./resource").routes());

module.exports = router;