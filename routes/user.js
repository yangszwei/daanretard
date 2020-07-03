const { security, external: { facebook } } = require("../config");
const Router = require("@koa/router"),
    jwt = require("jsonwebtoken");
const User = require("../app/user");
const router = new Router();
const SECONDS_IN_A_MONTH = 60 * 60 * 24 * 30;

router.get("/", async (ctx) => {
    ctx.redirect(ctx.user ? "/user/me" : "/user/register");
});

router.get("/login", async (ctx) => {
    if (ctx.user) {
        ctx.redirect("/user/me");
    } else {
        await ctx.render("user-login", {
            title: "登入",
            appId: facebook.appId
        });
    }
});

router.get("/register", async (ctx) => {
    if (ctx.user) {
       ctx.redirect("/user/me");
    } else {
        await ctx.render("user-register", {
            title: "建立帳戶",
            appId: facebook.appId
        });
    }
});

router.get("/me", async (ctx) => {
    if (ctx.user) {
        // TODO: handle error when fetching user profile fail
        await ctx.render("user-me", {
            title: "我的帳戶",
            user: await User.getUserByObjectId(ctx.user.id)
        });
    } else {
        await ctx.redirect("/user/register");
    }

})

router.post("/login", async (ctx) => {
    try {
        let { email, password } = ctx.request.body;
        let user = await User.getUserByPassword(email, password);
        let token = jwt.sign({ id: user._id }, security.jwt);
        ctx.cookies.set("user", token, {
            httpOnly: true,
            signed: true,
            maxAge: SECONDS_IN_A_MONTH * 2
        });
        await ctx.json({ status: "success" });
    } catch(err) {
        if (err === "Invalid Credential") {
            await ctx.json({ status: "failed", reason: err });
        } else if (err === "No Valid Provider") {
            await ctx.json({ status: "failed", reason: err });
        } else {
            console.error(err);
            await ctx.throw(500, "Internal Server Error");
        }
    }
});

router.post("/register", async (ctx) => {
    try {
        let user = await User.createUserWithPassword(ctx.request.body);
        let token = jwt.sign({ id: user._id }, security.jwt);
        ctx.cookies.set("user", token, { httpOnly: true, signed: true });
        await ctx.json({ status: "success" });
    } catch(err) {
        if (err === "User Already Exists") {
            await ctx.json({ status: "failed", reason: err });
        } else if (err.isJoi) {
            await ctx.json({
                status: "failed",
                reason: "Validation Failed",
                details: err.details
            });
        } else {
            console.error(err);
            await ctx.throw(500, "Internal Server Error");
        }
    }
});

router.post("/oauth/facebook", async (ctx) => {
    try {
        let { accessToken } = ctx.request.body;
        let user = await User.continueWithFacebook(accessToken);
        let token = jwt.sign({ id: user._id }, security.jwt);
        ctx.cookies.set("user", token, { httpOnly: true, signed: true });
        await ctx.json({ status: "success" });
    } catch(err) {
        if (err === "User Already Exists") {
            await ctx.json({ status: "failed", reason: err });
        } else {
            console.error(err);
            await ctx.throw(500, "Internal Server Error");
        }
    }
});

module.exports = router;