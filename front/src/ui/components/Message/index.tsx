import { Alert } from "@mui/material";
import MessageTemplateProps from "./interface";

export function MessageTemplate(props: MessageTemplateProps) {
    return <>
        {props.message.status === 'success' && (<Alert severity="success">{props.message.description}</Alert>)}
        {props.message.status === 'error' && (<Alert severity="error">{props.message.description}</Alert>)}
    </>
}