// !!!! Должен генерироваться автоматически и возвращать min версию
// Производит инициализацию sdk. 
(function(window, document){
    const SDK_NAME = 'SholokhovForm';
    const CONFIG_NAME = '___sholokhov_form_cfg';
    // URL должен генерироваться на основе .env
    const SDK_URL = 'feedback.127.0.0.1.nip.io/sdk/v1';
    // хеш файла по пути feedback.127.0.0.1.nip.io/sdk/v1, для защиты от подмены
    const INTEGRITY_HASH = '';

    const cfg = window[CONFIG_NAME] = window[CONFIG_NAME] || {};

    // Очередь функции ready()
    cfg.fns = cfg.fns || [];

    // Создаем глобальный объект sdk
    const sdk = window[SDK_NAME] = window[SDK_NAME] || {};
    sdk.ready = sdk.ready || function(fn) {
        cfg.fns.push(fn);
    }

    // Асинхронная загрузка основного SDK
    const script = document.createElement('script');
    script.type = 'text/javascript';
    script.async = true;
    script.charset = 'utf-8';
    script.src = SDK_URL;
    script.crossOrigin = 'anonymous';
    script.integrity = INTEGRITY_HASH;

    // CSP nonce поддержка
    const firstScript = document.querySelector('script[nonce]');
    const nonce = firstScript && (firstScript.nonce || firstScript.getAttribute('nonce'));
    if (nonce) {
        script.setAttribute('nonce', nonce);
    }

    // Вставляем скрипт на страницу
    const s = document.getElementsByTagName('script')[0];
    s.parentNode.insertBefore(script, s);
} )(window, document);