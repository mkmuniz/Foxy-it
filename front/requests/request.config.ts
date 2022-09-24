export async function post(url: string, body: any) {
    try {
        const response = await fetch('http://localhost:3000/api' + url, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Accept': 'application/json',
            },
            body: JSON.stringify(body)
        });

        const data = await response.json();
        
        return data;
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

export async function patch(url: string, body: any) {
    try {
        const response = await fetch('http://localhost:3000/api' + url, {
            method: 'PATCH',
            headers: {
                'Content-Type': 'application/json',
                'Accept': 'application/json',
            },
            body: JSON.stringify(body)
        });

        const data = await response.json();
        
        return data;
    } catch (err) {
        console.log(err);
        throw err;
    }
};