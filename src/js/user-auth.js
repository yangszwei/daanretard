import { POST, getQueryVariable, setFormMessage } from "./form";
const RESULT = require("../../app/result-code");

function initFacebookButton() {
    const fbButton = document.getElementById("facebook");
    fbButton.addEventListener("click", () => {
        FB.login((fb_response) => {
            if (fb_response.status === "connected") {
                POST("/user/oauth/facebook", fb_response.authResponse).then((response) => {
                    response.json().then((result) => {
                        if (result.code === RESULT.SUCCESS) {
                            location.href = getQueryVariable("redirect_url") || "/";
                        } else if (
                            result.code === RESULT.DISABLED ||
                            result.code === RESULT.INVALID_TARGET
                        ) {
                            setFormMessage("此帳戶未啟用以Facebook登入");
                        } else {
                            setFormMessage(`伺服器發生錯誤（錯誤代碼：${result.code}）`);
                        }
                    });
                });
            } else {
                setFormMessage("Facebook登入失敗");
            }
        }, {
            scope: ["email", "public_profile"].join(","),
            return_scopes: true
        });
    });
}

export { initFacebookButton };