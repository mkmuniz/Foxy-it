import { useRouter } from "next/router";
import { ElementType, useEffect } from "react";
import Cookie from 'js-cookie';

export default function withAuth(WrapperComponent: ElementType) {
    const Wrapper = (props: unknown) => {
        const router = useRouter();
        
        useEffect(() => {
            const token = Cookie.get('token');

            if(!token) router.replace('/login')
        }, [])
    };

    return Wrapper;
}