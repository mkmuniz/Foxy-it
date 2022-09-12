import React from 'react';
import { Box, Grid, BottomNavigation, Typography, ListItem, ListItemText, List } from '@mui/material';
import { BoxFooterStyle } from './style';

export default function Footer() {
    return <>
        <BottomNavigation sx={BoxFooterStyle}>
            <Grid container>
                <Grid xs={4}>
                    <Typography variant='h4'>
                        About
                    </Typography>
                    <List>
                        <ListItem>
                            <ListItemText>
                                Teste
                            </ListItemText>
                        </ListItem>
                        <ListItem>
                            <ListItemText>
                                Teste
                            </ListItemText>
                        </ListItem>
                        <ListItem>
                            <ListItemText>
                                Teste
                            </ListItemText>
                        </ListItem>
                    </List>
                </Grid>
                <Grid xs={4}>
                    <Typography variant='h4'>
                        Contact
                    </Typography>
                    <ListItem>
                        <ListItemText>
                            Teste
                        </ListItemText>
                    </ListItem>
                    <ListItem>
                        <ListItemText>
                            Teste
                        </ListItemText>
                    </ListItem>
                    <ListItem>
                        <ListItemText>
                            Teste
                        </ListItemText>
                    </ListItem>
                </Grid>
                <Grid xs={3}>
                    <Typography variant='h4'>
                        Team
                    </Typography>
                    <ListItem>
                        <ListItemText>
                            Teste
                        </ListItemText>
                    </ListItem>
                    <ListItem>
                        <ListItemText>
                            Teste
                        </ListItemText>
                    </ListItem>
                    <ListItem>
                        <ListItemText>
                            Teste
                        </ListItemText>
                    </ListItem>
                </Grid>
            </Grid>
        </BottomNavigation>
    </>
}