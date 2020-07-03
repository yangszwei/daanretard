const { security } = require("../config");
const jwt = require("jsonwebtoken");
const SECONDS_IN_A_MONTH = 60 * 60 * 24 * 30;

module.exports = () => {
    return async (ctx, next) => {
        if (ctx.cookies.get("user")) {
            let user = ctx.cookies.get("user");
            ctx.user = jwt.verify(user, security.jwt);
            let token = jwt.sign({ id: user._id }, security.jwt);
            ctx.cookies.set("user", token, {
                httpOnly: true,
                signed: true,
                maxAge: SECONDS_IN_A_MONTH * 2
            });
        }
        await next();
    };
};