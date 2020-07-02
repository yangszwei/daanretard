const Router = require("@koa/router");
const router = new Router();
const PostSubmission = require("../app/post-submission");

router.get("/create", async (ctx) => {
    await ctx.render("post-create", {
        title: "建立貼文",
        user: ctx.user
    });
});

router.post("/create", async (ctx) => {
    let { content, media, email } = ctx.request.body;
    let post = {
        content: content || "",
        media: media || [],
        ...(email ? { email: email } : {}),
        verified: Boolean(ctx.user)
    };
    post.author = ctx.user || "email";
    if (!ctx.user && !email) {
        await ctx.json({
            status: "failed",
            reason: "Invalid Identity"
        });
    } else {
        let result = await PostSubmission.create(ctx.request.body);
        if (!ctx.user) {
            await ctx.json({
                status: "success",
                id: result.insertedId
            });
        } else {
            await ctx.json({ status: "success" });
        }
    }
});

module.exports = router;