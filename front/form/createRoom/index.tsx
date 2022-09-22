import { Box, Card, CardContent, Grid, Typography, CardMedia, Button, FormControl, TextField } from '@mui/material';
import React, { useState } from 'react';
import { TextFieldStyle } from '../login/style';

type Form = {
    name: string,
}

export default function CreateRoom() {
    const [form, setForm] = useState({} as Form);

    const getValues = (e: any) => {
        setForm({ ...form, [e.target.name]: e.target.value });
    };

    console.log(form);
    
    return <>
        <Box>
            <Grid container spacing={2} direction="column" alignItems="center" sx={{ mt: '10%' }}>
                <Grid item xs={12} justifyContent="center" sx={{ width: '40% '}}>
                    <Card>
                        <CardMedia sx={{ textAlign: 'center'}}>
                            <Typography variant='h5'>
                                Create a Room
                            </Typography>
                        </CardMedia>
                        <CardContent>
                            <FormControl sx={{ alignItems: 'center', width: '100%' }}>
                                <TextField sx={TextFieldStyle}
                                    required
                                    id="outlined-required"
                                    label="Room name"
                                    name="name"
                                    onChange={getValues}
                                />
                                <Button variant="contained">Create</Button>
                            </FormControl>
                        </CardContent>
                    </Card>
                </Grid>
            </Grid>
        </Box>
    </>;
}