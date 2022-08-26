import { Box, Button, Container, FormControl, Grid, Stack, TextField, Typography } from "@mui/material";
import { BoxStyle, ContainerStyle, TextFieldStyle } from "./style";


export default function FormLogin() {

    return <>
        <Container sx={ContainerStyle}>
            <Grid>
                <Box sx={BoxStyle}>
                    <Typography align="center">
                        <h1>Login</h1>
                    </Typography>
                    <Stack>
                        <FormControl sx={{ alignItems: "center" }}>
                            <TextField sx={TextFieldStyle}
                                required
                                id="outlined-required"
                                label="Username"
                            />
                            <TextField sx={TextFieldStyle}
                                required
                                id="outlined-required"
                                label="Password"
                            />
                            <Button variant="contained">Contained</Button>
                        </FormControl>
                    </Stack>
                </Box>
            </Grid>
        </Container>
    </>
}