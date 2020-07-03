const { database } = require("../config").external;
const mongodb = require("mongodb");
const client = new mongodb.MongoClient(getDatabaseUri(), {
    useNewUrlParser: true,
    useUnifiedTopology: true
});

function getDatabaseUri() {
    const { user, password, host, port } = database;
    return `mongodb://${user}:${password}@${host}:${port}`;
}

module.exports = {
    client: client,
    _id: mongodb.ObjectID,
    connect: async () => {
        return await client.connect();
    },
    collection: (name) => {
        return client.db(database.name).collection(name);
    }
};