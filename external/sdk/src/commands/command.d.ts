import type { HttpClient } from '../utils/http/client';

export interface CommandInterface {
    execute(options: object): void
}

export type CommandConstructor = new (client: HttpClient) => CommandInterface;

export type CommandRepository = Map<string, CommandInterface>