// При инициализации делаю запрос к микросервису передавая project_id и Origin
// Микросервис проверяет project_id и Origin
// Микросервис возвращает сессию и возвращает csrf


// TODO: mvp набросок
;(function(window){
    'use strict'

    if (window.SholokhovForm) {
        return;
    }

    window.SholokhovForm = new class {
        debug = true;
        _initialized = false;
        _version = '1.0';
        _api = 'https://feedback.127.0.0.1.nip.io';
        _csrfToken;

        init(siteId) {
            if (this._initialized) {
                return Promise.resolve();
            }

            log('Initializing SDK for project', siteId)

            // Делаем запрос на получение cookie и site_id
            return this._request(
                '/sdk/init',
                {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                        'VERSION-SDK': this._version,
                    },
                    credentials: 'include',
                    body: { siteId }
                }
            )
                .then(({csrf}) => {
                    this._csrfToken = csrf;
                    this._initialized = true;
                    this._log('initialized');
                })
        }

        _request(path, options = {}) {
            const headers = Object.assign({}, options.headers || {})

            if (this._csrfToken) {
                headers['X-CSRF-Token'] = this._csrfToken;
            }

            return fetch(
                this._api + path,
                {
                    method: options.method || 'POST',
                    headers,
                    credentials: 'include',
                    body: options.body ? JSON.stringify(options.body) : undefined
                }
            )
            .catch((res) => {
                this._log('Error request', res)
            })
        }

        _log(...args) {
            if (this.debug) {
                console.log('[Sholokhov form]', ...args);
            }
        }
    }
})(window);


// Это скрипт на стороне клиента, который размещается в head
(function(siteId) {

})('site_id');