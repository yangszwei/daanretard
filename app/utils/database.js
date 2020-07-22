const { database } = require("../../config").external;
const mongodb = require("mongodb"),
    sanitize = require("mongo-sanitize");
const { RESULT } = require("./codes");

class Collection {

    constructor(name, db) {
        this.collection = db.collection(name);
    }

    async insertOne(doc) {
        doc = Collection.sanitizeQuery(doc);
        return this.collection.insertOne(doc);
    }

    async upsertOne(filter, doc) {
        filter = Collection.sanitizeQuery(filter);
        doc = Collection.sanitizeQuery(doc);
        return this.collection.updateOne(filter, doc, { upsert: true });
    }

    async list(filter) {

    }

    async findOne(filter) {
        filter = Collection.sanitizeQuery(filter);
        let doc = await this.collection.findOne(filter);
        if (!doc) throw RESULT.NOT_EXIST;
        return doc;
    }

    async updateOne(filter, doc) {
        filter = Collection.sanitizeQuery(filter);
        if (doc.$set) doc.$set = Collection.sanitizeQuery(doc.$set);
        if (doc.$unset) doc.$unset = Collection.sanitizeQuery(doc.$unset);
        return this.collection.updateOne(filter, doc);
    }

    async deleteOne(filter) {
        filter = Collection.sanitizeQuery(filter);
        return this.collection.deleteOne(filter);
    }

    static sanitizeQuery(query) {
        if ("_id" in query) {
            query._id = mongodb.ObjectID(query._id);
        }
        return sanitize(query);
    }

}

class Database {

    #collections = {};

    constructor() {
        this.client = new mongodb.MongoClient(this.#getDatabaseUri(), {
            useUnifiedTopology: true,
            useNewUrlParser: true
        });
    }

    async connect() {
        await this.client.connect();
        this.db = this.client.db(database.name);
        return this.db;
    }

    collection(name) {
        if (!this.#collections[name]) {
            this.#collections[name] = new Collection(name, this.db);
        }
        return this.#collections[name];
    }

    #getDatabaseUri = () => {
        let { user, password, host, port } = database;
        return `mongodb://${user}:${password}@${host}:${port}`;
    }

}

module.exports = new Database();