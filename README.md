# subtitle ass shifter
Shift Subtitle of [`.ass`, `.srt`] format

> [!CAUTION]
> It will replace your subtitle, **backup** your subtitle files before run. Use `--dry` for test.  
> From version `v1.2.0`, the program will automatically convert subtitle files to UTF8 encoding.


## Usage
```
ass-shifter [path] -t [shift ms]

  Positional Variables:
    path   the subtitle path to shift (Required)
  Flags:
       --version       Displays the program version string.
    -h --help          Displays help with available flag, subcommand, and positional value parameters.
    -t --shift         shift ms (default: 0)
    -s --start         start from HH:MM:SS
    -e --end           end at HH:MM:SS
    -sr --startRegexp  start from regular expression
    -er --endRegexp    end at regular expression
    -d --dry           dry run
```

The `--start` and `--end` parameters can be used to qualify the time range of the subtitle offset.   
The `--startRegexp` and `--endRegexp` parameters can be used to match the content of the subtitle with a regular expression as the start and end of the offset time range.  
The time and regular expression parameters can be used together.

## Example
The `start` and `end` parameters are optional. They can also be used together.
```bash
ass-shifter ../Better.Call.Saul/S03 -t 3200
ass-shifter ../Better.Call.Saul/S03 -t 3200 -s 0:06:13
ass-shifter ../Better.Call.Saul/S03 -t -3200 -s 0:06:13 -e 0:24:12
ass-shifter ../Better.Call.Saul/S06 -sr "第.季\s*第.+集" -t 3200
ass-shifter ../Better.Call.Saul/S06 -sr "第.季\s*第.+集" -e 0:24:12 -t 3200
```

The program prints the result of the execution like the following.
```bash
ass-shifter ../Better.Call.Saul/S06 -t -3200 -s 0:06:13
Better.Call.Saul.S06E01.2022.1080p.WEB-DL.x265.10bit.ass
From 0:06:13.00 to end, 1152 lines shifted 3200ms

Better.Call.Saul.S06E02.2022.1080p.WEB-DL.x265.10bit.ass
From 0:06:13.00 to end, 1222 lines shifted 3200ms

...

Better.Call.Saul.S06E12.2022.1080p.WEB-DL.x265.10bit.ass
From 0:06:13.00 to end, 1035 lines shifted 3200ms

Better.Call.Saul.S06E13.2022.1080p.WEB-DL.x265.10bit.ass
From 0:06:13.00 to end, 1629 lines shifted 3200ms

[Info] 13 subtitle files updated.
```

```bash
ass-shifter ../Better.Call.Saul/S06 -sr "第.季\s*第.+集" -t 3234 --dry

Better.Call.Saul.S06E01.2022.1080p.WEB-DL.x265.10bit.ass
From 0:05:32.87 to end, 1155 lines shifted 3234ms

Better.Call.Saul.S06E02.2022.1080p.WEB-DL.x265.10bit.ass
From 0:07:20.77 to end, 1221 lines shifted 3234ms

...

Better.Call.Saul.S06E11.2022.1080p.WEB-DL.x265.10bit.ass
From 0:01:39.00 to end, 1387 lines shifted 3234ms

Better.Call.Saul.S06E12.2022.1080p.WEB-DL.x265.10bit.ass
From 0:02:50.43 to end, 1145 lines shifted 3234ms

Better.Call.Saul.S06E13.2022.1080p.WEB-DL.x265.10bit.ass
From 0:05:17.70 to end, 1651 lines shifted 3234ms

[Info] Dry run, no file changes.
```
