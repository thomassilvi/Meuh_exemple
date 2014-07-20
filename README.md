Meuh_exemple
============

```
 ______________________________________
/ Voici un exemple minimal de site web \
\ utilisant Meuh                       /
 --------------------------------------
        \   ^__^
         \  (oo)\_______
            (__)\       )\/\
                ||----w |
                ||     ||

```

Prerequis:
- un port libre pour http en local
- un serveur de base de donnees MongoDb : serveur, port, nom de la base, login et password

Pour tester l'exemple

1. recuperation de l'exemple:
git clone https://github.com/thomassilvi/Meuh_exemple.git

2. mettez a jour votre GOPATH 

3. recuperation des librairies

```
cd Meuh_exemple
go get github.com/gorilla/sessions
go get labix.org/v2/mgo
go get code.google.com/p/go.crypto/pbkdf2
go get crypto/sha256
go get github.com/thomassilvi/Meuh
```

4. compilation

```
cd src/MeuhWebExemple
go install
cd ../../
```

5. modification du fichier bin/default.conf pour changer les donnees MongoDB

6. lancer l'application

cd bin
./MeuhWebExemple

7. utiliser un browser avec l'adresse http://localhost:9000
 

