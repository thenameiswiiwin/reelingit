export const API = {
  baseURL: "/api",
  getTopMovies: async () => await API.fetch("movies/top"),
  getRandomMovies: async () => await API.fetch("movies/random"),
  getGenres: async () => await API.fetch("genres"),
  getMovieById: async (id) => await API.fetch(`movies/${id}`),
  searchMovies: async (q, order, genre) =>
    await API.fetch("movies/search", { q, order, genre }),
  register: async (name, email, password) =>
    await API.send("account/register", { name, email, password }),
  login: async (email, password) =>
    await API.send("account/authenticate", { email, password }),
  send: async (serviceName, data) => {
    try {
      const url = `${API.baseURL}/${serviceName}`;
      const response = await fetch(url, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
      });
      if (!response.ok) {
        const errorData = await response.json().catch(() => ({}));
        throw new Error(
          errorData.message || `HTTP error! status: ${response.status}`,
        );
      }
      return await response.json();
    } catch (error) {
      console.error("API send error:", error);
      throw error;
    }
  },
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
        const errorData = await response.json().catch(() => ({}));
        throw new Error(
          errorData.message || `HTTP error! status: ${response.status}`,
        );
      }
      return await response.json();
    } catch (error) {
      console.error("API fetch error:", error);
      throw error;
    }
  },
};
