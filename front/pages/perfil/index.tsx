import { Container, Card, CardMedia, CardContent, Grid, List, ListItem, Link, ListItemButton, ListItemIcon, ListItemText } from '@mui/material';
import React, { useContext, useState } from 'react';
import Navbar from '../../components/navbar';
import withAuth from '../../utils/auth/withAuth';
import { AuthContext } from '../../auth/provider';
import EditPerfil from '../../form/editPerfil';
import ChangePassword from '../../form/changePassword';
import KeyIcon from '@mui/icons-material/Key';
import AccountCircleIcon from '@mui/icons-material/AccountCircle';

function Perfil() {
    const { user } = useContext(AuthContext);
    const [component, setComponent] = useState('');

    return <>
        <Navbar />
        <Container sx={{ mt: '40px' }}>
            <Card>
                <CardContent>
                    <Grid container spacing={2}>
                        <Grid xs={4}>
                            <List component="nav" aria-label="main mailbox folders">
                                <ListItemButton onClick={() => setComponent('edit')}>
                                    <ListItemIcon>
                                        <AccountCircleIcon />
                                    </ListItemIcon>
                                    <ListItemText primary="Edit perfil" />
                                </ListItemButton>
                                <ListItemButton onClick={() => setComponent('changepassword')}>
                                    <ListItemIcon>
                                        <KeyIcon />
                                    </ListItemIcon>
                                    <ListItemText primary="Change password" />
                                </ListItemButton>
                            </List>
                        </Grid>
                        <Grid xs={8} sx={{ alignItems: 'center', textAlign: 'center' }}>
                            {component === 'edit' &&
                                <>
                                    <EditPerfil />
                                </>
                            }
                            {component === 'changepassword' &&
                                <>
                                    <ChangePassword />
                                </>
                            }
                        </Grid>
                    </Grid>
                </CardContent>
            </Card>
        </Container>
    </>;
};

export default withAuth(Perfil);