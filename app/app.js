const db = require("./database");

class App {

    // 版規
    static get rules () {
        let app = db.collection("app");
        return app.findOne({ name: "rules" }) || { content: [] };
    }

}

module.exports = App;