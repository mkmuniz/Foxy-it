import { Box, Button, Container, FormControl, Grid, Stack, TextField, Typography } from "@mui/material";
import { BoxStyle, ContainerStyle, TextFieldStyle } from "./style";
import { useState } from 'react';


export default function FormLogin() {
    const [feedback, setFeedback] = useState("");
    const [form, setForm] = useState({
        email: "",
        password: ""
    });

    const getForm = (e: React.ChangeEvent<HTMLInputElement>) => {
        setForm({
            ...form,
            [e.target.name]: e.target.value
        });
    }

    console.log(form);
    return <>
        <Container sx={ContainerStyle}>
            <Grid>
                <Box sx={BoxStyle}>
                    <Typography align="center">
                        <h1>Login</h1>
                    </Typography>
                    <Stack>
                        <FormControl sx={{ alignItems: "center" }}>
                            <TextField sx={TextFieldStyle}
                                required
                                id="outlined-required"
                                label="E-mail"
                                name="email"
                                onChange={getForm}
                            />
                            <TextField sx={TextFieldStyle}
                                required
                                id="outlined-required"
                                label="Password"
                                name="password"
                                onChange={getForm}
                            />
                            <Button variant="contained">Contained</Button>
                        </FormControl>
                    </Stack>
                </Box>
            </Grid>
        </Container>
    </>
}