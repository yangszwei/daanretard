const coll = require("./database");

class Post {

    constructor(data) {
        this.content = data.content;
    }



    async publish() {
        let posts = await coll("posts");
        posts.updateOne({ id: this.id }, {
            ...this.callback()
        });
    }

}

module.exports = Post;