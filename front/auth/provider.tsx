import { setCookie } from 'nookies';
import React, { createContext, useState } from 'react';
import jwt from 'jwt-decode';

type ContextType = {
    user: null,
    authenticate: (token: any) => Promise<void>,
    isAuthenticated: boolean,
};

type TokenType = {
    status: number,
    data: string,
    message: string,
};

export const AuthContext = createContext({} as ContextType);

export function AuthContextProvider({ children }) {
    const [user, setUser] = useState(null);
    const [isAuthenticated, setAuthentication] = useState(false);

    async function authenticate(token: TokenType) {
        if (token.status === 200) {
            setUser(jwt(token.data));
            setCookie(undefined, 'token', token.data, {
                maxAge: 60 * 60 * 1,
            });
            setAuthentication(true);
        }
    };

    return (
        <AuthContext.Provider value={{ user, authenticate, isAuthenticated }}>
            {children}
        </AuthContext.Provider>
    )
}
