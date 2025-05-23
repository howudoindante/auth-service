function classNames(...args: Array<string | false | null | undefined | Record<string, boolean>>): string {
    return args
        .flatMap(arg => {
            if (typeof arg === 'string') return [arg];
            if (typeof arg === 'object' && arg !== null) {
                return Object.entries(arg)
                    .filter(([_, value]) => Boolean(value))
                    .map(([key]) => key);
            }
            return [];
        })
        .join(' ');
}


export default classNames;