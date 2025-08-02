import { API } from "../services/API.js";

import { MovieItemComponent } from "./MovieItem.js";

export class HomePage extends HTMLElement {
  async render() {
    const topSection = document.querySelector("#top-10");
    const randomSection = document.querySelector("#random");
    topSection.classList.add("loading");
    randomSection.classList.add("loading");

    try {
      const topMovies = await API.getTopMovies();
      this.renderMoviesInList(topMovies, topSection.querySelector("ul"));
    } catch (error) {
      this.showErrorInSection(topSection, error);
    } finally {
      topSection.classList.remove("loading");
    }

    try {
      const randomMovies = await API.getRandomMovies();
      this.renderMoviesInList(randomMovies, randomSection.querySelector("ul"));
    } catch (error) {
      this.showErrorInSection(randomSection, error);
    } finally {
      randomSection.classList.remove("loading");
    }
  }

  renderMoviesInList(movies, ul) {
    ul.innerHTML = "";
    if (!Array.isArray(movies)) {
      throw new Error(
        movies.error || "An error occurred while fetching movies.",
      );
    }
    movies.forEach((movie, index) => {
      const li = document.createElement("li");
      li.style.animationDelay = `${index * 0.1}s`;
      li.appendChild(new MovieItemComponent(movie));
      ul.appendChild(li);
    });
  }

  showErrorInSection(section, error) {
    const ul = section.querySelector("ul");
    ul.innerHTML = `<li class="error-message">${error.message || "Failed to load movies."} <button onclick="location.reload()">Retry</button></li>`;
  }

  connectedCallback() {
    const template = document.getElementById("template-home");
    const content = template.content.cloneNode(true);
    this.appendChild(content);
    this.render();
  }
}

customElements.define("home-page", HomePage);
