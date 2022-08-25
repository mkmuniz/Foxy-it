import { Box, Container, Stack, TextField, Typography } from "@mui/material";
import { BoxStyle, ContainerStyle, FormControlStyle } from "./style";
import FormControl from '@mui/material/FormControl';


export default function FormLogin() {

    return <>
        <Container sx={ContainerStyle}>
            <Box sx={BoxStyle}>
                <Typography>
                    <h1>Login</h1>
                </Typography>
                <Stack>
                    <TextField
                        required
                        id="outlined-required"
                        label="Required"
                        defaultValue="Hello World"
                        />
                        <TextField
                        required
                        id="outlined-required"
                        label="Required"
                        defaultValue="Hello World"
                        />
                        <TextField
                        required
                        id="outlined-required"
                        label="Required"
                        defaultValue="Hello World"
                        />
                </Stack>
            </Box>
        </Container>
    </>
}