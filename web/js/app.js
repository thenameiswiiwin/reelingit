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
    const q = document.querySelector("input[type=search]").value;
    app.Router.go("/movies?q=" + q);
  },
  api: API,
};
