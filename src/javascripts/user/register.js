function initRegisterButton() {
    const registerButton = document.getElementById("register");
    registerButton.addEventListener("click", async () => {
        const name = document.getElementById("name"),
            email = document.getElementById("email"),
            password = document.getElementById("password");
        if (!name.value || !email.value || !password.value) {
            if (!name.value) name.nextElementSibling.innerText = "請輸入姓名";
            if (!email.value) email.nextElementSibling.innerText = "請輸入正確的電子郵件";
            if (!password.value) password.nextElementSibling.innerText = "請輸入密碼";
        } else {
            let response = await (await fetch("/user/register", {
                method: "post",
                headers: {
                    'Accept': 'application/json, text/plain, */*',
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                    name: name.value,
                    email: email.value,
                    password: password.value
                })
            })).json();
            if (response.status === "success") {
                location.href = "/";
            } else {
                if (response.reason === "Validation Failed") {
                    // TODO: add response.details to corresponding form-control.status
                    registerButton.nextElementSibling.innerText = "無效的使用者資料";
                } else if (response.reason === "User Already Exists") {
                    registerButton.nextElementSibling.innerText = "此電子郵件已使用";
                } else {
                    registerButton.nextElementSibling.innerText = "伺服器錯誤";
                }
            }
        }

    });
}

function initFbLogin() {
    FB.Event.subscribe('auth.authResponseChange', async (response) => {
        if (response.status === "connected") {
            let profile = await (await fetch("/user/oauth/facebook", {
                method: "post",
                headers: {
                    'Accept': 'application/json, text/plain, */*',
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(response.authResponse)
            })).json();
            console.log(profile)
        } else {
            const loginButton = document.getElementById("login");
            loginButton.nextElementSibling.innerText = "Facebook登入失敗";
        }
    });
}

onload = () => {
    initRegisterButton();
    initFbLogin();
};