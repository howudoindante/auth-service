import {AxiosInstance, AxiosResponse} from "axios";

export interface LoginDTO {
    username: string;
    password: string;
}

export interface RegisterDTO {
    username: string;
    password: string
    email: string;
}

export interface AuthService {
    axios: AxiosInstance;

    login(dto: LoginDTO): Promise<AxiosResponse<any, any>>;

    register(dto: RegisterDTO): Promise<AxiosResponse<any, any>>;
}