const config = {
    app: {
        version: ""
    },
    server: {
        port: 8000,
        proxy: true
    },
    security: {
        jwt: "",
        keys: [ "" ]
    },
    external: {
        database: {
            // mongodb
            user: "",
            password: "",
            name: "",
            host: "localhost",
            port: 27017
        },
        facebook: {
            appId: "",
            secret: ""
        },
        browserSync: {
            // browserSync config here
        }
    }
};

module.exports = config;