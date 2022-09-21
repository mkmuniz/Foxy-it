import { Box, Card, CardContent, Grid, Typography } from '@mui/material';
import React from 'react';

export default function CreateRoom() {
    return <>
        <Box>
            <Grid container spacing={2} direction="column" alignItems="center" sx={{ mt: '10%'}}>
                <Grid item xs={12} justifyContent="center">
                    <Card>
                        <CardContent>
                            <Typography variant='h6'>
                                Create room
                            </Typography>
                        </CardContent>
                    </Card>
                </Grid>
            </Grid>
        </Box>
    </>;
}