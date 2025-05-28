# tspo2
## Эндпоинты
POST   /api/v1/auth/register         - Регистрация пользователя
POST   /api/v1/auth/login            - Вход в систему
POST   /api/v1/auth/refresh          - Обновление токена
GET    /api/v1/users/me              - Получение данных текущего пользователя
PUT    /api/v1/users/me              - Обновление данных пользователя
GET    /api/v1/users/me/addresses    - Получение адресов пользователя
POST   /api/v1/users/me/addresses    - Добавление адреса
PUT    /api/v1/users/me/addresses/:id - Обновление адреса
DELETE /api/v1/users/me/addresses/:id - Удаление адреса

## Каталог товаров
GET    /api/v1/categories            - Список категорий
GET    /api/v1/categories/:id        - Детали категории
GET    /api/v1/products              - Список товаров с фильтрацией
GET    /api/v1/products/:id          - Детали товара
GET    /api/v1/products/:id/reviews  - Отзывы о товаре
POST   /api/v1/products/:id/reviews  - Добавление отзыва

## Корзина
bash
GET    /api/v1/cart                  - Просмотр корзины
POST   /api/v1/cart/items            - Добавление товара в корзину
PUT    /api/v1/cart/items/:id        - Изменение количества товара
DELETE /api/v1/cart/items/:id        - Удаление товара из корзины
DELETE /api/v1/cart                  - Очистка корзины
Заказы
bash
POST   /api/v1/orders                - Создание заказа
GET    /api/v1/orders                - Список заказов пользователя
GET    /api/v1/orders/:id            - Детали заказа
GET    /api/v1/orders/:id/payment    - Информация об оплате
POST   /api/v1/orders/:id/payment    - Создание платежа
GET    /api/v1/orders/:id/delivery   - Информация о доставке

## Административные эндпоинты
bash
GET    /api/v1/admin/users           - Список пользователей
POST   /api/v1/admin/categories      - Добавление категории
PUT    /api/v1/admin/categories/:id  - Обновление категории
POST   /api/v1/admin/products        - Добавление товара
PUT    /api/v1/admin/products/:id    - Обновление товара
DELETE /api/v1/admin/products/:id    - Удаление товара
GET    /api/v1/admin/orders          - Список всех заказов
PUT    /api/v1/admin/orders/:id      - Обновление статуса заказа


digital-market/
├── cmd/
│   └── api/
│       └── main.go
├── config/
│   └── config.go
├── internal/
│   ├── models/
│   │   ├── user.go
│   │   ├── product.go
│   │   ├── category.go
│   │   ├── order.go
│   │   └── ...
│   ├── handlers/
│   │   ├── auth.go
│   │   ├── user.go
│   │   ├── product.go
│   │   └── ...
│   ├── repositories/
│   │   ├── user.go
│   │   ├── product.go
│   │   └── ...
│   ├── services/
│   │   ├── auth.go
│   │   ├── product.go
│   │   └── ...
│   └── middleware/
│       ├── auth.go
│       └── ...
├── pkg/
│   ├── database/
│   │   └── postgres.go
│   ├── jwt/
│   │   └── jwt.go
│   └── validator/
│       └── validator.go
├── go.mod
└── go.sum
