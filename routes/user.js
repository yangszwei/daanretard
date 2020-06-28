const { external: { facebook } } = require("../config");
const Router = require("@koa/router"),
    got = require("got");
const router = new Router();
const graph = "https://graph.facebook.com/v7.0";


router.get("/auth", async (ctx) => {
    await ctx.render("auth", {
        title: "登入",
        appId: facebook.appId
    });
});

router.post("/auth/get-fb-profile", async (ctx) => {
    let access_token = ctx.request.body.access_token;
    if (!access_token) {
        ctx.body = JSON.stringify({
            status: "failed"
        });
    } else {
        let fields = "id,name,email";
        let url = `${graph}/me?fields=${fields}&access_token=${access_token}`;
        ctx.body = (await got(url)).body;
    }

});

module.exports = router;