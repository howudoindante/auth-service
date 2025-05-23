import React from "react";

export class ErrorBoundary extends React.Component {
    state: { hasError: boolean } = {hasError: false};

    constructor(public props: { children: React.ReactNode }) {
        super(props);
    }

    static getDerivedStateFromError() {
        // Update state so the next render will show the fallback UI.
        return {hasError: true};
    }

    componentDidCatch() {
        // You can also log the error to an error reporting service
    }

    render() {
        if (this.state.hasError) {
            // You can render any custom fallback UI
            return <h1>Something went wrong.</h1>;
        }

        return this.props.children;
    }
}