import { Alert } from '@mui/material';
import React from 'react';


interface MessageProps {
    status: string;
    description: string;
};

export default function MessageTemplate(message: MessageProps) {

    return <>
        {message.status === "success" && <>
            <Alert severity="success">{message.description}</Alert>
        </>
        }
        {message.status === "error" && <>
            <Alert severity="error">{message.description}</Alert>
        </>
        }
    </>
}