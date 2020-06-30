function initLoginButton() {
    const loginButton = document.getElementById("login");
    loginButton.addEventListener("click", async () => {
        const email = document.getElementById("email"),
            password = document.getElementById("password");
        if (!email.value || !password.value) {
            if (!email.value) email.nextElementSibling.innerText = "請輸入正確的電子郵件";
            if (!password.value) password.nextElementSibling.innerText = "請輸入密碼";
        } else {
            let response = await (await fetch("/user/login", {
                method: "post",
                headers: {
                    'Accept': 'application/json, text/plain, */*',
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                    email: email.value,
                    password: password.value
                })
            })).json();
            if (response.status === "success") {
                location.href = "/";
            } else {
                if (response.reason === "Invalid Credential") {
                    loginButton.nextElementSibling.innerText = "登入失敗";
                } else {
                    loginButton.nextElementSibling.innerText = "登入失敗(伺服器錯誤)";
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
    initLoginButton();
    initFbLogin();
};