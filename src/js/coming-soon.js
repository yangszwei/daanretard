function convertTime(seconds) {
    let pad = (num) => num.toString().padStart(2, 0);
    let d = Math.floor(seconds / 60 / 60 / 24),
        h = pad(Math.floor(seconds / 60 / 60) % 24),
        m = pad(Math.floor(seconds / 60) % 60),
        s = pad(Math.floor(seconds % 60));
    return `${d}天 ${h}:${m}:${s}`;
}

window.onload = function() {
    const timer = document.getElementsByClassName("countdown")[0],
        dest = 1596113094000;
    function countdown() {
        let remaining_time = (dest - Date.now()) / 1000;
        timer.innerText = convertTime(Math.max(remaining_time, 0));
    }
    countdown();
    setInterval(countdown, 1000);
    const kotori = document.getElementsByClassName("kotori")[0];
    let kotoriSize = 6;
    setInterval(() => {
        kotoriSize = Math.min(kotoriSize + 0.01, 35);
        kotori.style.maxHeight = `${kotoriSize}rem`;
    }, 100);
    kotori.onclick = (event) => {
        event.preventDefault();
        event.stopPropagation();
        kotoriSize = Math.max(6, kotoriSize - 0.1);
    };
}