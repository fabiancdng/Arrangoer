<div align="center">
    <img src="assets/Arrangør-transparent.png" width="230px" />
    <hr>
    <h1>Arrangør</h1>
    <strong>
        Ein Discord Bot, der bei der Organisation des Programmier-Wettbewerbs hilft und zahlreiche Automatisierungen sowie ein (Web-) Dashboard bietet. 
    </strong><br><br>
</div>

---

## Funktionen des Bots
#### Begrüßung neuer Nutzer
Der Bot schickt automatisch eine Begrüßungsnachricht für neu gejointe Nutzer in den 'lobbychannel', welchen man in der config.json festlegen kann.

<img width="400px" src="assets/screenshots/welcome-message.png" />

## Einrichtung / Konfiguration des Bots
Der Bot verfügt über eine 'Config-Datei', in die alle Daten eingetragen werden können, die er braucht, um zu funktionieren. Eine ```example.config.yml``` befindet sich im Ordner 'config'. Diese kann einfach kopiert und/oder umbenannt werden zu ```config.yml```. Nachdem dies erledigt ist, muss die Config ausgefüllt werden (siehe Kommentare in der Datei).

Zudem muss in ```web/src/config.js``` noch die URL eingetragen werden, unter der der Bot (bzw. seine API) erreichbar ist.

### Woher bekomme ich Client Secret und Client ID?
Den Token, das Client Secret und die Client ID gibt es auf der Discord Developers Seite im Oauth Menü. Zudem muss dort eine 'redirect uri' eingetragen werden, also eine URL, an die der Nutzer nach dem Login mit Discord weitergeleitet wird. Dort muss die URL der API des Bots eingetragen werden + ```'/api/auth/callback'``` (wie im Bild zu sehen).

<img width="600px" src="assets/screenshots/discord-developers-oauth.png" />

Weitere Dokumentation folgt.

<br>

**Einsendung für den Programmier-Wettbewerb der "Digitalen Woche 2021 Leer" von Fabian Reinders.**