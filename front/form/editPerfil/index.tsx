import React, { useContext } from 'react';
import { Typography, TextField } from '@mui/material';
import { TextFieldStyle } from '../login/style';
import { AuthContext } from '../../auth/provider';

export default function EditPerfil() {
    const { user } = useContext(AuthContext);

    return <>
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
    </>
}