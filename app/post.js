const Facebook = require("./utils/facebook"),
    { RESULT } = require("./utils/codes"),
    Resource = require("./resource"),
    database = require("./utils/database"),
    crypto = require("crypto"),
    BSON = require("bson");
const posts = database.collection("posts");

function validatePost(post) {
    if (post.content && post.content.length) {
        if (post.content.length > 500) throw "content too long";
    }
    let images = [];
    if (post.images) {
        for (let image of post.images) {
            if (image.size > 4 * 1024 * 1024) throw "image too large";
            images.push(BSON.serialize(image));
        }
    }
    post.images = images;
    return post;
}

class Post {

    static async submit(post, author) {
        let id = crypto.randomBytes(12).toString("base64");
        post = validatePost(post);
        const cleanPost = {
            id: id,
            stage: "submission",
            author: (author && author._id) || null,
            content: post.content,
            images: post.images,
            submit_time: Date.now()
        };
        let isValidAuthor = author && author._id;
        cleanPost.verified = Boolean(isValidAuthor);
        if (!isValidAuthor) cleanPost.email = post.email;
        await posts.insertOne(cleanPost);
        return id;
    }

    static async review(id, review) {
        let stage = "submission";
        let post = this.getSubmissionById(id);
        if (!post.verified) throw RESULT.UNAUTHORIZED;
        if (post.review && post.review.result) throw RESULT.ALREADY_EXIST;
        if (review.result === "approved") {
            stage = "pending post";
        } else if (review.result === "rejected") {
            stage = "rejected";
        } else {
            stage = "reviewed";
        }
        return posts.updateOne({
            $setOrInsert: {
                stage: stage,
                review: {
                    result: review.result,
                    comment: review.comment || "",
                    reviewer: review.reviewer._id,
                    timestamp: Date.now()
                }
            }
        });
    }

    static async publish(id) {
        let post = await this.getSubmissionById();
        if (post.stage !== "pending post" ||
            post.review.result !== "approved") throw RESULT.UNAUTHORIZED;
        if (post.images) {
            let urls = await Resource.uploadMultiple();
            post.attachMedia = await Facebook.uploadImages(urls);
        }
        let response = await Facebook.publishPost(post);
        await posts.updateOne({ id: id }, {
            $setOrInsert: {
                stage: "published",
                post_id: "000001",
                fb_post_id: response, // TODO
                publish_time: Date.now()
            }
        });
    }

    static getSubmissionById(id) {
        return posts.findOne({ id: id });
    }

}

module.exports = Post;