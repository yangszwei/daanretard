const { app } = require("../config");
const Router = require("@koa/router");
const Post = require("../app/post"),
    { RESULT, STATUS } = require("../app/utils/codes");
const router = new Router();

router.get("/create", async (ctx) => {
    await ctx.render("post-create", {
        title: "建立貼文",
        user: ctx.user
    });
});

router.get("/s/:id", async (ctx) => {
    let submission = await Post.getSubmissionById(ctx.params.id);
    let status = "無貼文資訊";
    if (submission && submission.stage) {
        if (submission.stage === "submission") {
            status = submission.verified ? "未審核" : "未驗證電子郵件";
        }
        if (submission.stage === "pending post") status = "已通過審核";
        if (submission.stage === "published") status = "已發佈";
    }
    await ctx.render("post-submission-status", {
        title: "貼文狀態",
        submission_id: ctx.params.id,
        submission_status: status
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

router.post("/send-verification", async (ctx) => {
    let { email } = ctx.request.fields;
    try {
        await Post.sendVerificationEmail(email);
        await ctx.json({ code: RESULT.SUCCESS });
    } catch (err) {
        await ctx.json(handleError(err));
    }
});

router.post("/verify-email", async (ctx) => {
    let { token, email } = ctx.request.fields;
    try {
        await Post.verifyEmail(token, email);
        await ctx.json({ code: RESULT.SUCCESS });
    } catch (err) {
        await ctx.json(handleError(err));
    }
});

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
            redirect_url: `${app.url}/post/s/${result}`
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

});

module.exports = router;