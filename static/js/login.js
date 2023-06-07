const showPasswordButton = document.getElementById('show');
const hidePasswordButton = document.getElementById('hide');

showPasswordButton.addEventListener(
  "click",
  () => {
    showPasswordButton.classList.add("password-input_show");
    hidePasswordButton.classList.add("password-input_hide");
    document.querySelector(".form__field-password").type = "text";
  }
)

hidePasswordButton.addEventListener(
    "click",
    () => {
      showPasswordButton.classList.remove("password-input_showe");
      hidePasswordButton.classList.remove("password-input_hide");
      document.querySelector(".form__field-password").type = "password";
    }
  )