import { post } from './request.config';

export async function createRoom(body: any) {
    try {
        const res = await post('/room/room', body);
        return { ...res};
    } catch (err) {
        console.log(err);
        return err;
    }
}