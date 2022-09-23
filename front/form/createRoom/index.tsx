import { Box, Card, CardContent, Grid, Typography, CardMedia, Button, FormControl, TextField } from '@mui/material';
import React, { useState } from 'react';
import MessageTemplate from '../../components/message';
import { TextFieldStyle } from '../login/style';
import { createRoom } from '../../requests/room';

type Form = {
    name: string,
}

export default function CreateRoom() {
    const [form, setForm] = useState({} as Form);
    const [feedback, setFeedback] = useState({
        status: "",
        description: ""
    })

    const getValues = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
        setForm({ ...form, [e.target.name]: e.target.value });
    };

    const submit = async (e: React.MouseEvent<HTMLButtonElement, MouseEvent>) => {
        e.preventDefault();

        try {
            await createRoom(form);

            return setFeedback({
                status: "success",
                description: "Sala criada com sucesso!"
            });
        } catch(err) {
            console.log(err);
        }
    }
    
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
                                    onChange={(e) => getValues(e)}
                                />
                                <MessageTemplate {...feedback} />
                                <Button onClick={e => submit(e)} variant="contained">Create</Button>
                            </FormControl>
                        </CardContent>
                    </Card>
                </Grid>
            </Grid>
        </Box>
    </>;
}