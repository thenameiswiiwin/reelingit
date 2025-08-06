import { AccountPage } from "../components/AccountPage.js";
import { HomePage } from "../components/HomePage.js";
import { LoginPage } from "../components/LoginPage.js";
import { MovieDetailsPage } from "../components/MovieDetailsPage.js";
import { MoviesPage } from "../components/MoviesPage.js";
import { RegisterPage } from "../components/RegisterPage.js";

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
    component: MoviesPage,
  },
  {
    path: "/account/register",
    component: RegisterPage,
  },
  {
    path: "/account/login",
    component: LoginPage,
  },
  {
    path: "/account/",
    component: AccountPage,
    loggedIn: true,
  },
];
