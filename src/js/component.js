function TopNav() {
    const topnav = document.getElementById("topnav"),
        navigation = document.querySelector("#topnav .navigation"),
        socialLinks = document.querySelector("#topnav .social-links"),
        menuToggle = document.querySelector("#topnav .menu-toggle"),
        socialLinksToggle = document.querySelector("#topnav a.title");
    (function initMenuToggle() {
        menuToggle.onclick = function (event) {
            event.stopPropagation();
            this.classList.toggle("active");
            socialLinks.classList.remove("active")
            navigation.classList.toggle("active");
        };
    })();
    (function initSocialLinksToggle() {
        socialLinksToggle.onclick = function (event) {
            if (matchMedia("(max-width: 480px)").matches) {
                event.preventDefault();
                this.classList.toggle("active");
                menuToggle.classList.remove("active");
                navigation.classList.remove("active");
                socialLinks.classList.toggle("active");
            }
        };
        let timer;
        socialLinksToggle.ontouchstart = function (event) {
            timer = setTimeout(() => {
                location.href = "/";
            }, 500);
        }
        socialLinksToggle.ontouchend = function () {
            if (timer) clearTimeout(timer);
        };
    })();
    (function closeMenuOnClickOutside() {
        function closeMenu(event) {
            if (!(event.target === topnav) && !topnav.contains(event.target)) {
                navigation.classList.remove("active");
                menuToggle.classList.remove("active");
                socialLinks.classList.remove("active");
            }
        }
        document.body.addEventListener("click", closeMenu);
        document.body.addEventListener("touchstart", closeMenu);
    })();
}

function Accordion(element) {
    let toggle = element.getElementsByClassName("toggle")[0],
        content = element.getElementsByClassName("content")[0];
    toggle.addEventListener("click", function () {
        let wasActive = this.classList.contains("active");
        let query = ".accordion > .active";
        while (document.querySelector(query)) {
            document.querySelector(query).classList.remove("active");
        }
        if (!wasActive) {
            let prevState = content.classList.contains("active");
            content.classList.toggle("active");
            toggle.classList[prevState ? "remove" : "add"]("active");
        }
    });
}

Accordion.Group = function (elements) {
    for (let element of elements) {
        Accordion(element);
    }
}

export { Accordion, TopNav };