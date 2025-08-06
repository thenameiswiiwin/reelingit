import "./components/AnimatedLoading.js";
import "./components/YoutubeEmbed.js";
import { API } from "./services/API.js";
import { Router } from "./services/Router.js";

window.addEventListener("DOMContentLoaded", (event) => {
  document.getElementById("global-loader").style.display = "none";
  app.Router.init();
});

window.app = {
  Router,
  showError: (
    title = "Error",
    message = "Oops! Something went wrong on our end. Please try again in a moment or refresh the page.",
    goToHome = false,
  ) => {
    const dialog = document.getElementById("alert-modal");
    dialog.querySelector(".modal-title").textContent = title;
    dialog.querySelector(".modal-message").textContent = message;
    dialog.showModal();
    if (goToHome) app.Router.go("/");
  },
  closeError: () => {
    document.getElementById("alert-modal").close();
  },
  showToast: (message, type = "success") => {
    const toast = document.createElement("div");
    toast.className = `toast ${type}`;
    toast.textContent = message;
    document.body.appendChild(toast);
    setTimeout(() => {
      toast.classList.add("show");
    }, 100);
    setTimeout(() => {
      toast.classList.remove("show");
      setTimeout(() => toast.remove(), 300);
    }, 3000);
  },
  search: (event) => {
    event.preventDefault();
    const q = document.querySelector("input[type=search]").value.trim();
    if (!q) {
      app.showError("Search Error", "Please enter a search query.");
      return;
    }
    app.Router.go("/movies?q=" + encodeURIComponent(q));
  },
  searchFilterChange: (value) => {
    const urlParams = new URLSearchParams(window.location.search);
    if (value && value !== "Filter by Genre") urlParams.set("genre", value);
    else urlParams.delete("genre");
    app.Router.go(`${window.location.pathname}?${urlParams.toString()}`);
  },
  searchOrderChange: (value) => {
    const urlParams = new URLSearchParams(window.location.search);
    urlParams.set("order", value);
    app.Router.go(`${window.location.pathname}?${urlParams.toString()}`);
  },
  validateField: (field, errorSpan, validationFn, errorMessage) => {
    const isValid = validationFn(field.value.trim());
    field.setAttribute("aria-invalid", !isValid);
    errorSpan.textContent = isValid ? "" : errorMessage;
    errorSpan.style.display = isValid ? "none" : "block";
    return isValid;
  },
  setupRealTimeValidation: (formId) => {
    const emailRegex = /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/;
    const forms = document.querySelectorAll(`#${formId}`);
    forms.forEach((form) => {
      const nameInput = form.querySelector("#register-name");
      const emailInput =
        form.querySelector("#register-email") ||
        form.querySelector("#login-email");
      const passwordInput =
        form.querySelector("#register-password") ||
        form.querySelector("#login-password");
      const confirmPasswordInput = form.querySelector(
        "#register-password-confirm",
      );

      const nameError = form.querySelector("#name-error");
      const emailError = form.querySelector("#email-error");
      const passwordError = form.querySelector("#password-error");
      const confirmError = form.querySelector("#confirm-error");

      if (nameInput) {
        nameInput.addEventListener("input", () =>
          app.validateField(
            nameInput,
            nameError,
            (value) => value.length >= 4,
            "Name must be at least 4 characters long.",
          ),
        );
      }
      if (emailInput) {
        emailInput.addEventListener("input", () =>
          app.validateField(
            emailInput,
            emailError,
            (value) => emailRegex.test(value),
            "Please enter a valid email address.",
          ),
        );
      }
      if (passwordInput) {
        passwordInput.addEventListener("input", () => {
          app.validateField(
            passwordInput,
            passwordError,
            (value) => value.length >= 6,
            "Password must be at least 6 characters long.",
          );
          if (confirmPasswordInput)
            app.validateField(
              confirmPasswordInput,
              confirmError,
              (value) => value === passwordInput.value.trim(),
              "Passwords do not match.",
            );
        });
      }
      if (confirmPasswordInput) {
        confirmPasswordInput.addEventListener("input", () =>
          app.validateField(
            confirmPasswordInput,
            confirmError,
            (value) => value === passwordInput.value.trim(),
            "Passwords do not match.",
          ),
        );
      }
    });
  },
  register: async (event) => {
    event.preventDefault();
    const name = document.getElementById("register-name").value.trim();
    const email = document.getElementById("register-email").value.trim();
    const password = document.getElementById("register-password").value.trim();
    const passwordConfirm = document
      .getElementById("register-password-confirm")
      .value.trim();
    const errors = [];
    const emailRegex = /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/;
    if (name.length < 4)
      errors.push("Name must be at least 4 characters long.");
    if (!emailRegex.test(email))
      errors.push("Please enter a valid email address.");
    if (password.length < 6)
      errors.push("Password must be at least 6 characters long.");
    if (password !== passwordConfirm) errors.push("Passwords do not match.");
    if (errors.length > 0) {
      app.showError("Registration Error", errors.join("\n"));
      return;
    }
    try {
      const response = await API.register(name, email, password);
      if (response.success) {
        app.showToast("Registration successful!");
        setTimeout(() => app.Router.go("/account/"), 3000);
      } else {
        app.showError(
          "Registration Error",
          response.message || "An error occurred during registration.",
        );
      }
    } catch (error) {
      app.showError("Registration Error", error.message);
    }
  },
  login: async (event) => {
    event.preventDefault();
    const email = document.getElementById("login-email").value.trim();
    const password = document.getElementById("login-password").value.trim();
    const errors = [];
    const emailRegex = /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/;
    if (!emailRegex.test(email))
      errors.push("Please enter a valid email address.");
    if (password.length < 6)
      errors.push("Password must be at least 6 characters long.");
    if (errors.length > 0) {
      app.showError("Login Error", errors.join("\n"));
      return;
    }
    try {
      const response = await API.login(email, password);
      if (response.success) {
        app.showToast("Login successful!");
        setTimeout(() => app.Router.go("/account/"), 3000);
      } else {
        app.showError(
          "Login Error",
          response.message || "An error occurred during login.",
        );
      }
    } catch (error) {
      app.showError("Login Error", error.message);
    }
  },
  api: API,
};

document.addEventListener("DOMContentLoaded", () => {
  app.setupRealTimeValidation("register-form");
  app.setupRealTimeValidation("login-form");
});
