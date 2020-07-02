const BSON = require('bson');
const db = require("./database");
const bson = new BSON();

class PostSubmission {

    static create(post) {
        return db.collection("submissions").insertOne({
            author: post.author || "email",
            content: post.content,
            media: post.media.map((file) => bson.serialize(file)),
            ...(post.email ? { email: post.email } : {})
        });
    }

    verify() {

    }

}

module.exports = PostSubmission;