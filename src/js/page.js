addEventListener("load", function initPage() {

    const topnav = document.getElementById("topnav"),
        menu = document.querySelector("#topnav .nav-menus"),
        menuToggle = document.getElementById("topnav-toggle-menu");

    // init menu toggle
    menuToggle.addEventListener("click", function () {
        let previousState = this.classList.contains("active");
        this.classList.toggle("active");
        menu.classList[previousState ? "remove" : "add"]("active");
    });

    // close menu on touch outside
    addEventListener("touchstart", (event) => {
        if ((topnav !== event.target) && !topnav.contains(event.target)) {
            menuToggle.classList.remove("active");
            menu.classList.remove("active");
        }
    })

    // init topnav shadow
    addEventListener("scroll", () => {
        let scrollTop = pageYOffset ||
            document.body.scrollTop ||
            document.documentElement.scrollTop;
        topnav.classList[(scrollTop > 0) ? "add" : "remove"]("shadow");
    });

});