import { setFormMessage, POST, getQueryVariable } from "./form";
import { initFacebookButton } from "./user-auth";
const RESULT = require("../../app/result-code");

function initLoginButton() {
    const submit = document.getElementById("submit"),
        email = document.getElementById("email"),
        password = document.getElementById("password");
    submit.addEventListener("click", async function () {
        if (!email.value || !password.value) {
            setFormMessage("請輸入電子郵件及密碼");
            return;
        }
        let response = await POST("/user/login", {
            email: email.value,
            password: password.value
        });
        let result = await response.json();
        if (result.code === RESULT.SUCCESS) {
            location.href = getQueryVariable("redirect_url") || "/";
        } else if (result.code === RESULT.INVALID_INPUT) {
            setFormMessage("電子郵件或密碼錯誤");
        } else if (result.code === RESULT.INVALID_TARGET) {
            setFormMessage("此帳戶未啟用以密碼登入");
        } else {
            setFormMessage(`伺服器發生錯誤（錯誤代碼：${result.code}）`);
        }
    });
}

onload = () => {
    initLoginButton();
    initFacebookButton();
};