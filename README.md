# ГОСТ 28147-89 (шифр Магма)

Программа для шифрования и расшифрования файлов по стандарту ГОСТ 28147-89 (шифр Магма) 

## Install

```bash
go install github.com/vctriusss/magma-cipher/cmd/magma@latest
```

or

```bash
git clone github.com/vctriusss/magma-cipher.git
cd magma-cipher
go build cmd/magma/magma.go
```

## Usage

### Шифрование
```bash
magma encrypt -i <input-file> -o <output-file> -k <key-file>
```

Также можно не задавать ключ, тогда он будет сгенерирован и записан в файл `key.txt`. Тогда запуск программы будет выглядеть так
```bash
magma encrypt -i <input-file> -o <output-file>
```

### Расшифрование
```bash
magma decrypt -i <input-file> -o <output-file> -k <key-file>
```

### Помощь
```bash
magma help
magma -h
```

## Особенности
- написан на языке Go с использованием библиотеки `cli`
- реализован только единственный режим работы - простой (ECB)
- таблица замены (S-block) задана согласно спецификации ГОСТ и имеет идентификатор *id-tc26-gost-28147-param-Z*
- при работе в режиме простой замены блоки с недостающей длиной дополняются нулями справа
