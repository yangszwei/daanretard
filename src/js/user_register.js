import {
    setFormMessage,
    POST,
    getQueryVariable,
    initFacebookButton
} from "./user_auth";

function initRegisterButton() {
    const submit = document.getElementById("submit"),
        username = document.getElementById("username"),
        email = document.getElementById("email"),
        password = document.getElementById("password");
    submit.addEventListener("click", async function () {
        if (!username.value || !email.value || !password.value) {
            setFormMessage("請填寫所有欄位");
            return;
        }
        if (password.value.length < 6 || password.value.length > 30) {
            setFormMessage("密碼長度需介於6-30");
            return;
        }
        let response = await POST("/user/register", {
            name: username.value,
            email: email.value,
            password: password.value
        });
        let result = await response.json();
        if (result.status === "success") {
            location.href = getQueryVariable("redirect_url") || "/";
        } else if (result.reason === "Validation Failed") {
            setFormMessage("請填寫正確的資料");
        } else if (result.reason === "User Already Exists") {
            setFormMessage("此電子郵件已被使用");
        } else {
            setFormMessage("伺服器發生錯誤");
        }
    });
}

onload = () => {
    initRegisterButton();
    initFacebookButton();
};