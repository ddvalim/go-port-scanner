# go-port-scanner

O `go-port-scanner` é um scanner TCP/IP construído na linguagem Go.

O TCP/IP é um protocolo de comunicação que opera na camada de transporte
do modelo de comunicação entre redes de computadores OSI. O modelo OSI é
composto por 7 camadas, cada uma com suas funcionalidades e protocolos.

![Modelo OSI](https://community.cisco.com/t5/image/serverpage/image-id/180291iDA59C8DFF9920CD8?v=v2) 

A camada de transporte é a quarta camada, responsável pela comunicação
ponta-a-ponta de dados, ao contrário da camada de rede, que trabalha com
os pacotes como entidades separadas. Ela garante que os pacotes de dados 
cheguem na ordem correta sem perda de informação, atuando como um 
intermediário entre as camadas de alto e baixo nível.

O protocolo TCP (Transport Control Protocol) é um protocolo baseado em conexão
segura, utilizado quando é essencial que os pacotes de dados cheguem intactos a
outra ponta. Outros protocolos como o UDP e o SCTP também operam nesta camada, 
ambos sem conexão.

A conexão ponta-a-ponta é feita por meio de um processo chamado **handshake de três
vias**, que garante que a comunição será feita de maneira segura e confiável. O handshake
é um acordo SYN/SYN-ACK/ACK. Após esses três passos, a conexão está estabelecida, e ambos 
os lados podem começar a trocar dados de forma confiável.