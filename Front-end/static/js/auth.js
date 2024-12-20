const signupBtn = document.getElementById("signup-btn");
const signinBtn = document.getElementById("signin-btn");
const mainContainer = document.querySelector(".container");

signupBtn.addEventListener("click", () => {
  mainContainer.classList.toggle("change");
});
signinBtn.addEventListener("click", () => {
  mainContainer.classList.toggle("change");
});

function isOnlyNumbers(str) {
  for (let i = 0; i < str.length; i++) {
    if (isNaN(str[i]) || str[i] === " ") {
      return false; // Contains a non-numeric character
    }
  }
  return true; // All characters are numeric
}

document.addEventListener("DOMContentLoaded", function () {
  // Signup form validation
  const signupForm = document.querySelector(".signup-form form");
  const usernameRegister = document.getElementById("username_register");
  const emailRegister = document.getElementById("email_register");
  const passwordRegister = document.getElementById("password_register");
  const confirmPasswordRegister = document.getElementById("confirm_password_register");

  signupForm.addEventListener("submit", function (e) {
    e.preventDefault();

    // Check username
    if (usernameRegister.value.trim() === ""||isOnlyNumbers(usernameValue)) {
      alert("Username is required for signup.");
      return;
    }

    // Check email format
    if (!/\S+@\S+\.\S+/.test(emailRegister.value)) {
      alert("Please enter a valid email address.");
      return;
    }

    // Check password length
    if (passwordRegister.value.length < 8) {
      alert("Password must be at least 8 characters long.");
      return;
    }

    // Check if passwords match
    if (passwordRegister.value !== confirmPasswordRegister.value) {
      alert("Passwords do not match. Please confirm your password.");
      return;
    }

    alert("Signup successful!");
    signupForm.submit(); // Submit the form if everything is valid
  });
});
//   // Signin form validation
//   const signinForm = document.querySelector(".signin-form form");
//   const usernameLogin = document.getElementById("username_login");
//   const passwordLogin = document.getElementById("password_login");

//   signinForm.addEventListener("submit", function (e) {
//     e.preventDefault();

//     // Check username
//     if (usernameLogin.value.trim() === "") {
//       alert("Username is required for signin.");
//       return;
//     }

//     // Check password length
//     if (passwordLogin.value.length < 8) {
//       alert("Password must be at least 8 characters long.");
//       return;
//     }

//     alert("Signin successful!");
//     signinForm.submit(); // Submit the form if everything is valid
//   });
// });


document.querySelector(".signin-form form").addEventListener("submit", async function (e) {
  e.preventDefault();

  const username = document.getElementById("username_login").value.trim();
  const password = document.getElementById("password_login").value.trim();

  if (username === "" || password === "") {
    alert("Both username and password are required.");
    return;
  }

  const response = await fetch("/login", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ username, password }),
  });
console.log(response);

  if (response.ok) {
    window.location.href = "/home";
  } else {
    alert("Invalid username or password.");
  }
});





