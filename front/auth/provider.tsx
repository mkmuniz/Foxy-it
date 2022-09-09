import { setCookie } from 'nookies';
import React, { createContext, useEffect, useMemo, useState } from 'react';
import jwt from 'jwt-decode';
import Cookie from 'js-cookie';

type ContextType = {
    user: any,
    authenticate: (token: any) => Promise<void>,
};

type TokenType = {
    status: number,
    data: string,
    message: string,
};
export const AuthContext = createContext({} as ContextType);

export function AuthContextProvider({ children }) {
    const [user, setUser] = useState({});
    const token = Cookie.get('token');

    useEffect(() => {
        if (token) setUser(jwt(token))
    }, []);

    async function authenticate(token: TokenType) {
        if (token.status === 200) {
            setUser(jwt(token.data));
            setCookie(undefined, 'token', token.data, {
                maxAge: 60 * 60 * 1,
            });
        }
    };

    return (
        <AuthContext.Provider value={{ user, authenticate }}>
            {children}
        </AuthContext.Provider>
    )
}
