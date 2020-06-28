function initModal() {
    let postBtn = document.getElementById("post"),
        menu = document.getElementById("menu"),
        prompt = document.getElementById("user_prompt");
    postBtn.addEventListener("click", (event) => {
        event.preventDefault();
        menu.style.display = "none";
        prompt.style.display = "flex";
    });
}

onload = () => {
    initModal();
};