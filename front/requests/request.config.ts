export async function post(url: string, body: any) {
    try {
        const response = await fetch(url, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Accept': 'application/json',
            },
            body: JSON.stringify(body)
        });

        return response;
    } catch (err) {
        console.log(err);
        throw err;
    }
};

export async function get(url: string) {
    try {
        const response = await fetch(url, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
                'Accept': 'application/json',
            }
        });

        return response;
    } catch (err) {
        console.log(err);
        throw err;
    }
};