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
  api: API,
};
