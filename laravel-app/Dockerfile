FROM php:8.4-cli-alpine

# Install system libraries and required PHP extensions
RUN apk add --no-cache \
    curl \
    git \
    unzip \
    libzip-dev \
    libpng-dev \
    oniguruma-dev && \
    docker-php-ext-install pdo_mysql mbstring zip bcmath

# Copy Composer binary from official image
COPY --from=composer:2 /usr/bin/composer /usr/bin/composer

WORKDIR /var/www/html

# Copy composer dependency definitions
COPY composer.json composer.lock ./

# Install vendor dependencies
RUN composer install --no-dev --no-interaction --prefer-dist --no-scripts

# Copy full application code
COPY . .

# Generate autoload mappings
RUN composer dump-autoload --optimize

# Ensure write permissions for storage and bootstrap/cache
RUN chown -R www-data:www-data storage bootstrap/cache && \
    chmod -R 775 storage bootstrap/cache

EXPOSE 8000

CMD ["sh", "-c", "if [ ! -f .env ]; then cp .env.example .env && php artisan key:generate; fi && php artisan serve --host=0.0.0.0 --port=8000"]
