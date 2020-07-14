const { app, external } = require("../../config");
const nodemailer = require("nodemailer");
const transporter = nodemailer.createTransport(external.mail);

function sendMail(mail) {
    return transporter.sendMail({
        from: app.mailAddress,
        to: mail.to,
        subject: mail.subject,
        html: mail.content
    });
}

module.exports = { send: sendMail };