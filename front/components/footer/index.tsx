import React from 'react';
import { Grid, BottomNavigation, Typography } from '@mui/material';
import { BoxFooterStyle } from './style';

export default function Footer() {
    return <>
        <BottomNavigation sx={BoxFooterStyle}>
            <Grid container>
                <Grid xs={4}>
                    <Typography variant='h6'>
                        @Copyright 2022
                    </Typography>
                </Grid>
            </Grid>
        </BottomNavigation>
    </>
}