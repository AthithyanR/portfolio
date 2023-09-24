install nginx - ```sudo apt install nginx```
note:- default static files under - ```ls /usr/share/nginx/html```

enable nginx service - ```sudo systemctl enable nginx```
to check nginx service status - ```sudo systemctl status nginx```

to start nginx service - ```sudo systemctl start nginx```
to stop nginx service - ```sudo systemctl stop nginx```
to reload nginx service - ```sudo systemctl reload nginx```
to restart nginx service - ```sudo systemctl restart nginx```

for SSL:
ref - https://www.nginx.com/blog/using-free-ssltls-certificates-from-lets-encrypt-with-nginx/
install certbot - ```sudo apt install certbot python3-certbot-nginx```
request certifice - ```sudo certbot --nginx -d athi.online -d www.athi.online```
* restart nginx service *

reverse proxy: