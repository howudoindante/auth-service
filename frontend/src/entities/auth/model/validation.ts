import * as yup from 'yup';

export const registerSchema = yup.object().shape({
    email: yup.string().required('Email is required'),
    username: yup.string().required('Username is required'),
    password: yup.string().required('Password is required')
})
export const loginSchema = yup.object().shape({
    username: yup.string().required('Username is required'),
    password: yup.string().required('Password is required')
})