import React from 'react'
import ReactDOM from 'react-dom/client'
import './index.css'
import { RouterProvider, createBrowserRouter } from 'react-router-dom'
import Root from './layout/Root.tsx'
import ErrorPage from './pages/ErrorPage.tsx'
import StockPage from './pages/StockPage.tsx'
import { initialLoader } from "./data/loaders.tsx"

const router = createBrowserRouter([
  {
    path: "/", 
    element: <Root></Root>,
    errorElement: <ErrorPage></ErrorPage>,
    loader: initialLoader,
    children: [
      {
        path: "/stock",
        element: <StockPage></StockPage>
      }
    ]
  }
])

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>,
)
