import { Box, Button, Container, FormControl, Grid, Stack, TextField, Typography } from "@mui/material";
import { BoxStyle, ContainerStyle, TextFieldStyle } from "./style";
import { useState } from 'react';
import { SignUpForm } from "./interface";
import React from "react";

export default function FormSignUp() {
    const [feedback, setFeedback] = useState({
        status: "",
        description: ""
    });
    const [form, setForm] = useState(SignUpForm);

    const getForm = (e: React.ChangeEvent<HTMLInputElement>) => {
        setForm({
            ...form,
            [e.target.name]: e.target.value
        });
    };

    return <>
        <div>
            <Container sx={ContainerStyle}>
                <Grid>
                    <Box sx={BoxStyle}>
                        <Typography align="center">
                            Sign Up
                        </Typography>
                        <Stack>
                            <FormControl sx={{ alignItems: "center" }}>
                                <TextField sx={TextFieldStyle}
                                    required
                                    id="outlined-required"
                                    label="Username"
                                    name="username"
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
                                    onChange={getForm}
                                />
                                <Button variant="contained">Sign Up</Button>
                            </FormControl>
                        </Stack>
                    </Box>
                </Grid>
            </Container>
        </div>
    </>
}
