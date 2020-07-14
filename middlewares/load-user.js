const { security } = require("../config");
const jwt = require("jsonwebtoken");
const MILLISECONDS_IN_A_DAY = 60 * 60 * 24 * 1000;

async function loadUser(ctx, next) {
    let token = ctx.cookies.get("user"),
        last_refresh = ctx.cookies.get("last_refresh");
    if (token) {
        ctx.user = jwt.verify(token, security.secret);
        let today = new Date().setHours(0,0,0,0);
        if (!last_refresh || (last_refresh < today)) {
            ctx.cookies.set("last_refresh", Date.now());
            ctx.cookies.set("user", jwt.sign(ctx.user, security.secret), {
                httpOnly: true,
                signed: true,
                maxAge: MILLISECONDS_IN_A_DAY * 60
            });
        }
    }
    await next();
}

module.exports = () => loadUser;