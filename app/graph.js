const { facebook } = require("../config").external;
const got = require("got");
const path = require("path");
const graphUrl = "https://graph.facebook.com/v7.0";

async function graph (command, query) {
    let [ method, target ] = command.split(" ");
    method = method.toLowerCase();
    query = Object.keys(query).map((key) => {
        let value = encodeURIComponent(query[key]);
        key = encodeURIComponent(key);
        return `${key}=${value}`;
    }).join('&');
    return JSON.parse((await got[method](`${path.join(graphUrl, target)}?${query}`)).body);
}

graph.exchangeAccessToken = (accessToken) => {
    return graph("GET /oauth/access_token", {
        grant_type: "fb_exchange_token",
        client_id: facebook.appId,
        client_secret: facebook.secret,
        fb_exchange_token: accessToken
    });
};

module.exports = graph;