export default interface ISignUpForm {
    username: string,
    email: string;
    password: string;
    confirmPassword: string;
}

export const SignUpForm: ISignUpForm = {
    username: '',
    email: '',
    password: '',
    confirmPassword: ''
}