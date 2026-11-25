import { createBrowserRouter } from "react-router";
import AboutPage from "../pages/AboutPage";
import HomePage from "../pages/HomePage";
import App from "../App";

const router = createBrowserRouter([
    {
        path: '/',
        Component: App,
        children: [
            { 
                index: true,
                Component: HomePage,
             },
             {
                path: 'about',
                Component: AboutPage,
             }
        ],
    },
]);

export default router;

