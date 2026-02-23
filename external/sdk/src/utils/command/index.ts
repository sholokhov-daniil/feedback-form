import type { HttpClient } from '../http/client';
import type { CommandRepository } from '../../commands/command';
import { Config } from '../config/index';

export function defineCommands(client: HttpClient): CommandRepository {
    const rep: CommandRepository = new Map();

    for (let key in Config.commands) {
        rep.set(key, new Config.commands[key](client));
    }

    return rep;
}