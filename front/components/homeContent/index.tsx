import { Button, Card, CardContent, CardMedia, Grid, Typography } from '@mui/material';
import { Container } from '@mui/system';
import AddIcon from '@mui/icons-material/Add';
import React from 'react';

export default function HomeContent() {
    return <>
        <Container sx={{ m: '30px', textAlign: 'center', width: '100%', minHeight: '55vh', height: '100%'}}>
            <Typography sx={{ m: '15px' }} variant="h4" >My rooms</Typography>
            <Card>
                <CardMedia>
                    <Grid container justifyContent='flex-end'>
                        <Button sx={{ m: '10px', bgcolor: '#FD2626' }} variant="contained" color='error'><AddIcon />Criar sala</Button>
                    </Grid>
                </CardMedia>
                <CardContent>

                </CardContent>
            </Card>
        </Container>
    </>
} 