import { Box, Button, Card, CardContent, CardMedia, Grid, Modal, Typography } from '@mui/material';
import { Container } from '@mui/system';
import AddIcon from '@mui/icons-material/Add';
import React from 'react';
import CreateRoom from '../../form/createRoom';

export default function HomeContent() {
    const [open, setOpen] = React.useState(false);
    const handleOpen = () => {
        setOpen(true);
    };
    const handleClose = () => {
        setOpen(false);
    };
    return <>
        <Container sx={{ m: '30px', textAlign: 'center', width: '100%', minHeight: '77vh', height: '100%' }}>
            <Typography sx={{ m: '15px' }} variant="h4" >My rooms</Typography>
            <Card>
                <CardMedia>
                    <Grid container justifyContent='flex-end'>
                        <Button onClick={handleOpen} sx={{ m: '10px', bgcolor: '#FD2626' }} variant="contained" color='error'><AddIcon />Criar sala</Button>
                        <Modal
                            open={open}
                            onClose={handleClose}
                            aria-labelledby="parent-modal-title"
                            aria-describedby="parent-modal-description"
                        >
                            <Box sx={{ width: "100%" }}>
                                <CreateRoom />
                            </Box>
                        </Modal>
                    </Grid>
                </CardMedia>
                <CardContent>

                </CardContent>
            </Card>
        </Container>
    </>
} 