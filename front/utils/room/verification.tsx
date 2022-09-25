import { useRouter } from "next/router";
import { ElementType, useEffect } from "react";

export default function isVerified(WrappedComponent: ElementType) {
    const Wrapper = (props: unknown) => {

        return <WrappedComponent />
    };

    return Wrapper;
}