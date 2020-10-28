/* eslint-env browser */
import '../page/topnav'

const form = document.getElementById('login-form')
const submit = document.getElementById('submit')

submit.onclick = async function (e) {
  e.preventDefault()
  const formData = new FormData(form)
  const request = new XMLHttpRequest()
  request.onload = function () {
    if (request.readyState === 4) {
      if (request.status === 200) {
        location.href = '/'
      }
    }
  }
  request.onerror = function () {
    console.log(request.responseText)
  }
  request.open('POST', '/api/user/session')
  request.send(formData)
}