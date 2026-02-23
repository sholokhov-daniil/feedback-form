# Подключение SDK

1. Загрузите JavaScript API, используя ваш ключ сайта.

```html
<script src="your_service_url/sdk/api.js?siteId=my_id"></script>
```

2. Вызовите `sform.execute` для выполнения действия с формой
```js
sform.ready(function(){
    sform.execute({action: 'getForm', data: {id: 'your_form_id'}})
        .then(function(form) {
            // Add your logic
        })
})
```

## Список доступных действий

### Детальная информация по форме

Действие: getForm  
Описание: Возвращает детальную информацию по форме

**Параметры**
ТУТ ДОЛЖНА БЫТЬ ССЫЛКА НА SWAGGER

**Возвращаемые данные**
ТУТ ДОЛЖНА БЫТЬ ССЫЛКА НА SWAGGER


**Пример**
```js
sform.ready(function(){
    sform.execute({action: 'getForm', data: {id: 'your_form_id'}})
        .then(function(form) {
            // Add your logic
        })
        .catch(function(errors) {
            // Add your logic
        })
})


### Отправка данных формы

Действие: submit 
Описание: Производит отправку данных формы на сервер

**Параметры**
ТУТ ДОЛЖНА БЫТЬ ССЫЛКА НА SWAGGER

**Возвращаемые данные**
ТУТ ДОЛЖНА БЫТЬ ССЫЛКА НА SWAGGER

**Пример**
sform.ready(function(){
    sform.execute({action: 'submit', data: {fields: yourFields}})
        .then(function(response) {
            // Add your logic
        })
        .catch(function(errors) {
            // Add your logic
        })
});
