const coll = require("./database");

class User {

    constructor(cred) {
        this.name = cred.name;
    }

    static async login(id, password) {
        let users = await coll("users");
        let profile = await users.findOne({ id: id });
        if (!profile) throw "User not found";
    }

    static async register(cred) {
        let users = await coll("users");
        let profile = new this(cred);
        profile.save();
    }

    static async recover() {
        let users = await coll("users");
    }

}

module.exports = User;