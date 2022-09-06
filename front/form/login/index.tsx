import React from "react";
import MessageTemplate from "../../components/message";
import { Box, Button, Container, FormControl, Grid, Stack, TextField, Typography } from "@mui/material";
import { BoxStyle, ContainerStyle, TextFieldStyle } from "./style";
import { useState } from 'react';
import { LoginForm } from "./interface";
import { login } from "../../requests/login";

export default function FormLogin() {
    const [feedback, setFeedback] = useState({
        status: "",
        description: ""
    });
    const [form, setForm] = useState(LoginForm);

    const getForm = (e: React.ChangeEvent<HTMLInputElement>) => {
        setForm({
            ...form,
            [e.target.name]: e.target.value
        });
    };

    const submitForm = async () => {
        try {
            await login(form);
            return setFeedback({
                status: "success",
                description: "Login realizado com sucesso!"
            });
        } catch (err) {
            console.log(err);

            return setFeedback({
                status: "error",
                description: "Erro ao enviar formulário"
            })
        }
    };

    return <>
        <Container sx={ContainerStyle}>
            <Grid>
                <Box sx={BoxStyle}>
                    <Typography component='h3' variant='h3'>
                        Login
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
                            <MessageTemplate {...feedback} />
                            <Button onClick={submitForm} variant="contained">Log In</Button>
                        </FormControl>
                    </Stack>
                </Box>
            </Grid>
        </Container>
    </>
}
