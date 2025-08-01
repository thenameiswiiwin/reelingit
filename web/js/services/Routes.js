import { HomePage } from "../components/HomePage";
import { MovieDetailsPage } from "../components/MovieDetailsPage";
import { MoviePage } from "../components/MoviePage";

export const routes = [
  {
    path: "/",
    component: HomePage,
  },
  {
    path: /\/movies\/(\d+)/,
    component: MovieDetailsPage,
  },
  {
    path: "/movies",
    component: MoviePage,
  },
];
