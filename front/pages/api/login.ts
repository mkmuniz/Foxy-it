import { generateCookie } from './../../utils/cookies/cookie';

export let AuthToken: any = null;

export default async function handler(req, res) {
  try {
    const { body } = req;
    const response = await fetch('http://localhost:4001/login', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Accept': 'application/json',
      },
      body: JSON.stringify(body)
    });

    const data = await response.json();

    AuthToken = data;

    generateCookie(data, 'token', { expires: new Date(Date.now() + 86400e3), httpOnly: true });
    res.status(200).send(data)
  } catch (err) {
    console.log(err);

    throw err;
  }
}
