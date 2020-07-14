import {RESULT} from "./codes";

function validate(data) {
    if (!data.name || !data.email || !data.password) return "請填寫所有欄位";
    let len = data.password.length;
    if (len < 6 || len > 30) return "密碼長度須介於6-30";
}

addEventListener("load", () => {

    const nameInput = document.getElementById("name"),
        emailInput = document.getElementById("email"),
        passwordInput = document.getElementById("password"),
        submitButton = document.getElementById("submit"),
        submitResult = document.getElementById("result");

    submitButton.addEventListener("click", async () => {
        let error = validate({
            name: nameInput.value,
            email: emailInput.value,
            password: passwordInput.value
        });
        if (error) {
            submitResult.innerText = error;
            return;
        }
        let form = new FormData();
        form.append("name", nameInput.value);
        form.append("email", emailInput.value);
        form.append("password", passwordInput.value);
        let response = await fetch("/user/register", {
            method: "POST",
            body: form
        });
        let result = await response.json();
        if (result.code) {
            if (result.code === RESULT.ALREADY_EXIST) {
                submitResult.innerText = "電子郵件已經被使用";
            } else if(result.code === RESULT.INVALID_QUERY) {
                submitResult.innerText = `登入失敗(錯誤訊息：${result.details})`;
            } else {
                submitResult.innerText = `登入失敗(錯誤代碼：${result.code})`;
            }
        } else {
            location.href = "/";
        }
    });

});