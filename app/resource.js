const { app } = require("../config");
const BSON = require("bson"),
    crypto = require("crypto");
const database = require("./utils/database");
const resources = database.collection("resources");
class Resource {

    static async upload(file) {
        let id = crypto.randomBytes(8).toString("base64");
        await resources.insertOne({
            id: id,
            file: BSON.serialize(file)
        });
        return app.domain + "/resource/" + id;
    }

    static async uploadMultiple(files) {
        let responses = [];
        for (let file of files) {
            responses.push(await this.upload(file));
        }
        return responses;
    }

    static async read(id) {
        let doc = await resources.findOne({ id: id });
        return BSON.deserialize(doc.file.buffer);
    }

}

module.exports = Resource;