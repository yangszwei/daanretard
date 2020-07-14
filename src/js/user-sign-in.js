import { RESULT } from "./codes";

function validate(data) {
    if (!data.email || !data.password) return "電子郵件或密碼錯誤";
}

addEventListener("load", () => {

    const emailInput = document.getElementById("email"),
        passwordInput = document.getElementById("password"),
        submitButton = document.getElementById("submit"),
        submitResult = document.getElementById("result");

    submitButton.addEventListener("click", async () => {
        let error = validate({
            email: emailInput.value,
            password: passwordInput.value
        });
        if (error) {
            submitResult.innerText = error;
            return;
        }
        let form = new FormData();
        form.append("email", emailInput.value);
        form.append("password", passwordInput.value);
        let response = await fetch("/user/sign-in", {
            method: "POST",
            body: form
        });
        let result = await response.json();
        if (result.code) {
            if (result.code === RESULT.INVALID_QUERY) {
                submitResult.innerText = "電子郵件或密碼錯誤";
            } else {
                submitResult.innerText = `登入失敗(錯誤代碼：${result.code})`;
            }
        } else {
            location.href = "/";
        }
    });

});