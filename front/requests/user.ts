import { post, patch } from './request.config';

export async function createUser(body: any) {
    try {
        const res = await post('/user/create', body);
        return { ...res};
    } catch (err) {
        console.log(err);
        return err;
    }
};

export async function login(body: any) {
    try {
        const res = await post('/login', body);
        return { ...res};
    } catch (err) {
        console.log(err);
        return err;
    }
};

export async function updateUser(id: string, body: any) {
    try {
        const res = await patch(`/user/update/${id}`, body);
        return { ...res};
    } catch (err) {
        console.log(err);
        return err;
    }
}