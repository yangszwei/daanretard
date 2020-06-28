async function getFbUserProfile(access_token) {
    return await fetch("/user/get-fb-profile", {
        body: JSON.stringify({ access_token: access_token })
    });
}

function listenLoginResponse() {
    FB.Event.subscribe('auth.authResponseChange', async (response) => {
        if (response.status === "connected") {
            let accessToken = response.authResponse.access_token;
            let profile = await  getFbUserProfile(accessToken);
            console.log(profile)
        }
    });
}

onload = () => {
    listenLoginResponse();
}