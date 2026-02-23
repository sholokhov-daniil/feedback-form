import { Config } from "../config/index";

export const Logger = new class {
    public debug(...args: any): void {
        if (this.enabled) {
            console.debug(this.prefix, ...args);
        }
    }

    public error(...args: any): void {
        console.error(this.prefix, ...args);
    }

    public get enabled(): boolean {
        return Config.debug;
    }

    private get prefix() {
        return '[Sholokhov form] ';
    }
};