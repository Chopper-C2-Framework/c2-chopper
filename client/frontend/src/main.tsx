import React from "react";
import ReactDOM from "react-dom/client";
import "./styles/global.css";
import {RouterProvider} from "react-router-dom";
import router from "./routes";
import { QueryClient, QueryClientProvider, useQuery } from 'react-query'

const queryClient = new QueryClient()

ReactDOM.createRoot(document.getElementById("root") as HTMLElement).render(
  <React.StrictMode>
    <QueryClientProvider client={queryClient}>
      <RouterProvider router={router}/>
      </QueryClientProvider>
  </React.StrictMode>
);
