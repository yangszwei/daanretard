function initTabButton() {
    let buttons = document.querySelectorAll(".subnav > [data-target]");
    for (let button of buttons) {
        button.addEventListener("click", function () {
            let prevButton = document.querySelector(".subnav > [data-target].active");
            if (prevButton) prevButton.classList.remove("active");
            let prevTab = document.querySelector(".subnav + .content > section.active");
            if (prevTab) prevTab.classList.remove("active");
            document.getElementById(this.dataset.target).classList.add("active");
            this.classList.add("active");
        });
    }
}

onload = () => {
    initTabButton();
};