import { Config } from "./utils/config/index";
import { Logger } from "./utils/logger/index";
import { Client } from './utils/http/client';
import { defineCommands } from "./utils/command/index";
import type { CommandRepository } from "./commands/command";
import type { HttpClient } from './utils/http/client';
import type { EngineConfig } from './engine';

export class Engine {
    private client: HttpClient;
    private commandRepository: CommandRepository;

    constructor() {
        this.client = new Client;
        this.commandRepository = defineCommands(this.client);

        this.initReadyQueue();
    }

    /**
     * Публичный и безопасный метод взаимодействия с формой
     * 
     * @param {string} action Тип действия
     * @param {any} args Параметры действия
     */
    public execute(action: string, options: object): void {
        if (this.commandRepository.has(action) === false) {
            Logger.debug('Command not found');
            return;
        }

        const command = this.commandRepository.get(action);
        command?.execute(options);
    }

    /**
     * Возвращает конфигурацию SDK
     * 
     * @return {object}
     */
    public get config(): EngineConfig {
        const key = Config.configKey;
        const win = window as Window & Record<string, any>;

        if (!win[key]) {
            win[key] = {};
        }

        return win[key];
    }

    /**
     * Вызываются все, кто подписался на событие после инициализации SDK.
     * Пока еще в разработке
     */
    private initReadyQueue() {
        this.config?.fns?.forEach((fn: Function) => {
            try {
                fn(this);
            } catch (e) {
                Logger.error(e);
            }
        });
        this.config.fns = [];
    }
}