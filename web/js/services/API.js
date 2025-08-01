export const API = {
  baseURL: "/api",
  getTopMovies: async () => await API.fetch("movies/top"),
  getRandomMovies: async () => await API.fetch("movies/random"),
  getMovieById: async (id) => await API.fetch(`movies/${id}`),
  searchMovies: async (q, order, genre) =>
    await API.fetch("movies/search", { q, order, genre }),

  fetch: async (endpoint, params) => {
    try {
      const queryString = new URLSearchParams(params || {}).toString();
      const url = queryString
        ? `${API.baseURL}/${endpoint}?${queryString}`
        : `${API.baseURL}/${endpoint}`;

      const response = await fetch(url, {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
        },
      });

      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }

      return await response.json();
    } catch (error) {
      console.error("API fetch error:", error);
      return {
        error: error.message || "An error occurred while fetching data.",
      };
    }
  },
};
