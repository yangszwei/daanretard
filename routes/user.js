const Router = require("@koa/router");
const User = require("../app/user"),
    { RESULT } = require("../app/utils/codes");
const router = new Router();
const COOKIE_OPTIONS = {
    httpOnly: true,
    signed: true,
    maxAge: 86400 * 30 * 2 * 1000 // 60 days in milliseconds
};

router.get("/sign-in", async (ctx) => {
    if (ctx.user) ctx.redirect("/user/profile");
    else await ctx.render("user-sign-in", {
        title: "登入"
    });
});

router.get("/register", async (ctx) => {
    if (ctx.user) ctx.redirect("/user/profile");
    else await ctx.render("user-register", {
        title: "建立帳戶"
    });
});

router.get("/forgot", async (ctx) => {
    if (ctx.user) ctx.redirect("/user/profile");
    else await ctx.render("user-forgot", {
        title: "忘記密碼"
    });
});

router.get("/reset-password", async (ctx) => {
    await ctx.render("user-reset-password", {
        title: "重設密碼"
    });
});

router.get("/profile", async (ctx) => {
    if (!ctx.user) ctx.redirect("/user/sign-in");
    else await ctx.render("user-profile", {
        title: ctx.user.name,
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

router.post("/sign-in", async (ctx) => {
    try {
        if (!ctx.user) {
            let { email, password } = ctx.request.fields;
            let token = await User.signIn(email, password);
            ctx.cookies.set("user", token, COOKIE_OPTIONS);
            await ctx.json({ code: RESULT.SUCCESS });
        }
    } catch (err) {
        await ctx.json(handleError(err));
    }
});

router.post("/register", async (ctx) => {
    try {
        if (!ctx.user) {
            let token = await User.register(ctx.request.fields);
            ctx.cookies.set("user", token, COOKIE_OPTIONS);
            await ctx.json({ code: RESULT.SUCCESS });
        }
    } catch(err) {
        if (err.isJoi) {
            await ctx.json({
                code: RESULT.INVALID_QUERY,
                details: err.details
            });
        } else {
            await ctx.json(handleError(err));
        }
    }
});

router.post("/recover", async (ctx) => {
    try {
        if (!ctx.user) {
            await User.sendPasswordRecoveryMail(ctx.request.fields.email);
            await ctx.json({ code: RESULT.SUCCESS });
        }
    } catch(err) {
        await ctx.json(handleError(err));
    }
});

router.post("/reset-password", async (ctx) => {
    try {
        if (ctx.request.fields.token) {
            await User.resetPasswordWithToken({
                password: ctx.request.fields.password,
                email: ctx.request.fields.email,
                token: ctx.request.fields.token
            });
            await ctx.json({ code: RESULT.SUCCESS });
        }
    } catch(err) {
        await ctx.json(handleError(err));
    }
})

router.post("/sign-in/facebook", async (ctx) => {
    try {
        let { access_token } = ctx.request.fields;
        let token = await User.signInWithFacebook(access_token);
        ctx.cookies.set("user", token, COOKIE_OPTIONS);
        await ctx.json({ code: RESULT.SUCCESS });
    } catch(err) {
        await ctx.json(handleError(err));
    }
});

router.post("/verify", async (ctx) => {
    try {
        await User.sendUserVerificationMail(ctx.user._id);
        await ctx.json({ code: RESULT.SUCCESS });
    } catch(err) {
        await ctx.json(handleError(err));
    }
});

router.post("/verify/complete", async (ctx) => {
    try {
        let { token } = ctx.request.fields;
        await User.verifyUserEmail(ctx.user._id, token);
        await ctx.json({ code: RESULT.SUCCESS });
    } catch(err) {
        await ctx.json(handleError(err));
    }
});

router.post("/connect/facebook", async (ctx) => {
    try {
        if (ctx.user) {
            let { access_token } = ctx.request.fields;
            await User.connectToFacebook(ctx.user._id, access_token);
            await ctx.json({ code: RESULT.SUCCESS });
        }
    } catch(err) {
        await ctx.json(handleError(err));
    }
});

router.post("/profile/update", async (ctx) => {
    try {
        await User.updateUserProfile(ctx.user._id, ctx.request.fields);
        await ctx.json({ code: RESULT.SUCCESS });
    } catch(err) {
        if (err.isJoi) {
            await ctx.json({
                code: RESULT.INVALID_QUERY,
                details: err.details
            });
        } else {
            await ctx.json(handleError(err));
        }
    }
});

module.exports = router;