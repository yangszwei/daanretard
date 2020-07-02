const { security } = require("../config");
const jwt = require("jsonwebtoken");

module.exports = () => {
    return async (ctx, next) => {
        if (ctx.cookies.get("user")) {
            ctx.user = jwt.verify(ctx.cookies.get("user"), security.jwt);
        }
        await next();
    };
};