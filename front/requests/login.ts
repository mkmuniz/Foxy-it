import { post } from './request.config';

export async function login(body: any) {
    try {
        const response = await post('/user', body);
        return response;
    } catch (err) {
        console.log(err);
        throw err;
    }
};