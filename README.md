# brutekit

[![Go Version](https://img.shields.io/badge/Go-1.21%2B-blue.svg)](https://golang.org/) [![License](https://img.shields.io/badge/license-MIT-red.svg)](LICENSE) ![Maintained](https://img.shields.io/badge/Maintained%3F-Yes-green.svg)

A high-performance wordlist mutation engine written in Go. Brutekit transforms simple keywords into comprehensive wordlists by applying common password patterns and variations. Perfect for penetration testing, password recovery, and security research. Many concepts and functions are directly inspired by [psudohash](https://github.com/t3l3machus/psudohash), but rewritten in Go for enhanced performance.

## Key Features

- **Lightning Fast**: Written in Go for maximum performance and minimal resource usage
- **Smart Mutations**: Generates variations using leet speak, case changes, and common patterns
- **Targeted Generation**: Focuses on realistic password patterns people actually use
- **Flexible Output**: Customizable output formats and mutation rules
- **Memory Efficient**: Streams results to file instead of holding them in memory
- **Customizable**: Easy to extend with new mutation rules and patterns

## Core Capabilities

- Advanced leet speak transformations (e.g., password → p@ssw0rd)
- Case variations (lower, upper, mixed)
- Common padding patterns (prefix/suffix symbols)
- Year appendages (single year, ranges, or specific years)
- Sequential numbering
- Custom pattern support

## Installation

```bash
# Clone the repository
git clone https://github.com/gnomegl/brutekit
cd brutekit

# Build the binary
go build
```

## Usage

Basic usage:
```bash
./brutekit -w keyword -cpa
```

Advanced usage:
```bash
./brutekit -w keyword -cpa -an 3 -y 1990-2022
```

### Command Line Options

| Flag | Description |
|------|-------------|
| `-w` | Comma separated keywords to mutate (required) |
| `-an` | Append numbering range at end of mutations |
| `-nl` | Max numbering limit (default: 50) |
| `-y` | Years to append (e.g., 2022 or 1990,2017,2022 or 1990-2000) |
| `-ap` | Additional padding values |
| `-cpb` | Append common paddings before mutations |
| `-cpa` | Append common paddings after mutations |
| `-cpo` | Use only user provided paddings |
| `-o` | Output filename (default: output.txt) |
| `-q` | Do not print banner |

## Examples

1. Basic mutation with common paddings:
```bash
./brutekit -w company -cpa
```

2. Multiple keywords with numbering:
```bash
./brutekit -w "admin,root" -cpa -an 2
```

3. Keywords with year range:
```bash
./brutekit -w system -cpa -y 2020-2024
```

4. Custom output file:
```bash
./brutekit -w password -cpa -o wordlist.txt
```

## Customization

### Leet Speak Transformations
You can customize character substitutions by modifying the `transformations` map in `main.go`:

```go
var transformations = map[rune][]string{
    'a': {"@", "4"},
    'e': {"3"},
    'i': {"1", "!"},
    // Add your own transformations
}
```

### Common Padding Values
Edit `common_padding_values.txt` to customize the padding values used for mutations.

## Credits

- Original Python Implementation: [t3l3machus/psudohash](https://github.com/t3l3machus/psudohash)
- Go Implementation: [gnomegl/brutekit](https://github.com/gnomegl/brutekit)

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Disclaimer

This tool is intended for legal security testing purposes only. Users are responsible for complying with applicable laws and regulations. The authors assume no liability for misuse or damage caused by this program.
