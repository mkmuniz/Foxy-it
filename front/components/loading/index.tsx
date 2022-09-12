import { Typography, LinearProgress } from '@mui/material';
import React from 'react';

export default function Loading() {
    return <>
        <Typography variant="h3">
            Loading...
        </Typography>
        <LinearProgress />
    </>
}