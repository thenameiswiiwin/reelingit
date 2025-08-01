import { API } from "../services/API.js";
import { MovieItemComponent } from "./MovieItem.js";

export class HomePage extends HTMLElement {
  async render() {
    const topMovies = await API.getTopMovies();
    renderMoviesInList(topMovies, document.querySelector("#top-10 ul"));

    const randomMovies = await API.getRandomMovies();
    renderMoviesInList(randomMovies, document.querySelector("#random ul"));

    function renderMoviesInList(movies, ul) {
      ul.innerHTML = "";

      if (!Array.isArray(movies)) {
        ul.innerHTML = `
          <li class="error">
            ${movies.error || "An error occurred while fetching movies."}
          </li>
        `;
        console.error("Error fetching movies:", movies.error);
        return;
      }

      movies.forEach((movie) => {
        const li = document.createElement("li");
        li.appendChild(new MovieItemComponent(movie));
        ul.appendChild(li);
      });
    }
  }

  connectedCallback() {
    const template = document.getElementById("template-home");
    const content = template.content.cloneNode(true);

    this.appendChild(content);
    this.render();
  }
}

customElements.define("home-page", HomePage);
