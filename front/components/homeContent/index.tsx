import { Button, Card, CardContent, CardMedia, Grid } from '@mui/material';
import { Container } from '@mui/system';
import AddIcon from '@mui/icons-material/Add';
import React from 'react';

export default function HomeContent() {
    return <>
        <Container sx={{ m: '30px' }}>
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