
. "$PSScriptRoot/../intcode.ps1"

$input = (Get-Content "$PSScriptRoot/input.txt") -split ','

$input[1] = 12
$input[2] = 2

(Invoke-IntCode $input) -join ','

