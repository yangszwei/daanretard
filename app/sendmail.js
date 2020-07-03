const { smtp } = require("../config").external;
const nodemailer = require("nodemailer");
const transporter = nodemailer.createTransport(smtp);

module.exports = (mail) => {
    return transporter.sendMail({
        from: "靠北大安4.0 <no-reply@daan-retard.yncc.nctu.me>",
        to: mail.to,
        subject: mail.subject,
        html: mail.content
    });
};