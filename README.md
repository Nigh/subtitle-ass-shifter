# subtitle ass shifter
Shift Subtitle of [`.ass`, `.srt`] format

> [!CAUTION]
> It will replace your subtitle, **backup** before run.


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
ass-shifter ../BCS -t -1002
[SUCCESS] Shifted -1002ms -> Better Call Saul - 1x01 - Uno.HDTV.KILLERS.en.srt
[SUCCESS] Shifted -1002ms -> Better Call Saul - 1x02 - Mijo.HDTV.LOL.en.srt
[SUCCESS] Shifted -1002ms -> Better Call Saul - 1x03 - Nacho.HDTV.x264-LOL.en.srt
[SUCCESS] Shifted -1002ms -> Better Call Saul - 1x04 - Hero.HDTV.LOL.en.srt
[SUCCESS] Shifted -1002ms -> Better Call Saul - 1x05 - Alpine Shepherd Boy.HDTV.x264-LOL.en.srt
[SUCCESS] Shifted -1002ms -> Better.Call.Saul.S05E06.2020.1080p.BluRay.x265.10bit.ass
[SUCCESS] Shifted -1002ms -> Better.Call.Saul.S05E07.2020.1080p.BluRay.x265.10bit.ass
[SUCCESS] Shifted -1002ms -> Better.Call.Saul.S05E08.2020.1080p.BluRay.x265.10bit.ass
[SUCCESS] Shifted -1002ms -> Better.Call.Saul.S05E09.2020.1080p.BluRay.x265.10bit.ass
[SUCCESS] Shifted -1002ms -> Better.Call.Saul.S05E10.2020.1080p.BluRay.x265.10bit.ass
```
