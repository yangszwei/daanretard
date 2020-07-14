async function errorPage(ctx, next) {
    try {
        await next();
        ctx.status = ctx.status || 404;
        if ([404, 405].includes(ctx.status)) ctx.throw(ctx.status);
    } catch(err) {
        ctx.status = err.status || 500;
        if ([404, 405].includes(ctx.status)) {
            await ctx.render("error-page", {
                user: ctx.user,
                title: 404,
                status: 404,
                content: "頁面不存在"
            });
        } else {
            console.error("An error was caught:", err);
            await ctx.render("error-page", {
                user: ctx.user,
                title: ctx.status,
                status: ctx.status,
                content: "伺服器發生錯誤"
            });
        }
    }
}
module.exports = () => errorPage;