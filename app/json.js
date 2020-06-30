module.exports = () => {
    return async (ctx, next) => {
        ctx.json = (object) => {
            ctx.body = JSON.stringify(object);
        };
        await next();
    };
};