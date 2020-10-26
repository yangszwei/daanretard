/* eslint-env browser */
import '../component/topnav'

const form = document.getElementById('register-form')
const name = document.getElementById('name')
const firstName = document.getElementById('first-name')
const lastName = document.getElementById('last-name')
const email = document.getElementById('email')
const password = document.getElementById('password')
const retypePassword = document.getElementById('retype-password')
const submit = document.getElementById('submit')

submit.onclick = function (e) {
  e.preventDefault()
  const formData = new FormData(form)
  formData.delete('retype-password')
  const request = new XMLHttpRequest()
  request.onload = function () {
    if (request.readyState === 4) {
      if (request.status === 200) {
        location.href = '/'
      }
    }
  }
  request.open('POST', '/api/user')
  request.send(formData)
}