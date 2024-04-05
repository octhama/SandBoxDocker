# Utiliser l'image officielle de WordPress
FROM wordpress:latest

# Définir l'auteur
LABEL maintainer="Hasler Tehou kuassi.tehou@gmail.com"

# Copier les fichiers de configuration personnalisés, thèmes et plugins
COPY custom-config.php /var/www/html/wp-content/
COPY custom-theme /var/www/html/wp-content/themes/custom-theme
COPY custom-plugin /var/www/html/wp-content/plugins/custom-plugin

# Exposer le port 80 pour le trafic web
EXPOSE 80

# Commande par défaut pour lancer WordPress
CMD ["apache2-foreground"]

