const got = require("got"),
    path = require("path");
const { facebook } = require("../../config").external;

class GraphRequest {

    static host = "https://graph.facebook.com/v7.0";

    constructor(query) {
        this.url = path.join(this.constructor.host, query.path);
        this.query = query.body;
        this.method = query.method.toLowerCase() || "GET";
    }

    async send() {
        let query = this.constructor.stringifyQuery(this.query);
        console.log(`${this.url}?${query}`)
        let response = await got[this.method](`${this.url}?${query}`);
        return JSON.parse(response.body);
    }

    static send(query) {
        let request = new this(query);
        return request.send();
    }

    static stringifyQuery(query) {
        return Object.keys(query).map((key) => {
            let value = encodeURIComponent(query[key]);
            key = encodeURIComponent(key);
            return `${key}=${value}`;
        }).join('&');
    }

}

class Facebook {

    static api(path, options) {
        let request = new GraphRequest(path, options);
        return request.send();
    }

    static getUserProfile(accessToken) {
        return GraphRequest.send({
            path: "/me",
            method: "GET",
            body: {
                fields: ["id", "name", "email"].join(","),
                access_token: accessToken
            }
        });
    }

    static exchangeLongLivedToken(accessToken) {
        return GraphRequest.send({
            path: "/oauth/access_token",
            method: "GET",
            body: {
                grant_type: "fb_exchange_token",
                client_id: facebook.appId,
                client_secret: facebook.secret,
                fb_exchange_token: accessToken
            }
        });
    }

    static publishPost(post) {
        return GraphRequest.send({
            path: `/${facebook.pageId}/feed`,
            method: "POST",
            body: {
                message: post.content,
                published: true,
                ...this.attachMedia(post.attachMedia),
                access_token: facebook.accessToken
            }
        });
    }

    static async uploadImages(urls) {
        let responses = [];
        for (let url of urls) {
            let response = await GraphRequest.send({
                path: `/${facebook.pageId}/photos`,
                method: "POST",
                body: {
                    url: url,
                    published: false,
                    access_token: facebook.accessToken
                }
            });
            console.log(response);
            responses.push(response);
        }
        return responses;
    }

    static attachMedia(ids) {
        let obj = {};
        ids.forEach((value, index) => {
            obj[`attached_media[${index}]`] = `{"media_fbid": "${value}"}`;
        });
        return obj;
    }

    static publishComment() {}

    static listComments() {}

    static deletePost() {}

}

module.exports = Facebook;