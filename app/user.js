const { security } = require("../config");
const Joi = require("joi"),
    bcrypt = require("bcrypt"),
    jwt = require("jsonwebtoken");
const db = require("./database"),
    graph = require("./graph"),
    RESULT = require("./result-code"),
    sendmail = require("./sendmail");

const USER_DEFAULT_PREFERENCES = {};

class User {

    // LOG IN

    static async loginWithPassword(email, password) {
        let user = await this.#getUserByEmail(email);
        if (!user) throw RESULT.NOT_FOUND;
        if (!user.keys.hasOwnProperty("password")) throw RESULT.INVALID_TARGET;
        if (!user.keys.password.enabled) throw RESULT.DISABLED;
        let hash = user.keys.password.data;
        if (!bcrypt.compareSync(password, hash)) throw RESULT.INVALID_INPUT;
        return user;
    }

    static async continueWithFacebook(accessToken) {
        let credential = await this.#getFacebookProfile(accessToken);
        if (!credential.hasOwnProperty("id")) throw RESULT.INVALID_INPUT;
        let user = await this.#getUserByFacebookId(credential.id);
        if (!user) {
            credential.accessToken = accessToken;
            user = await this.#getUserByEmail(credential.email);
            if (!user) return this.registerWithFacebook(credential);
            // TODO: prompt to connect facebook.
            throw RESULT.INVALID_TARGET;
        }
        if (!user.keys.facebook.enabled) throw RESULT.DISABLED;
        return user;
    }

    // REGISTER

    static async registerWithPassword(credential) {
        let validation = this.#validateRegistration(credential);
        if (validation.error) throw validation.error;
        let user = await User.#getUserByEmail(credential.email);
        if (user) throw RESULT.ALREADY_EXIST;
        let result = await db.collection("users").insertOne({
            name: credential.name,
            email: credential.email,
            verified: false,
            keys: {
                password: {
                    enabled: true,
                    data: bcrypt.hashSync(credential.password, 7)
                }
            },
            apps: {},
            preferences: USER_DEFAULT_PREFERENCES
        });
        return result.ops[0];
    }

    static async registerWithFacebook(credential) {
        if (
            !credential.hasOwnProperty("name") ||
            !credential.hasOwnProperty("email")
        ) throw RESULT.INVALID_INPUT;
        let { access_token } = await graph.exchangeAccessToken(credential.accessToken);
        let result = await db.collection("users").insertOne({
            name: credential.name,
            email: credential.email,
            verified: true,
            keys: {
                facebook: {
                    enabled: true,
                    data: credential.id
                }
            },
            apps: {
                facebook: {
                    enabled: true,
                    name: credential.name,
                    email: credential.email,
                    accessToken: access_token
                }
            },
            preferences: USER_DEFAULT_PREFERENCES
        });
        return result.ops[0];
    }

    // UPDATE PROFILE

    static async updateUserProfile(credential) {
        let validation = this.#validateRegistration(credential);
        if (validation.error) throw validation.error;
        let user = await User.#getUserByEmail(credential.email);
        if (!user) throw RESULT.NOT_FOUND;
        let update = {};
        if (credential.name) update.name = credential.name;
        if (credential.email) {
            update.email = credential.email;
            update.verified = (user.apps.facebook) &&
                (user.apps.facebook.email === update.email);
        }
        let result = await db.collection("users").updateOne({

        }, {
            name: credential.name,
            email: credential.email
        });
        return result.ops[0];
    }

    // VERIFY

    static async verifyUserByEmail(_id) {}

    // GET LOCAL PROFILE

    static #getUserByEmail = (email) => {
        return this.#getUserByFilter({ email: email });
    }

    static #getUserByFacebookId = (id) => {
        return this.#getUserByFilter({
            "keys.facebook.data": id
        });
    }

    static getUserByObjectID (_id) {
        return this.#getUserByFilter({ _id: db._id(_id) });
    }

    static #getUserByFilter = async (filter) => {
        let user = db.collection("users").findOne(filter);
        if (!user) throw RESULT.NOT_FOUND;
        return user;
    }

    // GET APP USER PROFILE

    static #getFacebookProfile = (accessToken) => {
        return graph("GET /me", {
            fields: ["id", "name", "email"].join(","),
            access_token: accessToken
        })
    }

    // TOKEN

    static signUserToken(user) {
        return jwt.sign({
            id: user._id,
            name: user.name,
            email: user.email,
            verified: user.verified,
            privileges: user.privileges
        }, security.jwt);
    }

    // INPUT VALIDATION

    static #validateRegistration = (credential) => {
        const schema = {
            name: Joi.string().max(50).required(),
            password: Joi.string().regex(/^[a-z\d\-_\s]{6,30}$/i).required(),
            email: Joi.string().email().required()
        };
        return Joi.validate(credential, schema);
    }

}

module.exports = User;