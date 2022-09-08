import jwt from 'jwt-decode';

export const generateCookie = (token: any = {}, name: string, options: any = {}) => {
    options = {
        path: '/',
        ...options
    };

    if (options.expires instanceof Date) {
        options.expires = options.expires.toUTCString();
    }

    if (token.status === 200) {
        let updatedCookie = encodeURIComponent(name) + '=' + encodeURIComponent(token.data);

        for (let optionKey in options) {
            updatedCookie += '; ' + optionKey;
            let optionValue = options[optionKey];
            if (optionValue !== true) {
                updatedCookie += '=' + optionValue;
            }
        }
        return updatedCookie;
    }
}

export const authenticatingUser = (token: any = {}, name: string) => {
    const isValid = validateCookie(name);
    const userData = jwt(token);

    if (isValid) {
        return {
            user: { userData },
            isAuthenticated: true,
            token: token.data
        }
    }
}

export const validateCookie = (name: string) => {
    if (process.browser) {
        var cookie = document.cookie.split(';').find((c) => c.trim().startsWith(name + '='));
    }
    if (!cookie) {
        return false;
    }
}

export const deleteCookie = (token: any = {}, name: string) => {
    generateCookie(name, '', {
        'max-age': -1
    });
};