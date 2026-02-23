import { Config } from '../config/index';
import { Logger } from '../logger/index';
import type { HttpClient, HttpHeaders, HttpOptions } from './client';

export class Client implements HttpClient {
    private csrfToken: string = '';

    /**
     * Generates and sends requests
     * 
     * @param {string} path 
     * @param {HttpOptions} options 
     */
    public async send(path: string, options: HttpOptions = {}): Promise<any> {
        const headers = this.normalizeHeaders(options);
        const method = options.method || 'POST';
        const credentials = 'include';
        const body = this.normalizeBody(options);
        const url = this.getUrl(path);

        const response = await fetch(
            url,
            {
                method,
                headers,
                credentials,
                body
            }
        )

        Logger.debug(response);
    }

    /**
     * Generates request headers
     * 
     * @param {HttpOptions} options 
     * @return {HttpHeaders}
     */
    private normalizeHeaders(options: HttpOptions): HttpHeaders {
        const headers = Object.assign({}, options.headers || {});

        if (this.csrfToken) {
            headers['X-CSRF-Token'] = this.csrfToken;
        }

        return headers;
    }

    /**
     * Generates the request body
     * 
     * @param {HttpOptions} options 
     * @return {string}
     */
    private normalizeBody(options: HttpOptions): string {
        return options.body ? JSON.stringify(options.body) : "";
    }

    /**
     * Generates the full request URL
     * 
     * @param {string} path
     * @return {string}
     */
    private getUrl(path: string): string {
        return this.target + path;
    }

    /**
     * Returns the URL of the data exchange service
     * 
     * @return {string}
     */
    private get target(): string {
        return Config.api;
    }
}