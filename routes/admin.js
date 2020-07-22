const Router = require("@koa/router");
const router = new Router();

router.get("/", async (ctx) => {
    ctx.redirect("/admin/review");
});

router.get("/review", async (ctx) => {
    await ctx.render("review.admin", {
        posts: await Post.listNotReviewed()
    });
});

module.exports = router;