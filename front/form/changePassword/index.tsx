import { Typography, TextField } from '@mui/material';
import React, { useContext } from 'react';
import { AuthContext } from '../../auth/provider';
import { TextFieldStyle } from '../login/style';

export default function ChangePassword() {
    return <>
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
    </>
}