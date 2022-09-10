import React from 'react';
import FormLogin from '../../form/login/index';
import withoutAuth from '../../utils/auth/withoutAuth';
function Login() {
    return <>
        <FormLogin />
    </>
}

export default withoutAuth(Login);