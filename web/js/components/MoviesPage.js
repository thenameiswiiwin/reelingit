import { API } from "../services/API.js";
import { MovieItemComponent } from "./MovieItem.js";

export class MoviesPage extends HTMLElement {
  async loadGenres() {
    try {
      const genres = await API.getGenres();
      const select = this.querySelector("#filter");
      select.innerHTML = `<option>Filter by Genre</option>`;
      genres.forEach((genre) => {
        const option = document.createElement("option");
        option.value = genre.id;
        option.textContent = genre.name;
        select.appendChild(option);
      });
    } catch (error) {
      console.error("Failed to load genres:", error);
    }
  }

  async render(query) {
    const section = this.querySelector("section");
    section.classList.add("loading");

    const urlParams = new URLSearchParams(window.location.search);
    const order = urlParams.get("order") ?? "";
    const genre = urlParams.get("genre") ?? "";

    try {
      const movies = await API.searchMovies(query, order, genre);
      const ulMovies = this.querySelector("ul");
      ulMovies.innerHTML = "";
      if (movies && movies.length > 0) {
        movies.forEach((movie, index) => {
          const li = document.createElement("li");
          li.style.animationDelay = `${index * 0.1}s`;
          li.appendChild(new MovieItemComponent(movie));
          ulMovies.appendChild(li);
        });
      } else {
        ulMovies.innerHTML = "<h3 class='error-message'>No movies found</h3>";
      }
    } catch (error) {
      this.querySelector("ul").innerHTML =
        `<h3 class="error-message">${error.message || "Failed to search movies."} <button onclick="location.reload()">Retry</button></h3>`;
    } finally {
      section.classList.remove("loading");
    }

    if (order) this.querySelector("#order").value = order;
    if (genre) this.querySelector("#filter").value = genre;
  }

  connectedCallback() {
    const template = document.getElementById("template-movies");
    const content = template.content.cloneNode(true);
    this.appendChild(content);

    const urlParams = new URLSearchParams(window.location.search);
    const query = urlParams.get("q");
    if (query) {
      this.querySelector("h2").textContent = `'${query}' movies`;
      this.render(query);
      this.loadGenres();
    } else {
      app.showError("Search Error", "No search query provided.");
      app.Router.go("/");
    }
  }
}

customElements.define("movies-page", MoviesPage);
