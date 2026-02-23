import type { CommandConstructor } from '../command/index';

export const Config = new class {
    /**
     * Возвращает URL сервиса управления формами
     * 
     * @return {string}
     */
    public get api(): string {
        return "https://" + String(process.env.HOST_NAME);
    }

    /**
     * Возвращает текущую веросию SDK
     * 
     * @return {string}
     */
    public get version(): string {
        return String(process.env.SDK_VERSION);
    }

    /**
     * Возвращает активности режима отладки
     * 
     * @return {boolean}
     */
    public get debug(): boolean {
        console.log(process.env.SDK_DEBUG);
        return Number(process.env.SDK_DEBUG || 0) === 1;
    }

    /**
     * Возвращает название глобального ключа хранящего конфигурацию SDK
     * 
     * @return {string}
     */
    public get configKey(): string
    {
        return String(process.env.SDK_PUBLIC_CONFIG_NAME || '___sholokhov_form_cfg');
    }

    /**
     * Возвращает все зарегистрированные комманды
     * 
     * @return {Record<string, CommandConstructor>}
     */
    public get commands(): Record<string, CommandConstructor> {
        return {

        }
    }
};