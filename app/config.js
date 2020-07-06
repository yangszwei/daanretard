const db = require("./database");

class Config {

    // 版規
    static get rules () {
        let app = db.collection("configs");
        return app.findOne({ name: "posting rules" }) || { data: [] };
    }

}

module.exports = Config;