const { security, external: { facebook } } = require("../config");
const Router = require("@koa/router"),
    jwt = require("jsonwebtoken");
const User = require("../app/user");
const router = new Router();

router.get("/login", async (ctx) => {
    await ctx.render("user_login", {
        title: "登入",
        appId: facebook.appId
    });
});

router.get("/register", async (ctx) => {
    await ctx.render("user_register", {
        title: "建立帳號",
        appId: facebook.appId
    });
});

router.post("/login", async (ctx) => {
    try {
        let { email, password } = ctx.request.body;
        let user = await User.getUserByPassword(email, password);
        let token = jwt.sign({ id: user._id }, security.jwt);
        ctx.cookies.set("user", token, { httpOnly: true, signed: true });
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