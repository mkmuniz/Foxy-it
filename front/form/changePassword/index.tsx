import { Typography, TextField, Button, FormControl } from '@mui/material';
import React, { useContext } from 'react';
import { AuthContext } from '../../auth/provider';
import { TextFieldStyle } from '../login/style';

export default function ChangePassword() {
    return <>
        <FormControl sx={{ alignItems: 'center', width: '75%' }}>
            <Typography variant="h4">
                Your password
            </Typography>
            <TextField sx={TextFieldStyle}
                required
                id="outlined-required"
                label="Password"
                name="name"
            />
            <TextField sx={TextFieldStyle}
                required
                id="outlined-required"
                label="New password"
                name="email"
            />
            <Button variant="contained">Save</Button>
        </FormControl>
    </>
}