async function comingSoon(ctx, next) {
    if (
        !ctx.cookies.get("developer") &&
        (ctx.url !== "/coming-soon")
    ) ctx.redirect("/coming-soon");
    await next();
}
module.exports = () => comingSoon;