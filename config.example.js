const config = {
    app: {
        version: "0.2.0",
        url: "",
        mailAddress: ""
    },
    development: true,
    server: {
        port: 8000,
        proxy: true
    },
    security: {
        secret: "",
        keys: [ "" ]
    },
    external: {
        mail: {
            // directly passed to nodemailer.createTransport
        },
        database: {
            // mongodb
            user: "",
            password: "",
            name: "",
            host: "localhost",
            port: 27017
        },
        facebook: {
            pageId: "",
            appId: "",
            secret: ""
        },
        browserSync: {
            // passed directly to browserSync.init
        }
    }
};

module.exports = config;