/**
 * Variables
 */
const signupButton = document.getElementById('signup-button'),
    loginButton = document.getElementById('login-button'),
    userForms = document.getElementById('user_options-forms')

/**
 * Add event listener to the "Sign Up" button
 */
signupButton.addEventListener('click', () => {
  userForms.classList.remove('bounceRight')
  userForms.classList.add('bounceLeft')
}, false)

/**
 * Add event listener to the "Login" button
 */
loginButton.addEventListener('click', () => {
  userForms.classList.remove('bounceLeft')
  userForms.classList.add('bounceRight')
}, false)

function login() {
  var arr = {
    "emailLogin":"emailLogin",
    "passwordLogin":"passwordLogin"
  }
  $.ajax({
    type: "POST",
    dataType: "json",
    url: "http://127.0.0.1:9300/index" ,
    data:{
      emailLogin: $("[name=emailLogin]").val(),
      passwordLogin: $("[name=passwordLogin]").val()
    },
    success: function(result){
      var obj =JSON.parse(result);
      if(obj.code === 200){
        alert("success")
      }else{
        alert("fail");
      }
    }
  });
}