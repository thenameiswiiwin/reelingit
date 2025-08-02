import { API } from "../services/API.js";

export class MovieDetailsPage extends HTMLElement {
  id = null;
  movie = null;

  async render() {
    const article = this.querySelector("#movie");
    if (article) article.classList.add("loading");

    try {
      this.movie = await API.getMovieById(this.id);
    } catch (error) {
      app.showError(
        "Movie Load Error",
        error.message || "Failed to load movie details.",
      );
      app.Router.go("/");
      return;
    } finally {
      if (article) article.classList.remove("loading");
    }

    const template = document.getElementById("template-movie-details");
    const content = template.content.cloneNode(true);
    this.appendChild(content);

    this.querySelector(".title-group").innerHTML = `
      <h2>${this.movie.title}</h2>
      <h3>${this.movie.tagline}</h3>
      <dl>
        <dt>Release Year</dt>
        <dd>${this.movie.release_year}</dd>
      </dl>
    `;

    const img = this.querySelector("img");
    img.src = this.movie.poster_url;
    img.alt = `${this.movie.title} Poster`;

    this.querySelector("#overview").textContent = this.movie.overview;

    const trailer = this.querySelector("#trailer");
    trailer.dataset.url = this.movie.trailer_url;
    trailer.setAttribute("aria-label", `${this.movie.title} Trailer`);

    this.querySelector("#metadata").innerHTML = `
      <dt>Score</dt>
      <dd>${this.movie.score} / 10</dd>
      <dt>Popularity</dt>
      <dd>${this.movie.popularity}</dd>
    `;

    const ulGenres = this.querySelector("#genres");
    ulGenres.innerHTML = "";
    this.movie.genres.forEach((genre) => {
      const li = document.createElement("li");
      li.textContent = genre.name;
      ulGenres.appendChild(li);
    });

    const ulCast = this.querySelector("#cast");
    ulCast.innerHTML = "";
    this.movie.casting.forEach((actor, index) => {
      const li = document.createElement("li");
      li.style.animationDelay = `${index * 0.1}s`;
      li.innerHTML = `
        <img src="${actor.image_url ?? "/images/generic_actor.jpg"}" alt="Picture of ${actor.first_name} ${actor.last_name}" loading="lazy">
        <p>${actor.first_name} ${actor.last_name}</p>
      `;
      ulCast.appendChild(li);
    });
  }

  connectedCallback() {
    this.id = this.params[0];
    if (!this.id) {
      app.showError("Invalid Movie ID", "Please select a valid movie.");
      app.Router.go("/");
      return;
    }
    this.render();
  }
}

customElements.define("movie-details-page", MovieDetailsPage);
