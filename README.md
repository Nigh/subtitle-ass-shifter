# subtitle-ass-shifter
Shift ASS format Subtitle


## Usage
```
ass-shifter [path] -t [shift ms]

  Positional Variables:
    path   the subtitle path to shift (Required)
  Flags:
       --version   Displays the program version string.
    -h --help      Displays help with available flag, subcommand, and positional value parameters.
    -t --shift     shift ms (default: 0)
```

## Example
```bash
go run main.go ../S05 -t -1002
[SUCCESS] Shifted -1002ms -> Better.Call.Saul.S05E01.2020.1080p.BluRay.x265.10bit.ass
[SUCCESS] Shifted -1002ms -> Better.Call.Saul.S05E02.2020.1080p.BluRay.x265.10bit.ass
[SUCCESS] Shifted -1002ms -> Better.Call.Saul.S05E03.2020.1080p.BluRay.x265.10bit.ass
[SUCCESS] Shifted -1002ms -> Better.Call.Saul.S05E04.2020.1080p.BluRay.x265.10bit.ass
[SUCCESS] Shifted -1002ms -> Better.Call.Saul.S05E05.2020.1080p.BluRay.x265.10bit.ass
[SUCCESS] Shifted -1002ms -> Better.Call.Saul.S05E06.2020.1080p.BluRay.x265.10bit.ass
[SUCCESS] Shifted -1002ms -> Better.Call.Saul.S05E07.2020.1080p.BluRay.x265.10bit.ass
[SUCCESS] Shifted -1002ms -> Better.Call.Saul.S05E08.2020.1080p.BluRay.x265.10bit.ass
[SUCCESS] Shifted -1002ms -> Better.Call.Saul.S05E09.2020.1080p.BluRay.x265.10bit.ass
[SUCCESS] Shifted -1002ms -> Better.Call.Saul.S05E10.2020.1080p.BluRay.x265.10bit.ass
```
