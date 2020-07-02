const Joi = require("joi"),
    bcrypt = require("bcrypt");
const db = require("./database"),
    graph = require("./graph");

class User {

    static async getUserByPassword(email, password) {
        let user = await db.collection("users").findOne({ email: email });
        if (!user) {
            throw "Invalid Credential";
        } else if (!user.providers.hasOwnProperty("password")) {
            throw "No Valid Provider";
        } else if (!bcrypt.compareSync(password, user.providers.password.password)) {
            throw "Invalid Credential";
        }
        return user;
    }

    static async continueWithFacebook(accessToken) {
        let facebookUser = await graph("GET /me", {
            fields: ["id", "name", "email"].join(","),
            access_token: accessToken
        });
        let users = db.collection("users");
        let user = await db.collection("users").findOne({
            "providers.facebook.id": facebookUser.id
        });
        if (!user) {
            if (await users.findOne({ email: facebookUser.email })) {
                // TODO: prompt add facebook to providers
                throw "User Already Exists";
            } else {
                return await this.createUserWithFacebook({
                    accessToken: accessToken,
                    ...facebookUser
                });
            }
        }
        return user;
    }

    static async createUserWithPassword(credential) {
        let validation = this.#validateRegistry(credential);
        if (validation.error) throw validation.error;
        if (await this.#userExistsWithEmail(credential.email)) {
            throw "User Already Exists";
        }
        return await db.collection("users").insertOne({
            name: credential.name,
            email: credential.email,
            providers: {
                password: {
                    password: bcrypt.hashSync(credential.password, 7)
                }
            }
        });
    }

    static async createUserWithFacebook(credential) {
        return await db.collection("users").insertOne({
            name: credential.name,
            email: credential.email,
            providers: {
                facebook: {
                    id: credential.id,
                    accessToken: credential.accessToken
                }
            }
        });
    }

    // TODO: add this to router
    static async addPasswordToProviders(_id, credential) {
        let schema = Joi.string().regex(/^[a-zA-Z0-9]{6,30}$/).required();
        Joi.validate(password, schema);
        return await db.collection("users").updateOne({ _id: _id }, {
            providers: {
                password: {
                    password: bcrypt.hashSync(credential.password, 7)
                }
            }
        });
    }

    // TODO: add this to router
    static async addFacebookToProviders(_id, credential) {
        let token = await graph.exchangeAccessToken(credential.access_token);
        return await db.collection("users").updateOne({ _id: _id }, {
            providers: {
                facebook: {
                    id: credential.id,
                    name: credential.name,
                    accessToken: token
                }
            }
        });
    }

    static async sendVerificationEmail(_id) {
        // TODO: add threshold
    }

    static #validateRegistry = (credential) => {
        const schema = {
            name: Joi.string().max(50).required(),
            password: Joi.string().regex(/^[a-z\d\-_\s]{6,30}$/i).required(),
            email: Joi.string().email().required()
        };
        return Joi.validate(credential, schema);
    }

    static #userExistsWithEmail = (email) => {
        return db.collection("users").findOne({ email: email });
    }

}

module.exports = User;