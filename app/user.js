const { security } = require("../config");
const Joi = require("joi"),
    bcrypt = require("bcrypt"),
    jwt = require("jsonwebtoken");
const database = require("./utils/database"),
    config = require("./config"),
    facebook = require("./utils/facebook"),
    { RESULT } = require("./utils/codes"),
    mail = require("./utils/mail");
const users = database.collection("users");

function generateToken() {
    return Math.floor(100000 + Math.random() * 900000).toString();
}

function validateCredentials(credentials) {
    const schema = {
        name: Joi.string().max(50).required(),
        password: Joi.string().regex(/^[a-z\d\-_\s]{6,30}$/i).required(),
        email: Joi.string().email().required()
    };
    return Joi.validate(credentials, schema);
}

function validateCredentialsUpdate(credentials) {
    const schema = {
        name: Joi.string().max(50),
        email: Joi.string().email()
    };
    return Joi.validate(credentials, schema);
}

class User {

    // Sign In

    static async signIn(email, password) {
        let user = await this.getUserByEmail(email);
        if (!user.password) throw RESULT.INACCESSIBLE;
        if (!bcrypt.compareSync(password, user.password)) {
            throw RESULT.INVALID_QUERY;
        }
        return this.signUserToken(user);
    }

    static async signInWithFacebook(accessToken) {
        let profile = await this.getFbProfile(accessToken);
        if (!profile) throw RESULT.INVALID_QUERY;
        let user = await this.getUserByFacebookId(profile.id);
        if (user) return user;
        profile.accessToken = accessToken;
        user = await this.getUserByEmail(profile.email);
        if (!user) throw RESULT.FORWARD;
        return this.registerWithFacebook(profile);
    }

    // Register

    static async register(credentials) {
        let validation = validateCredentials(credentials);
        if (validation.error) throw validation.error;
        try {
            await User.getUserByEmail(credentials.email);
            throw RESULT.ALREADY_EXIST;
        } catch (err) {
            if (err === RESULT.NOT_EXIST) {
                let result = await users.insertOne({
                    name: credentials.name,
                    email: credentials.email,
                    verified: false,
                    password: bcrypt.hashSync(credentials.password, 7)
                });
                return this.signUserToken(result.ops[0]);
            } else if (typeof err === "number") {
                throw err;
            }
        }
    }

    static async registerWithFacebook(credentials) {
        if (
            !credentials.hasOwnProperty("name") ||
            !credentials.hasOwnProperty("email")
        ) throw RESULT.INVALID_INPUT;
        let exchange = await facebook.exchangeLongLivedToken(credentials.accessToken);
        let result = await users.insertOne({
            name: credentials.name,
            email: credentials.email,
            verified: true,
            fb_token: exchange.access_token
        });
        return result.ops[0];
    }

    // Update User Profile

    static async updateUserProfile(_id, credentials) {
        let validation = validateCredentialsUpdate(credentials);
        if (validation.error) throw  validation.error;
        let update = {},
            user = await User.getUserByObjectID(_id);
        if ((user.email !== credentials.email) && credentials.email) {
            if (await this.getUserByEmail(credentials.email)) {
                throw RESULT.ALREADY_EXIST;
            }
        }
        if (credentials.name) update.name = credentials.name;
        if (credentials.email) {
            update.email = credentials.email;
            if (user.fb_token) {
                let profile = await this.getFbProfile(user.fb_token);
                update.verified = Boolean(credentials.email === profile.email);
            }
        }
        return users.updateOne({ _id: _id }, { $set: update });
    }

    // User Verification

    static async sendUserVerificationMail(_id) {
        let { email } = await this.getUserByObjectID(_id);
        let template = await config.getUserVerificationMailTemplate();
        let token = generateToken();
        await users.updateOne({ _id: _id }, { $set: { verification: token } });
        let params = { token: token };
        await mail.send({
            to: email,
            subject: config.fillTemplate(template.subject, params),
            content: config.fillTemplate(template.content, params)
        });
    }

    static async verifyUserEmail(_id, token) {
        let user = await this.getUserByObjectID(_id);
        if (user.verification !== token) throw RESULT.INVALID_QUERY;
        return users.updateOne({ _id: _id }, {
            $set: { verified: true },
            $unset: { verification: "" }
        });
    }

    // Recover Account

    static async sendPasswordRecoveryMail(email) {
        let template = await config.getPasswordRecoveryMailTemplate();
        let token = generateToken();
        await users.updateOne({ email: email }, { $set: { recovery: token } });
        let params = { email: email, token: token };
        await mail.send({
            to: email,
            subject: config.fillTemplate(template.subject, params),
            content: config.fillTemplate(template.content, params)
        });
    }

    static async resetPasswordWithToken(query) {
        let user = await User.getUserByEmail(query.email);
        if (!user.recovery) throw RESULT.NOT_ENABLED;
        if (user.recovery !== query.token) throw RESULT.INVALID_QUERY;
        if (!query.password) throw RESULT.INVALID_QUERY;
        await users.updateOne({ email: query.email }, {
            $set: { password: bcrypt.hashSync(query.password, 7) },
            $unset: { recovery: "" }
        });
    }

    // Get Local User

    static getUserByEmail = (email) => {
        return users.findOne({ email: email });
    }

    static getUserByFacebookId = (id) => {
        return users.findOne({ "keys.facebook.data": id });
    }

    static getUserByObjectID (_id) {
        return users.findOne({ _id: _id });
    }

    // Facebook

    static connectToFacebook(_id, accessToken) {
        return users.updateOne({ _id: _id }, {
            $set: { fb_token: accessToken }
        });
    }

    // Get Facebook Profile

    static getFbProfile(accessToken) {
        return facebook.getUserProfile(accessToken);
    }

    // Sign Tokens

    static signUserToken(user) {
        return jwt.sign({
            _id: user._id,
            email: user.email,
            name: user.name || "",
            verified: user.verified || false,
            fb_token: user.fb_token || null,
        }, security.secret);
    }

}

module.exports = User;