const BSON = require('bson');
const db = require("./database");
const sendmail = require("./sendmail");

class PostSubmission {

    static async create(post) {
        await db.collection("submissions").insertOne({
            author: post.author || "email",
            content: post.content,
            media: post.media.map((file) => BSON.serialize(file)),
            ...(post.email ? { email: post.email } : {})
        });
    }

}

module.exports = PostSubmission;