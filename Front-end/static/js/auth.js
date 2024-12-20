const signupBtn = document.getElementById("signup-btn");
const signinBtn = document.getElementById("signin-btn");
const mainContainer = document.querySelector(".container");

signupBtn.addEventListener("click", () => {
  mainContainer.classList.toggle("change");
});
signinBtn.addEventListener("click", () => {
  mainContainer.classList.toggle("change");
});

// Handle login form submission
document.querySelector(".signin-form form").addEventListener("submit", async function (e) {
  e.preventDefault();

  const username = document.getElementById("username_login").value.trim();
  const password = document.getElementById("password_login").value.trim();
  const errorContainer = document.getElementById("login-error");

  // Clear previous errors
  errorContainer.textContent = "";
  errorContainer.style.display = "none";
  console.log("Error container:", errorContainer);
  // Input validation
  if (username === "" || password === "") {
    errorContainer.textContent = "Both username and password are required.";
    errorContainer.style.display = "block";
    return;
  }

  try {
    const response = await fetch("/login", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ username, password }),
    });

    if (response.ok) {
      window.location.href = "/home";
    } else {
      const errorData = await response.json();
      errorContainer.textContent = errorData.error || "Invalid username or password.";
      errorContainer.style.display = "block";
      console.log("Error container:", errorContainer);
    }
  } catch (error) {
    console.error("Error during login:", error);
    errorContainer.textContent = "An error occurred. Please try again later.";
    errorContainer.style.display = "block";
  }
});

document.querySelector(".signup-form form").addEventListener("submit", async function (e) {
  e.preventDefault();

  // Get input values
  const username = document.getElementById("username_register").value.trim();
  const email = document.getElementById("email_register").value.trim();
  const password = document.getElementById("password_register").value.trim();
  const confirmPassword = document.getElementById("confirm_password_register").value.trim();
  const errorContainer = document.getElementById("register-error");

  // Clear previous errors
  errorContainer.textContent = "";
  errorContainer.style.display = "none";

  // Input validation
  if (!username || !email || !password || !confirmPassword) {
    errorContainer.textContent = "All fields are required.";
    errorContainer.style.display = "block";
    return;
  }

  if (password !== confirmPassword) {
    errorContainer.textContent = "Passwords do not match.";
    errorContainer.style.display = "block";
    return;
  }

  if (password.length < 6) {
    errorContainer.textContent = "Password must be at least 6 characters long.";
    errorContainer.style.display = "block";
    return;
  }

  // Send data to the server
  try {
    const response = await fetch("/register", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        username,
        email,
        password,
        confirmPassword,
      }),
    });

    if (response.ok) {
      window.location.href = "/home";
    } else {
      const errorText = await response.text();
      errorContainer.textContent = errorText || "Registration failed. Please try again.";
      errorContainer.style.display = "block";
    }
  } catch (error) {
    console.error("Error during registration:", error);
    errorContainer.textContent = "An error occurred. Please try again later.";
    errorContainer.style.display = "block";
  }
});