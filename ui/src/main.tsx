import { createRoot } from "react-dom/client";
import { RouterProvider } from "react-router/dom";
import './index.css'
import router from './routes/router.ts';

const root = document.getElementById('root')!;

createRoot(root).render(
  <RouterProvider router={router} />,
);

