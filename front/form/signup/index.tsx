import { Box, Button, Container, FormControl, Grid, Stack, TextField, Typography } from "@mui/material";
import { BoxStyle, ContainerStyle, TextFieldStyle } from "./style";
import { useState } from 'react';
import React from "react";
import { createUser } from "../../requests/login";
import MessageTemplate from "../../components/message";
import { useRouter } from "next/router";

type SignUpFormType = {
    name: string,
    email: string,
    confirmEmail: string,
    password: string
};

export default function FormSignUp() {
    const [feedback, setFeedback] = useState({
        status: "",
        description: ""
    });
    const [form, setForm] = useState({} as SignUpFormType);
    const router = useRouter();

    const getForm = (e: React.ChangeEvent<HTMLInputElement>) => {
        setForm({
            ...form,
            [e.target.name]: e.target.value
        });
    };

    const submitForm = async () => {
        if (form.name === '' || !form.name || form.password === '' || !form.password) {
            return setFeedback({
                status: 'error',
                description: 'Fill out all the fields!'
            })
        }
        if (form.confirmEmail != form.email) {
            return setFeedback({
                status: 'error',
                description: 'E-mail are not equals!'
            })
        };

        try {
            await createUser(form);
            setFeedback({
                status: 'success',
                description: 'User created successfully'
            })
            return router.replace('/login')
        } catch (err) {
            console.log(err);

            return setFeedback({
                status: 'error',
                description: 'Error on create account, try again'
            })
        }
    }

    return <>
        <div>
            <Container sx={ContainerStyle}>
                <Grid>
                    <Box sx={BoxStyle}>
                        <Typography variant="h4" align="center">
                            Sign Up
                        </Typography>
                        <Stack>
                            <FormControl sx={{ alignItems: "center" }}>
                                <TextField sx={TextFieldStyle}
                                    required
                                    id="outlined-required"
                                    label="Username"
                                    name="name"
                                    onChange={getForm}
                                />
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
                                    label="Confirm E-mail"
                                    name="confirmEmail"
                                    onChange={getForm}
                                />
                                <TextField sx={TextFieldStyle}
                                    required
                                    id="outlined-required"
                                    label="Password"
                                    name="password"
                                    type="password"
                                    onChange={getForm}
                                />
                                <MessageTemplate {...feedback} />
                                <Button onClick={submitForm} variant="contained">Sign Up</Button>
                            </FormControl>
                        </Stack>
                    </Box>
                </Grid>
            </Container>
        </div>
    </>
}
