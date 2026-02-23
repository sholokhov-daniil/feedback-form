import type { HttpClient } from '../utils/http/client';
import { Logger } from '../utils/logger/index';
import { Config } from '../utils/config/index';

// TODO: в работе
export class Init {
    private initialized: boolean = false;
    private HttpClient client;

    constructor(client: HttpClient) {
        this.client = client;
    }

    run(...args: any): void {
        const siteId = String(args[0]);

        if (this.initialized) {
            return;
        }

        Logger.debug('Initializing SDK for project', siteId);

        this.client.send(
            '/sdk/init',
            {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    'VERSION-SDK': Config.version,
                },
                credentials: 'include',
                body: { siteId }
            }
        )
    }
}