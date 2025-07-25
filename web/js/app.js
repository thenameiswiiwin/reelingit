import { API } from "./services/API.js";

window.app = {
  search: (event) => {
    event.preventDefault();
    const keywords = document.querySelector("input[type=search]").value;
    // TODO: Implement search functionality
  },
  api: API,
};
