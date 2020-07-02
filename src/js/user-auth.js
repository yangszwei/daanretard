import { POST, getQueryVariable, setFormMessage } from "./form";

function initFacebookButton() {
    const fbButton = document.getElementById("facebook");
    fbButton.addEventListener("click", () => {
        FB.login((fb_response) => {
            if (fb_response.status === "connected") {
                POST("/user/oauth/facebook", fb_response.authResponse).then((response) => {
                    response.json().then((result) => {
                        if (result.status === "success") {
                            location.href = getQueryVariable("redirect_url") || "/";
                        } else {
                            setFormMessage("伺服器發生錯誤");
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