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

router.get("/enable-dev", async (ctx) => {
    ctx.cookies.set("developer", true);
    ctx.redirect("/");
});

router.get("/coming-soon", async (ctx) => {
    await ctx.render("coming-soon", {
        user: ctx.user,
        title: "即將上線"
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
router.use("/admin", require("./admin").routes());

module.exports = router;