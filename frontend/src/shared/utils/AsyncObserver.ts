import {makeAutoObservable, runInAction} from "mobx";

type TrackedStatus = 'pending' | 'fulfilled' | 'rejected';

type TrackedPromise<T> = {
    promise: Promise<T>;
    status: TrackedStatus;
    data?: T;
    error?: any;
};

export class AsyncObserver {
    observed = new Map<string, TrackedPromise<any>>();

    constructor() {
        makeAutoObservable(this)
    }

    observe<T>(name: string, promise: Promise<T>) {
        const tracked: TrackedPromise<T> = {
            promise,
            status: 'pending',
        };

        runInAction(() => {
            makeAutoObservable(tracked);
        })
        this.observed.set(name, tracked);

        promise
            .then((res) => {
                runInAction(() => {
                    tracked.status = 'fulfilled';
                    tracked.data = res;
                })
            })
            .catch((err) => {
                runInAction(() => {
                    tracked.status = 'rejected';
                    tracked.error = err;
                })
            });

        return tracked;
    }

    getValue<T>(name: string): T | undefined {
        const tracked = this.observed.get(name);
        return tracked?.data;
    }

    getStatus(name: string): { isFulfilled: boolean; isRejected: boolean, isPending: boolean, status: TrackedStatus } {
        const status = this.observed.get(name)?.status ?? 'pending';
        return {
            isFulfilled: status === 'fulfilled',
            isRejected: status === 'rejected',
            isPending: status === 'pending',
            status
        };
    }

    getError(name: string): any {
        return this.observed.get(name)?.error;
    }

    isLoading(name: string): boolean {
        return this.observed.get(name)?.status === 'pending';
    }
}