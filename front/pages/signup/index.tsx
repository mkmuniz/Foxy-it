import React from 'react';
import FormSignUp from '../../form/signup/index';
import withoutAuth from '../../utils/auth/withoutAuth';

function SignUp() {
    return <>
        <FormSignUp />
    </>
}

export default withoutAuth(SignUp);