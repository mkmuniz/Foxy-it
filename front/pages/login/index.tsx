import React from 'react';
import FormLogin from '../../form/login/index';

export default function Login() {
    return <>
        <FormLogin />
    </>
}

export const getServerSideProps: any = async ({ req }) => {
    const { token } = req.cookies;

    if (token) {
        return {
            redirect: {
                destination: '/',
                permanent: false,
            }
        }
    }

    return {
        props: {},
    }
}