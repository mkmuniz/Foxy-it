import { post } from './request.config';

export async function createParticipation(body: any) {
    try {
        const res = await post('/participation/create', body);
        return { ...res};
    } catch (err) {
        console.log(err);
        return err;
    }
}