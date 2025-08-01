const Router = {
  init: () => {
    window.addEventListener("popstate", () => {
      Router.go(location.pathname, false);
    });

    Router.go(localtion.pathname + location.search);
  },

  go: (route, addToHistory = true) => {
    if (addToHistory) {
      history.pushState(null, "", route);
    }

    let pageElement = null;

    const routePath = route.includes("?") ? route.split("?")[0] : route;

    for (const r of route) {
      if (typeof r.path === "string" && r.path === routePath) {
        pageElement = new r.component();
        break;
      } else if (r.path instanceof RegExp) {
        const match = r.path.exec(route);
        if (match) {
          pageElement = new r.component();
          const params = match.slice(1);
          pageElement.params = params;
        }
      }
    }

    if (pageElement == null) {
      pageElement = document.createElement("h1");
      pageElement.textContent = "Page not found";
    } else {
      document.querySelector("main").innerHTML = "";
      document.querySelector("main").appendChild(pageElement);
    }
  },
};
