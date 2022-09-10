import { Container, Card, CardMedia, CardContent, Grid, List, ListItem, Link } from '@mui/material';
import React, { useContext, useState } from 'react';
import Navbar from '../../components/navbar';
import withAuth from '../../utils/auth/withAuth';
import { AuthContext } from '../../auth/provider';
import EditPerfil from '../../form/editPerfil';
import ChangePassword from '../../form/changePassword';

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
                            <List>
                                <ListItem>
                                    <Link onClick={() => setComponent('edit')} sx={{ textDecoration: 'none', color: 'black' }}>
                                        Edit perfil
                                    </Link>
                                </ListItem>
                                <ListItem>
                                    <Link onClick={() => setComponent('changepassword')} sx={{ textDecoration: 'none', color: 'black' }}>
                                        Change password
                                    </Link>
                                </ListItem>
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