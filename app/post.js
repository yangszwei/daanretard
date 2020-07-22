const Facebook = require("./utils/facebook"),
    { RESULT } = require("./utils/codes"),
    database = require("./utils/database"),
    mail = require("./utils/mail"),
    Config = require("./config"),
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

function generateToken() {
    return Math.floor(100000 + Math.random() * 900000).toString();
}

class Post {

    static async submit(post, author) {
        let id = crypto.randomBytes(12).toString("base64");
        post = validatePost(post);
        const cleanPost = {
            id: id,
            stage: "submission",
            content: post.content,
            images: post.images,
            submit_time: Date.now()
        }
        if (author) {
            cleanPost.author = author._id;
            cleanPost.verified = true;
            await posts.insertOne(cleanPost);
        } else {
            let filter = {
                stage: "not submitted",
                email: post.email,
                verified: true
            };
            console.log(await posts.findOne(filter))
            if (await posts.findOne(filter)) {
                await posts.updateOne(filter, { $set: cleanPost });
            } else {
                throw RESULT.NOT_EXIST;
            }
        }
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

    static async listNotReviewed() {

    }

    static async publish(id) {
        let post = await this.getSubmissionById();
        if (post.stage !== "pending post" ||
            post.review.result !== "approved") throw RESULT.UNAUTHORIZED;
        if (post.images) {
            // let urls = await Resource.uploadMultiple();
            post.attachMedia = await Facebook.uploadImages(urls);
        }
        let response = await Facebook.publishPost({

        });
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

    static async sendVerificationEmail(email) {
        let token = generateToken();
        if (!email) throw RESULT.INVALID_QUERY;
        const post = {
            stage: "not submitted",
            email: email,
            token: token,
            verified: false
        };
        try {
            let filter = {
                email: email,
                stage: "not submitted"
            }
            await posts.findOne(filter);
            await posts.deleteOne(filter);
        } catch(err) {}
        await posts.insertOne(post);
        let template = await Config.getEmailVerificationTemplate();
        let params = { token: token };
        await mail.send({
            to: email,
            subject: Config.fillTemplate(template.subject, params),
            content: Config.fillTemplate(template.content, params)
        });
    }

    static async verifyEmail(token, email) {
        let filter = {
            email: email,
            token: token,
            stage: "not submitted"
        }
        if (!await posts.findOne(filter)) {
            throw RESULT.INVALID_QUERY;
        }
        await posts.updateOne(filter, {
            $set: { verified: true },
            $unset: { token: "" }
        });
    }

}

module.exports = Post;