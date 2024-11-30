# subtitle ass shifter
Shift Subtitle of [`.ass`, `.srt`] format

> [!CAUTION]
> It will replace your subtitle, **backup** your subtitle files before run.


## Usage
```
ass-shifter [path] -t [shift ms]

  Positional Variables:
    path   the subtitle path to shift (Required)
  Flags:
       --version   Displays the program version string.
    -h --help      Displays help with available flag, subcommand, and positional value parameters.
    -t --shift     shift ms (default: 0)
    -s --start     start from HH:MM:SS
    -e --end       end at HH:MM:SS
```

## Example
The `start` and `end` parameters are optional. They can also be used together.
```bash
ass-shifter ../Better.Call.Saul/S03 -t -3200
ass-shifter ../Better.Call.Saul/S03 -t -3200 -s 0:06:13
ass-shifter ../Better.Call.Saul/S03 -t -3200 -s 0:06:13 -e 0:24:12
```

The program prints the result of the execution like the following.
```bash
ass-shifter ../Better.Call.Saul/S03 -t -3200 -s 0:06:13

[SUCCESS] Better.Call.Saul.S03E01.2017.1080p.BluRay.x265.10bit.ass
[SUCCESS] Better.Call.Saul.S03E02.2017.1080p.BluRay.x265.10bit.ass
[SUCCESS] Better.Call.Saul.S03E03.2017.1080p.BluRay.x265.10bit.ass
[SUCCESS] Better.Call.Saul.S03E04.2017.1080p.BluRay.x265.10bit.ass
[SUCCESS] Better.Call.Saul.S03E05.2017.1080p.BluRay.x265.10bit.ass
[SUCCESS] Better.Call.Saul.S03E06.2017.1080p.BluRay.x265.10bit.ass
[SUCCESS] Better.Call.Saul.S03E07.2017.1080p.BluRay.x265.10bit.ass
[SUCCESS] Better.Call.Saul.S03E08.2017.1080p.BluRay.x265.10bit.ass
[SUCCESS] Better.Call.Saul.S03E09.2017.1080p.BluRay.x265.10bit.ass
[SUCCESS] Better.Call.Saul.S03E10.2017.1080p.BluRay.x265.10bit.ass
Total 10 files shifted -3200ms from 0:06:13 to end
```
