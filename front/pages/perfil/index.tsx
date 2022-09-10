import { Container, Card, CardMedia, CardContent, Grid, Avatar, List, ListItem, ListItemAvatar, ListItemText, Typography } from '@mui/material';
import React from 'react';
import Navbar from '../../components/navbar';
import AccountBox from '@mui/icons-material/AccountBox';
import withAuth from '../../utils/auth/withAuth';

function Perfil() {
    return <>
        <Navbar />
        <Container sx={{ mt: '40px' }}>
            <Card>
                <CardContent>
                    <Grid container spacing={2}>
                        <Grid xs={6}>
                            <Grid container spacing={2}>
                                <Grid item xs={12} md={6}>
                                    <List>
                                        <ListItem>
                                            <Avatar>
                                                <AccountBox />
                                            </Avatar>
                                            Perfil
                                        </ListItem>
                                        <ListItem>
                                            asfasfdsaf
                                        </ListItem>
                                        <ListItem>
                                            asfasfdsaf
                                        </ListItem>
                                        <ListItem>
                                            asfasfdsaf
                                        </ListItem>
                                    </List>
                                </Grid>
                            </Grid>
                            <Grid xs={6}>
                                retegerg
                            </Grid>
                        </Grid>
                    </Grid>
                </CardContent>
            </Card>
        </Container>
    </>;
};

export default withAuth(Perfil);