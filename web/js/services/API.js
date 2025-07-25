export const API = {
  baseURL: "/api",
  getTopMovies: () => API.fetch("movies/top"),
  getRandomMovies: () => API.fetch("movies/random"),
  getMovieByID: (id) => API.fetch(`movies/${id}`),
  searchMoviesByName: (keywords, order, genre) =>
    API.fetch(`movies/search`, { q: keywords, order, genre }),

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
