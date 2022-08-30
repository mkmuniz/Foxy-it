import { useContext } from "react";
import PrivateRoutes from "./private";
import PublicRoutes from "./public";

export default function Routes() {
    const context: any = {};
    const authContext = useContext(context);
    {authContext ? (
        <PrivateRoutes />
    ) : (
        <PublicRoutes />
    )}
}