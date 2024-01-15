import React from "react";
import ReactDOM from "react-dom/client";
import "./index.css";
import { RouterProvider, createBrowserRouter } from "react-router-dom";
import Layout from "./layout/Layout.tsx";
import ErrorPage from "./pages/ErrorPage.tsx";
import PortfolioPage from "./pages/PortfolioPage.tsx";
import PortfoliosPage from "./pages/PortfoliosPage.tsx";
import BudgetPage from "./pages/BudgetPage.tsx";
import { layoutPageLoader, portfolioPageLoader, portfoliosPageLoader } from "./data/loaders.ts";

const router = createBrowserRouter([
  {
    path: "/",
    element: <Layout></Layout>,
    errorElement: <ErrorPage></ErrorPage>,
    loader: layoutPageLoader,
    children: [
      {
        path: "/budget",
        element: <BudgetPage></BudgetPage>
      },
      {
        path: "/portfolios",
        loader: portfoliosPageLoader,
        element: <PortfoliosPage></PortfoliosPage>
      },
      {
        path: "/portfolios/:portfolioId",
        loader: portfolioPageLoader,
        element: <PortfolioPage></PortfolioPage>,
      },
    ],
  },
]);

ReactDOM.createRoot(document.getElementById("root")!).render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>
);
