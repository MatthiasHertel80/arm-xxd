# arm-xxd

`arm-xxd` is a simplified reimplementation of the `xxd` command-line utility, written in Go. The `xxd` utility creates a hexdump of a given file or converts a hexdump back to binary. The `arm-xxd` version currently supports creating a hexdump, with the optional feature to generate a hexdump in C programming language syntax.

## Usage

```bash
go build arm-xxd.go
./arm-xxd [-i] [filename]
```

If no filename is given, `arm-xxd` reads from standard input.

## Options
- -i: Output in C include file style. If this option is used, the output is a hexdump in C programming language syntax.

## Example
```bash
echo "Hello, World!" | ./arm-xxd -i
```

Output:
```C
unsigned char data[] = {
  0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x2c, 
  0x20, 0x57, 0x6f, 0x72, 0x6c, 0x64, 
  0x21, 0x0a, 
};
```

## Limitations
`arm-xxd` does not currently support all of the features of the original xxd tool. In particular, it does not support the reverse operation, which converts a hexdump back into binary.

Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License
MIT?




