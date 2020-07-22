const database = require("./utils/database");
const configs = database.collection("configs");

class Config {

    constructor(name) {
        this.name = name;
    }

    async set(data) {
        return configs.upsertOne({ name: this.name }, {
            name: this.name,
            data: data
        });
    }

    async get() {
        let doc = await configs.findOne({ name: this.name });
        return doc.data;
    }

}

class Configs {

    // posting rules

    static async getPostingRules() {
        return (await new Config("posting rules").get()) || [];
    }

    static async setPostingRules(data) {
        return await new Config("posting rules").set(data || []);
    }

    // user verification mail template

    static async getUserVerificationMailTemplate() {
        let config = new Config("user verification mail");
        return await config.get() || [];
    }

    static async setUserVerificationMailTemplate(data) {
        let config = new Config("user verification mail");
        return await config.set(data);
    }

    // post - email verification template
    static async getEmailVerificationTemplate() {
        let config = new Config("post email verification mail");
        return await config.get() || [];
    }

    static async setEmailVerificationMailTemplate(data) {
        let config = new Config("post email verification mail");
        return await config.set(data);
    }

    // password recovery mail template

    static async getPasswordRecoveryMailTemplate() {
        let config = new Config("password recovery mail");
        return await config.get() || { subject: "", content: "" };
    }

    static async setPasswordRecoveryMailTemplate(data) {
        let config = new Config("password recovery mail");
        return await config.set(data);
    }

    static async getPostTemplate() {
        let config = new Config("password recovery mail");
        return await config.set(data);
    }

    static async setPostTemplate(data) {
        let config = new Config("password recovery mail");
        return await config.get();
    }

    static fillTemplate(template, params) {
        Object.keys(params).forEach((index) => {
            template = template.replaceAll(`{{${index}}}`, params[index]);
        });
        return template;
    }

}

module.exports = Configs;