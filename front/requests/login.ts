import { post } from './request.config';

export async function login(body: any) {
    try {
        const res = await post('/user', body);
        console.log(res);
        return { ...res};
    } catch (err) {
        console.log(err);
        throw err;
    }
};