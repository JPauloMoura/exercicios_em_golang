1) Crie um programa que salve os exemplos de log em um aquivo de txto
use: `os.OpenFile()` -> `ponteiro.WriteString`

dica de flags:
```go
 // Exatamente um dos O_RDONLY, O_WRONLY, ou O_RDWR deve ser especificado. 
  O_RDONLY int = syscall . O_RDONLY  // abrir o arquivo somente leitura. 
  O_WRONLY int = syscall . O_WRONLY  // abrir o arquivo somente gravação. 
  O_RDWR int = syscall . O_RDWR    // abrir o arquivo de leitura e escrita. 

  // os valores restantes podem ser or'ed para controlar o comportamento. 
  O_APPEND int = syscall . O_APPEND  // anexar dados para o arquivo quando a escrita. 
  O_CREATE int = syscall. O_CREAT   // cria um novo arquivo se não houver nenhum. 
  O_EXCL int = syscall . O_EXCL    // usado com O_CREATE, o arquivo não deve existir. 
  O_SYNC int = syscall . O_SYNC    // abre para E / S síncrona. 
  O_TRUNC int = syscall . O_TRUNC   // trunca o arquivo gravável regular quando aberto. 
```

2) O go tem um jeito diferente de tratamento de datas, mostre alguns exemplos
use: `time.Now().Format()`

```go
//constants para formatação da data
  stdMonth                                       // "Jan"
  stdNumMonth                                    // "1"
  stdDay                                         // "2"
  stdLongYear              = iota + stdNeedDate  // "2006"
  stdYear                                        // "06"
  stdPM                    = iota + stdNeedClock // "PM"
  stdpm                                          // "pm"

  stdHour                  = iota + stdNeedClock // "15"
  stdHour12                                      // "3"
  stdZeroHour12                                  // "03"
  stdMinute                                      // "4"
  stdZeroMinute                                  // "04"
  stdSecond                                      // "5"
```
Exemplo: `` 