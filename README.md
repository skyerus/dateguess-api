# history-api

## Debugger for VS-Code

Change program to your path in .vscode/launch.json

## Keys set up
```
cd keys
ssh-keygen -t rsa -b 4096 -m PEM -f RS256.key
openssl rsa -in RS256.key -pubout -outform PEM -out RS256.key.pub
```
