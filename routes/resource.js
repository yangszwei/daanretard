const Router = require("@koa/router"),
    send = require("koa-send");
const router = new Router();
const Resource = require("../app/resource");

router.get("/:id", async (ctx) => {
    let file = await Resource.read(ctx.params.id);
    ctx.set("Content-disposition", `inline; filename=${escape(file.name)}`);
    ctx.set("Content-type", file.type);
    await send(ctx, file.path, { root: "/" });
});

module.exports = router;