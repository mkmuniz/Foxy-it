import { post } from './request.config';

export async function login(body: any) {
    try {
        const res = await post('/login', body);
        return { ...res};
    } catch (err) {
        console.log(err);
        throw err;
    }
};