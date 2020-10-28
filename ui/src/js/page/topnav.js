/* eslint-env browser */
const topnav = document.getElementById('topnav')
const toggle = document.getElementById('topnav-toggle')
const menus = topnav.getElementsByTagName('ul')

let visible = false

toggle.onclick = function () {
  const iconMenu = toggle.querySelector("[data-icon='mdi:menu']")
  const iconClose = toggle.querySelector("[data-icon='mdi:close']")
  visible = !visible
  if (visible) {
    for (const menu of menus) {
      menu.classList.remove('hidden')
    }
    iconMenu.classList.add('hidden')
    iconClose.classList.remove('hidden')
  } else {
    for (const menu of menus) {
      menu.classList.add('hidden')
    }
    iconMenu.classList.remove('hidden')
    iconClose.classList.add('hidden')
  }
}
