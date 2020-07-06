const { external: { facebook } } = require("../config");
const Router = require("@koa/router");
const User = require("../app/user"),
    RESULT = require("../app/result-code");
const router = new Router();
const MILLISECONDS_IN_A_MONTH = 60 * 60 * 24 * 30 * 1000;

function getCookieOption() {
    return {
        httpOnly: true,
        signed: true,
        maxAge: MILLISECONDS_IN_A_MONTH * 2000
    };
}

router.get("/", async (ctx) => {
    ctx.redirect(ctx.user ? "/user/me" : "/user/register");
});

router.get("/login", async (ctx) => {
    if (ctx.user) {
        ctx.redirect("/user");
        return;
    }
    await ctx.render("user-login", {
        title: "登入",
        appId: facebook.appId
    });
});

router.get("/register", async (ctx) => {
    if (ctx.user) {
       ctx.redirect("/user/me");
       return;
    }
    await ctx.render("user-register", {
        title: "建立帳戶",
        appId: facebook.appId
    });
});

router.get("/me", async (ctx) => {
    if (!ctx.user) {
        ctx.redirect("/user");
        return;
    }
    try {
        await ctx.render("user-me", {
            title: "我的帳戶",
            user: await User.getUserByObjectID(ctx.user.id)
        });
    } catch(err) {
        ctx.throw(500);
    }
})

function handleError(err) {
    if (typeof err === "number") {
        return { code: err };
    } else {
        console.error(err);
        return { code: RESULT.INTERNAL_ERROR };
    }
}

router.post("/login", async (ctx) => {
    try {
        let { email, password } = ctx.request.body;
        let user = await User.loginWithPassword(email, password);
        let token = User.signUserToken(user);
        ctx.cookies.set("user", token, getCookieOption());
        await ctx.json({ code: RESULT.SUCCESS });
    } catch(err) {
        if (err === RESULT.NOT_FOUND) {
            await ctx.json({ code: RESULT.INVALID_INPUT });
        } else {
            await ctx.json(handleError(err));
        }
    }
});

router.get("/logout", async (ctx) => {
    ctx.cookies.set("user", null);
    ctx.cookies.set("updated", null);
    ctx.redirect("/");
});

router.post("/register", async (ctx) => {
    try {
        let user = await User.registerWithPassword(ctx.request.body);
        let token = User.signUserToken(user);
        ctx.cookies.set("user", token, getCookieOption());
        await ctx.json({ code: RESULT.SUCCESS });
    } catch(err) {
        if (err.isJoi) {
            await ctx.json({
                code: RESULT.INVALID_INPUT,
                details: err.details
            });
        } else {
            await ctx.json(handleError(err));
        }
    }
});

router.post("/oauth/facebook", async (ctx) => {
    try {
        let { accessToken } = ctx.request.body;
        let user = await User.continueWithFacebook(accessToken);
        let token = User.signUserToken(user);
        ctx.cookies.set("user", token, getCookieOption());
        await ctx.json({ code: RESULT.SUCCESS });
    } catch(err) {
        await ctx.json(handleError(err));
    }
});

module.exports = router;