import {makeAutoObservable, reaction} from "mobx";
import {AuthService} from "@entities/auth/model/api.ts";
import {AsyncObserver} from "@shared/utils/AsyncObserver.ts";
import {toaster} from "@components/ui/toaster.tsx";
import {LoginDTO, RegisterDTO} from "@entities/auth/model/types.ts";

class Authentication {

    authService = new AuthService()
    asyncObserver = new AsyncObserver()

    constructor() {
        makeAutoObservable(this)
    }

    login = (dto: LoginDTO) => {
        const promise = this.authService.login(dto)


        this.asyncObserver.observe("login", promise)
    }


    register = (dto: RegisterDTO) => {
        const promise = this.authService.register(dto)


        this.asyncObserver.observe("register", promise)
    }
}

const authStore = new Authentication();


reaction(
    () => authStore.asyncObserver.getStatus("login").status,
    (status) => {
        if (status === 'fulfilled') {
            toaster.create({title: "Successfully login into your account", type: "success"});
        } else if (status === 'rejected') {
            const err = authStore.asyncObserver.getError("login");
            toaster.create({
                title: "Cannot login to your account",
                description: err.response.data.message ?? err.response.data.error,
                type: "error",
                action: {
                    label: "Info",
                    onClick: () => console.log("Undo"),
                },
            });
        }
    }
);

reaction(
    () => authStore.asyncObserver.getStatus("register").status,
    (status) => {
        if (status === 'fulfilled') {
            toaster.create({title: "Successfully register your account", type: "success"});
        } else if (status === 'rejected') {
            const err = authStore.asyncObserver.getError("register");
            toaster.create({
                title: "Cannot register your account",
                description: err.response.data.message ?? err.response.data.error,
                type: "error",
                action: {
                    label: "Info",
                    onClick: () => console.log("Undo"),
                },
            });
        }
    }
);


export default authStore;
