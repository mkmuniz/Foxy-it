import React, { createRef, useContext, useEffect } from "react";
import MessageTemplate from "../../components/message";
import { Box, Button, Card, CardContent, Container, FormControl, Grid, Link, Stack, TextField, Typography } from "@mui/material";
import { BoxStyle, ContainerStyle, TextFieldStyle } from "./style";
import { useState } from 'react';
import { LoginForm } from "./interface";
import { login } from "../../requests/login";
import { AuthContext } from "../../auth/provider";
import { useRouter } from "next/router";
import ReCAPTCHA from "react-google-recaptcha";
import Script from 'next/script'

export default function FormLogin() {
    const { authenticate } = useContext(AuthContext);
    const router = useRouter();
    const [feedback, setFeedback] = useState({
        status: "",
        description: ""
    });
    const [form, setForm] = useState(LoginForm);
    const [captcha, setCaptcha] = useState(null);

    const getForm = (e: React.ChangeEvent<HTMLInputElement>) => {
        setForm({
            ...form,
            [e.target.name]: e.target.value
        });
    };

    const submitForm = async () => {
        try {
            const res = await login(form);

            if (!captcha) return setFeedback({
                status: 'error',
                description: 'Confirme o captcha, por favor'
            });

            switch (res.status) {
                case 200:
                    setFeedback({
                        status: "success",
                        description: "Login realizado com sucesso!"
                    });
                    authenticate(res);
                    router.replace('/');
                    break;
                case 401:
                    setFeedback({
                        status: "error",
                        description: "Email ou senha inválidos!"
                    });
                    break;
                case 404:
                    setFeedback({
                        status: "error",
                        description: "Não há usuário cadastrado com esse email."
                    });
                    break;
                default:
                    setFeedback({
                        status: "error",
                        description: "Erro desconhecido!"
                    });
                    break;
            }
        } catch (err) {
            console.log(err);

            return setFeedback({
                status: "error",
                description: "Erro ao realizar o login"
            })
        }
    };

    return <>
        <Box>
            <Script src="https://www.recaptcha.net/recaptcha/api.js" async defer />
            <Container sx={ContainerStyle}>
                <Grid>
                    <Box sx={BoxStyle}>
                        <Card>
                            <CardContent>
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
                                            type="password"
                                            onChange={getForm}
                                        />
                                        <ReCAPTCHA
                                            sitekey="6Ler4_IhAAAAAPracsrAPRsWsLio3U6FHXvsHJBK"
                                            onChange={(e) => setCaptcha(e)}
                                            onLoad={(e) => console.log(e)}
                                        />
                                        <Link sx={{ textDecoration: 'none', m: '10px' }} href="/signup">Não possui uma conta? Clique aqui</Link>
                                        <Button sx={{ m: '5px' }} onClick={submitForm} variant="contained">Log In</Button>
                                        <MessageTemplate {...feedback} />
                                    </FormControl>
                                </Stack>
                            </CardContent>
                        </Card>
                    </Box>
                </Grid>
            </Container>
        </Box>
    </>
}
