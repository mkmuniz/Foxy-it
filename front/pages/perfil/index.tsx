import { Container, Card, CardMedia, CardContent, Grid, Avatar, List, ListItem, ListItemAvatar, ListItemText, Typography, TextField, Button } from '@mui/material';
import React, { useContext } from 'react';
import Navbar from '../../components/navbar';
import withAuth from '../../utils/auth/withAuth';
import { TextFieldStyle } from '../../form/login/style';
import { AuthContext } from '../../auth/provider';

function Perfil() {
    const { user } = useContext(AuthContext);

    return <>
        <Navbar />
        <Container sx={{ mt: '40px' }}>
            <Card>
                <CardContent>
                    <Grid container spacing={2}>
                        <Grid xs={4}>
                            <List>
                                <ListItem>
                                    Edit perfil
                                </ListItem>
                                <ListItem>
                                    Change password
                                </ListItem>
                            </List>
                        </Grid>
                        <Grid xs={8} sx={{ alignItems: 'center', textAlign: 'center' }}>
                            <Typography variant="h4">
                                Your perfil
                            </Typography>
                            <TextField sx={TextFieldStyle}
                                required
                                id="outlined-required"
                                label="Username"
                                name="name"
                                value={user.Name}
                            />
                            <TextField sx={TextFieldStyle}
                                required
                                id="outlined-required"
                                label="E-mail"
                                name="email"
                                value={user.Email}
                            />
                        </Grid>
                    </Grid>
                </CardContent>
            </Card>
        </Container>
    </>;
};

export default withAuth(Perfil);