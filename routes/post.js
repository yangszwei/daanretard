const { app } = require("../config");
const Router = require("@koa/router");
const Post = require("../app/post"),
    { RESULT } = require("../app/utils/codes");
const router = new Router();

router.get("/create", async (ctx) => {
    await ctx.render("post-create", {
        title: "建立貼文",
        user: ctx.user
    });
});

function handleError(err) {
    if (typeof err === "number") {
        return { code: err };
    } else {
        console.error(err);
        return { code: RESULT.FATAL_ERROR };
    }
}

router.post("/submit", async (ctx) => {
    let submission = ctx.request.fields;
    try {
        let result = await Post.submit({
            content: submission.content,
            images: submission.images,
            email: submission.email || null
        }, ctx.user);
        await ctx.json({
            code: RESULT.SUCCESS,
            redirect_url: `${app.url}/post/submission_id/${result}`
        });
    } catch(err) {
        if (typeof err === "string") {
            await ctx.json({
                code: RESULT.INVALID_QUERY,
                details: err
            });
        } else {
            await ctx.json(handleError(err));
        }

    }
});

router.post("/review", async (ctx) => {

})

module.exports = router;