import { Container, LinearProgress, Typography } from '@mui/material';
import { useRouter } from 'next/router';
import React, { useEffect } from 'react';
import { removeCookie } from '../../utils/cookies/cookie';

export default function LogOut() {
    const router = useRouter();

    useEffect(() => {
        removeCookie();
        router.replace('/login');
    }, []);

    return <>
        <Container sx={{ textAlign: 'center', mt: '15%' }}>
            <Typography variant="h3">
                Logging out..
            </Typography>
            <LinearProgress />
        </Container>
    </>
}