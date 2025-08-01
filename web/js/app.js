import "./components/AnimatedLoading.js";
import "./components/YoutubeEmbed.js";
import { API } from "./services/API.js";
import { Router } from "./services/Router.js";

window.addEventListener("DOMContentLoaded", (event) => {
  app.Router.init();
});

window.app = {
  Router,
  showError: (
    message = "An unexpected error occurred. Please try again later.",
    goToHome = false,
  ) => {
    const dialog = document.getElementById("alert-modal");
    dialog.showModal();
    dialog.querySelector("p").textContent = message;
    if (goToHome) app.Router.go("/");
  },
  closeError: () => {
    document.getElementById("alert-modal").close();
  },
  search: (event) => {
    event.preventDefault();
    const q = document.querySelector("input[type=search]").value;
    // TODO: Implement search functionality
  },
  api: API,
};
