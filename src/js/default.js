window.addEventListener("load", () => {

    const navbar = document.getElementById("navbar");

    (function initMenuToggle() {
        let toggle = document.querySelector("#navbar > .toggle-menu");
        toggle.onclick = (event) => {
            event.stopPropagation();
            navbar.classList.toggle("active");
        };
    })();

    (function closeMenuOnClickOutside() {
        document.body.addEventListener("click", (event) => {
          if (!(this === navbar) || !navbar.contains(this)) {
              navbar.classList.remove("active");
          }
        });
    })();

});