import { Box, Grid, Container } from '@mui/material';
import React from 'react';
import Navbar from '../../components/navbar';
import { ChatStyle, ScreenStyle, UserStyle } from './style';

export default function Room() {
    return <>
        <Box>
            <Navbar />
            <Grid container sx={{ alignItems: 'center', textAlign: 'center' }}>
                <Grid xs={8}>
                    <Container sx={ScreenStyle}>
                        Teste
                    </Container>
                    <Container sx={UserStyle}>
                        <Box>
                            Users
                        </Box>
                    </Container>
                </Grid>
                <Grid xs={4}>
                    <Box sx={ChatStyle}>
                        Chat
                    </Box>
                </Grid>
            </Grid>
        </Box>
    </>
}