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
  register: async (event) => {
    event.preventDefault();
    const name = document.getElementById("register-name").value.trim();
    const email = document.getElementById("register-email").value.trim();
    const password = document.getElementById("register-password").value.trim();
    const passwordConfirm = document
      .getElementById("register-password-confirm")
      .value.trim();

    const errors = [];
    if (name.length < 4)
      errors.push("Name must be at least 4 characters long.");
    if (password.length < 6)
      errors.push("Password must be at least 6 characters long.");
    if (email.length < 5 || !email.includes("@"))
      errors.push("Please enter a valid email address.");
    if (password !== passwordConfirm) errors.push("Passwords do not match.");

    if (errors.length === 0) {
      try {
        const response = await API.register(name, email, password);
        if (response.success) {
          app.Router.go("/account/");
        } else {
          app.showError(
            "Registration Error",
            response.message || "An error occurred during registration.",
          );
        }
      } catch (error) {
        app.showError("Registration Error", error.message);
      }
    } else {
      app.showError("Registration Error", errors.join("\n"));
    }
  },
  login: async (event) => {
    event.preventDefault();
    const email = document.getElementById("login-email").value.trim();
    const password = document.getElementById("login-password").value.trim();
    const errors = [];
    if (email.length < 5 || !email.includes("@"))
      errors.push("Please enter a valid email address.");
    if (password.length < 6)
      errors.push("Password must be at least 6 characters long.");
    if (errors.length === 0) {
      try {
        const response = await API.login(email, password);
        if (response.success) {
          app.Router.go("/account/");
        } else {
          app.showError(
            "Login Error",
            response.message || "An error occurred during login.",
          );
        }
      } catch (error) {
        app.showError("Login Error", error.message);
      }
    } else {
      app.showError("Login Error", errors.join("\n"));
    }
  },
  api: API,
};
