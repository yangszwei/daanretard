const { security } = require("../config");
const jwt = require("jsonwebtoken");
const MILLISECONDS_IN_A_DAY = 60 * 60 * 24 * 1000;

module.exports = () => {
    return async (ctx, next) => {
        if (ctx.cookies.get("user")) {
            let user = ctx.cookies.get("user");
            ctx.user = jwt.verify(user, security.jwt);
            if (!ctx.cookies.get("updated")) {
                let token = jwt.sign(ctx.user, security.jwt);
                ctx.cookies.set("user", token, {
                    httpOnly: true,
                    signed: true,
                    maxAge: MILLISECONDS_IN_A_DAY * 30 * 2
                });
                ctx.cookies.set("updated", true, {
                    httpOnly: true,
                    maxAge: MILLISECONDS_IN_A_DAY
                });
            }
        }
        await next();
    };
};