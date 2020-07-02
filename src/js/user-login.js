import { setFormMessage, POST, getQueryVariable } from "./form";
import { initFacebookButton } from "./user-auth";

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
        if (result.status === "success") {
            location.href = getQueryVariable("redirect_url") || "/";
        } else if (result.reason === "Invalid Credential") {
            setFormMessage("電子郵件或密碼錯誤");
        } else if (result.reason === "No Valid Provider") {
            setFormMessage("此帳戶無法使用密碼登入");
        } else {
            setFormMessage("伺服器發生錯誤");
        }
    });
}

onload = () => {
    initLoginButton();
    initFacebookButton();
};