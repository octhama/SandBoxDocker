# Utilisez l'image de base officielle d'Odoo
FROM odoo:latest

# Définissez l'environnement de travail
ENV ODOO_USER=odoo
ENV ODOO_HOME=/opt/odoo
ENV ODOO_DATA=/var/lib/odoo
ENV ODOO_LOG=/var/log/odoo

# Définissez le répertoire de travail
WORKDIR $ODOO_HOME

# Exécutez les commandes en tant qu'utilisateur root pendant la construction de l'image
USER root

# Créez les répertoires nécessaires avec les permissions appropriées
RUN mkdir -p $ODOO_HOME $ODOO_DATA $ODOO_LOG \
    && chown -R $ODOO_USER:$ODOO_USER $ODOO_HOME $ODOO_DATA $ODOO_LOG

# Revenez à l'utilisateur odoo pour l'exécution du conteneur
USER $ODOO_USER

# Exposez le port par défaut d'Odoo
EXPOSE 8069

# Démarrez Odoo
CMD ["odoo", "--config=/etc/odoo/odoo.conf"]

