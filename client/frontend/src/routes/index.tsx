import AuthLayout from "../layouts/auth-layout.tsx";
import { createBrowserRouter } from "react-router-dom";
import NavbarLayout from "../layouts/navbar-layout.tsx";
import Home from "./home.tsx";
import Login from "./login.tsx";
import Register from "./register.tsx";
import Dashboard from "./dashboard.tsx";

const router = createBrowserRouter([
  {
    path: "/",
    element: <NavbarLayout />,
    children: [
      {
        index: true,
        element: <Home />,
      },
    ],
  },
  {
    path: "/app",
    children: [
      {
        path: "dashboard",
        element: <Dashboard />,
      }
    ]

  },

  {
    element: <AuthLayout />,
    children: [
      {
        path: "/register",
        element: <Register />,
      },
      {
        path: "/login",
        element: <Login />,
      },
    ],
  },
]);

export default router;
