import jwt from 'jwt-decode';

export function validateCookie(token: string) {
    const decoded: any = jwt(token);

    if (decoded.exp < Date.now() / 1000) {
        return false;
    }

    return true;
};

export function removeCookie() {
    document.cookie = 'token=; expires=Thu, 01 Jan 1970 00:00:01 GMT;';
};