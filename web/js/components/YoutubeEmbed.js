export class YoutubeEmbed extends HTMLElement {
  static get observedAttributes() {
    return ["data-url"];
  }

  attributeChangedCallback(prop) {
    if (prop === "data-url") {
      const url = this.dataset.url;

      if (!url) return;

      const videoId = url.split("v=")[1]?.split("&")[0];

      this.innerHTML = `
        <div class="youtube-wrapper">
          <iframe
            src="https://www.youtube.com/embed/${videoId}"
            title="YouTube video player"
            frameborder="0"
            allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share"
            referrerpolicy="strict-origin-when-cross-origin"
            allowfullscreen
          ></iframe>
        </div>
      `;
    }
  }
}

customElements.define("youtube-embed", YoutubeEmbed);
