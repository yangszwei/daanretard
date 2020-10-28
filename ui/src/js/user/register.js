/* eslint-env browser */
import '../page/topnav'

const form = document.getElementById('register')
const submit = document.getElementById('submit')

function validate () {
  let isValid = true
  const firstName = document.getElementById('first-name')
  const lastName = document.getElementById('last-name')
  const displayName = document.getElementById('display-name')
  const email = document.getElementById('email')
  const password = document.getElementById('password')
  const retypePassword = document.getElementById('retype-password')
  const isFirstNameValid = firstName.value.length >= 1 && firstName.value.length <= 50
  const isLastNameValid = lastName.value.length >= 1 && lastName.value.length <= 50
  if (isFirstNameValid && isLastNameValid) {
    document.getElementById('error-name').innerText = ''
  } else {
    document.getElementById('error-name').innerText = '請輸入姓名 (長度需介於1-50個字元)'
    isValid = false
  }
  if (displayName.value.length >= 1 && displayName.value.length <= 50) {
    document.getElementById('error-display-name').innerText = ''
  } else {
    document.getElementById('error-display-name').innerText = '請輸入暱稱 (長度需介於1-50個字元)'
    isValid = false
  }
  if (email.validity.valid) {
    document.getElementById('error-email').innerText = ''
  } else {
    document.getElementById('error-email').innerText = '請輸入有效的電子郵件'
    isValid = false
  }
  if (password.validity.valid) {
    document.getElementById('error-password').innerText = ''
  } else {
    document.getElementById('error-password').innerText = '密碼長度需介於8-30個字元'
    isValid = false
  }
  if (password.value === retypePassword.value) {
    document.getElementById('error-retype-password').innerText = ''
  } else {
    document.getElementById('error-retype-password').innerText = '這些密碼不相符，請重新輸入'
    isValid = false
  }
  return isValid
}

submit.onclick = async function (e) {
  e.preventDefault()
  if (!validate()) { return }
  let request = await fetch('/api/user', {
    method: 'POST',
    body: new FormData(form)
  })
  if (request.ok) {
    location.pathname = '/'
  } else if (request.status === 400) {
    const response = await request.json()
    if (response.message === 'Invalid credentials') {
      alert("註冊失敗! 表單內容未通過驗證")
    } else if (response.message === 'Email already taken') {
      document.getElementById('error-email').innerText = '此電子郵件已被註冊'
    }
  }
}