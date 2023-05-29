import { createBrowserRouter } from "react-router-dom";
import AppLayout from "../layouts/app-layout.tsx";
import AuthLayout from "../layouts/auth-layout.tsx";
import NavbarLayout from "../layouts/navbar-layout.tsx";
import Dashboard from "./dashboard.tsx";
import Findings from "./findings.tsx";
import Home from "./home.tsx";
import Login from "./login.tsx";
import Plugins from "./plugins.tsx";
import Register from "./register.tsx";
import Report from "./report.tsx";
import Tasks from "./tasks";

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
    element: <AppLayout />,
    children: [
      {
        path: "dashboard",
        element: <Dashboard />,
      },
      {
        path: "plugins",
        element: <Plugins />,
      },
      {
        path: "findings",
        element: <Findings />,
      },
      {
        path: "tasks",
        element: <Tasks />,
      },
      {
        path: "report",
        element: <Report />,
      },
    ],
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
