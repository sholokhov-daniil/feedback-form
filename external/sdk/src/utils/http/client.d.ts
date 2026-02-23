export interface HttpClient {
    /**
     * Generates and sends requests
     * 
     * @param {string} path 
     * @param {HttpOptions} options 
     */
    send(path: string, options: HttpOptions = {}): Promise<any>;
}

export interface HttpHeaders {
    [key: string]: string
}

export interface HttpOptions {
    headers?: Record<string, string>,
    method?: string,
    body?: any
}
