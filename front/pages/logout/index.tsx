import { Container } from '@mui/material';
import { useRouter } from 'next/router';
import React, { useEffect } from 'react';
import Loading from '../../components/loading';
import { removeCookie } from '../../utils/cookies/cookie';

export default function LogOut() {
    const router = useRouter();

    useEffect(() => {
        removeCookie();
        router.replace('/login').then(() => router.reload());
    }, []);

    return <>
        <Container sx={{ textAlign: 'center', mt: '15%' }}>
            <Loading />
        </Container>
    </>
}