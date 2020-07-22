const Router = require("@koa/router"),
    Post = require("../app/post");
const router = new Router();

router.get("/", async (ctx) => {
    ctx.redirect("/admin/review");
});

router.get("/review", async (ctx) => {
    await ctx.render("admin-review", {
        posts: await Post.listNotReviewed(0)
    });
});

module.exports = router;