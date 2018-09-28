# airdb
Execute queries and analyse data from your databases through a easy golang web server for those who work with
alot of database data. 

![Alt text](ui/assets/img/preview.png?raw=true "Title")

# Installation
Clone the repo into your desired directory. Cd into the root of airdb and configure the application.

To config the application run the command below and follow the steps. airdb has to main deployment
methods being docker and supervisord. Docker deployment is not yet complete. 
```bash
sudo ./airdb -i 
```
# Deploy with supervisord
To run the application you must have installed supervisord on your server to finish the application
configuration. Once the configuration is completed successfully run the commands below to reread your 
supervisord config and start airdb in supervisorctl 

```bash
supervisorctl reread
```
```bash
supervisorctl update
```
```bash
supervisorctl start airdb
```
# ToDo
- Docker application installer
- Add users
