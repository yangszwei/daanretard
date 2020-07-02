function toggleMenu() {
    document.querySelector("body > nav .toggle-menu")
        .addEventListener("click", () => {
            document.querySelector("body > nav").classList.toggle("active");
        });
}

addEventListener("load", event => {
    toggleMenu();
});