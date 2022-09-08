import React, { createContext, useState } from 'react';
import { AuthToken } from '../pages/api/login';
import { authenticatingUser } from '../utils/cookies/cookie';
import { ContextInterface } from './interfaces';

export let AuthContext: any = createContext({} as ContextInterface);

export function AuthContextProvider({ children }) {
    const [user, setUser] = useState(null);

    if (AuthContext) {
        AuthContext = authenticatingUser(AuthToken, 'token');
    }

    return (
        <AuthContext.Provider value={AuthContext}>
            {children}
        </AuthContext.Provider>
    )
}
