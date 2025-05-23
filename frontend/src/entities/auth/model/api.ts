import axios, {AxiosInstance} from "axios";
import type {AuthService as AuthServiceType, LoginDTO, RegisterDTO} from "@entities/auth/model/types.ts";

export class AuthService implements AuthServiceType {
    axios: AxiosInstance = axios.create({
        withCredentials: true,
    });


    async login(data: LoginDTO) {
        return this.axios.post("http://localhost:8080/api/auth/login", data)
    }

    async register(dto: RegisterDTO) {
        return this.axios.post("http://localhost:8080/api/auth/register", dto)
    }
}