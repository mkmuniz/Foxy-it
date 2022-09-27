import { ElementType } from "react";

export default function isVerified(WrappedComponent: ElementType) {
    const Wrapper = (props: unknown) => {

        return <WrappedComponent />
    };

    return Wrapper;
}